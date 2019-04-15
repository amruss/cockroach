// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.

package tpcc

import (
	"bytes"
	gosql "database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/exp/rand"
)

// partitioner encapsulates all logic related to partitioning discrete numbers
// of warehouses into disjoint sets of roughly equal sizes. Partitions are then
// evenly assigned "active" warehouses, which allows for an even split of live
// warehouses across partitions without the need to repartition when the active
// count is changed.
type partitioner struct {
	total  int // e.g. the total number of warehouses
	active int // e.g. the active number of warehouses
	parts  int // the number of partitions to break `total` into

	partBounds   []int       // the boundary points between partitions
	partElems    [][]int     // the elements active in each partition
	partElemsMap map[int]int // mapping from element to partition index
	totalElems   []int       // all active elements
}

func makePartitioner(total, active, parts int) (*partitioner, error) {
	if total <= 0 {
		return nil, errors.Errorf("total must be positive; %d", total)
	}
	if active <= 0 {
		return nil, errors.Errorf("active must be positive; %d", active)
	}
	if parts <= 0 {
		return nil, errors.Errorf("parts must be positive; %d", parts)
	}
	if active > total {
		return nil, errors.Errorf("active > total; %d > %d", active, total)
	}
	if parts > total {
		return nil, errors.Errorf("parts > total; %d > %d", parts, total)
	}

	// Partition boundary points.
	//
	// bounds contains the boundary points between partitions, where each point
	// in the slice corresponds to the exclusive end element of one partition
	// and and the inclusive start element of the next.
	//
	//  total  = 20
	//  parts  = 3
	//  bounds = [0, 6, 13, 20]
	//
	bounds := make([]int, parts+1)
	for i := range bounds {
		bounds[i] = (i * total) / parts
	}

	// Partition sizes.
	//
	// sizes contains the number of elements that are active in each partition.
	//
	//  active = 10
	//  parts  = 3
	//  sizes  = [3, 3, 4]
	//
	sizes := make([]int, parts)
	for i := range sizes {
		s := (i * active) / parts
		e := ((i + 1) * active) / parts
		sizes[i] = e - s
	}

	// Partitions.
	//
	// partElems enumerates the active elements in each partition.
	//
	//  total     = 20
	//  active    = 10
	//  parts     = 3
	//  partElems = [[0, 1, 2], [6, 7, 8], [13, 14, 15, 16]]
	//
	partElems := make([][]int, parts)
	for i := range partElems {
		partAct := make([]int, sizes[i])
		for j := range partAct {
			partAct[j] = bounds[i] + j
		}
		partElems[i] = partAct
	}

	// Partition reverse mapping.
	//
	// partElemsMap maps each active element to its partition index.
	//
	//  total        = 20
	//  active       = 10
	//  parts        = 3
	//  partElemsMap = {0:0, 1:0, 2:0, 6:1, 7:1, 8:1, 13:2, 14:2, 15:2, 16:2}
	//
	partElemsMap := make(map[int]int)
	for p, elems := range partElems {
		for _, elem := range elems {
			partElemsMap[elem] = p
		}
	}

	// Total elements.
	//
	// totalElems aggregates all active elements into a single slice.
	//
	//  total      = 20
	//  active     = 10
	//  parts      = 3
	//  totalElems = [0, 1, 2, 6, 7, 8, 13, 14, 15, 16]
	//
	var totalElems []int
	for _, elems := range partElems {
		totalElems = append(totalElems, elems...)
	}

	return &partitioner{
		total:  total,
		active: active,
		parts:  parts,

		partBounds:   bounds,
		partElems:    partElems,
		partElemsMap: partElemsMap,
		totalElems:   totalElems,
	}, nil
}

// randActive returns a random active element.
func (p *partitioner) randActive(rng *rand.Rand) int {
	return p.totalElems[rng.Intn(len(p.totalElems))]
}

// configureZone sets up zone configs for previously created partitions. By default it adds constraints
// in terms of racks, but if the zones flag is passed into tpcc, it will set the constraints based on the
// geographic zones provided.
func configureZone(db *gosql.DB, table, partition string, constraint int, zones []string) error {
	var constraints string
	if len(zones) > 0 {
		constraints = fmt.Sprintf("[+zone=%s]", zones[constraint])
	} else {
		constraints = fmt.Sprintf("[+rack=%d]", constraint)
	}

	// We are removing the EXPERIMENTAL keyword in 2.1. For compatibility
	// with 2.0 clusters we still need to try with it if the
	// syntax without EXPERIMENTAL fails.
	// TODO(knz): Remove this in 2.2.
	sql := fmt.Sprintf(`ALTER PARTITION %s OF TABLE %s CONFIGURE ZONE USING constraints = '%s'`,
		partition, table, constraints)
	_, err := db.Exec(sql)
	if err != nil && strings.Contains(err.Error(), "syntax error") {
		sql = fmt.Sprintf(`ALTER PARTITION %s OF TABLE %s EXPERIMENTAL CONFIGURE ZONE 'constraints: %s'`,
			partition, table, constraints)
		_, err = db.Exec(sql)
	}
	if err != nil {
		return errors.Wrapf(err, "Couldn't exec %q", sql)
	}
	return nil
}

// partitionObject partitions the specified object (TABLE or INDEX) with the
// provided name, given the partitioning. Callers of the function must specify
// the associated table and the partition's number.
func partitionObject(
	db *gosql.DB, p *partitioner, zones []string, obj, name, col, table string, idx int,
) error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "ALTER %s %s PARTITION BY RANGE (%s) (\n", obj, name, col)
	for i := 0; i < p.parts; i++ {
		fmt.Fprintf(&buf, "  PARTITION p%d_%d VALUES FROM (%d) to (%d)",
			idx, i, p.partBounds[i], p.partBounds[i+1])
		if i+1 < p.parts {
			buf.WriteString(",")
		}
		buf.WriteString("\n")
	}
	buf.WriteString(")\n")
	if _, err := db.Exec(buf.String()); err != nil {
		return errors.Wrapf(err, "Couldn't exec %q", buf.String())
	}

	for i := 0; i < p.parts; i++ {
		if err := configureZone(db, table, fmt.Sprintf("p%d_%d", idx, i), i, zones); err != nil {
			return err
		}
	}
	return nil
}

func partitionTable(
	db *gosql.DB, p *partitioner, zones []string, table, col string, idx int,
) error {
	return partitionObject(db, p, zones, "TABLE", table, col, table, idx)
}

func partitionIndex(
	db *gosql.DB, p *partitioner, zones []string, table, index, col string, idx int,
) error {
	if exists, err := indexExists(db, table, index); err != nil {
		return err
	} else if !exists {
		// If the index doesn't exist then there's nothing to do. This is the
		// case for a few of the indexes that are only needed for foreign keys
		// when foreign keys are disabled.
		return nil
	}
	indexStr := fmt.Sprintf("%s@%s", table, index)
	return partitionObject(db, p, zones, "INDEX", indexStr, col, table, idx)
}

func partitionWarehouse(db *gosql.DB, wPart *partitioner, zones []string) error {
	return partitionTable(db, wPart, zones, "warehouse", "w_id", 0)
}

func partitionDistrict(db *gosql.DB, wPart *partitioner, zones []string) error {
	return partitionTable(db, wPart, zones, "district", "d_w_id", 0)
}

func partitionNewOrder(db *gosql.DB, wPart *partitioner, zones []string) error {
	return partitionTable(db, wPart, zones, "new_order", "no_w_id", 0)
}

func partitionOrder(db *gosql.DB, wPart *partitioner, zones []string) error {
	if err := partitionTable(db, wPart, zones, `"order"`, "o_w_id", 0); err != nil {
		return err
	}
	return partitionIndex(db, wPart, zones, `"order"`, "order_idx", "o_w_id", 1)
}

func partitionOrderLine(db *gosql.DB, wPart *partitioner, zones []string) error {
	if err := partitionTable(db, wPart, zones, "order_line", "ol_w_id", 0); err != nil {
		return err
	}
	return partitionIndex(db, wPart, zones, "order_line", "order_line_stock_fk_idx", "ol_supply_w_id", 1)
}

func partitionStock(db *gosql.DB, wPart *partitioner, zones []string) error {
	// The stock_item_fk_idx can't be partitioned because it doesn't have a
	// warehouse prefix. It's an all-around unfortunate index that we only
	// need because of a restriction in SQL. See #36859 and #37255.
	return partitionTable(db, wPart, zones, "stock", "s_w_id", 0)
}

func partitionCustomer(db *gosql.DB, wPart *partitioner, zones []string) error {
	if err := partitionTable(db, wPart, zones, "customer", "c_w_id", 0); err != nil {
		return err
	}
	return partitionIndex(db, wPart, zones, "customer", "customer_idx", "c_w_id", 1)
}

func partitionHistory(db *gosql.DB, wPart *partitioner, zones []string) error {
	if err := partitionTable(db, wPart, zones, "history", "h_w_id", 0); err != nil {
		return err
	}
	if err := partitionIndex(db, wPart, zones, "history", "history_customer_fk_idx", "h_c_w_id", 1); err != nil {
		return err
	}
	return partitionIndex(db, wPart, zones, "history", "history_district_fk_idx", "h_w_id", 2)
}

// replicateItem creates a covering "replicated index" for the item table for
// each of the zones provided. The item table is immutable, so this comes at a
// negligible cost and allows all lookups into it to be local.
func replicateItem(db *gosql.DB, zones []string) error {
	for i, zone := range zones {
		idxName := fmt.Sprintf("replicated_idx_%d", i)

		create := fmt.Sprintf(`
			CREATE UNIQUE INDEX %s
			ON item (i_id)
			STORING (i_im_id, i_name, i_price, i_data)`,
			idxName)
		if _, err := db.Exec(create); err != nil {
			return errors.Wrapf(err, "Couldn't exec %q", create)
		}

		configure := fmt.Sprintf(`
			ALTER INDEX item@%s
			CONFIGURE ZONE USING lease_preferences = '[[+zone=%s]]'`,
			idxName, zone)
		if _, err := db.Exec(configure); err != nil {
			return errors.Wrapf(err, "Couldn't exec %q", configure)
		}
	}
	return nil
}

func partitionTables(db *gosql.DB, wPart *partitioner, zones []string) error {
	if err := partitionWarehouse(db, wPart, zones); err != nil {
		return err
	}
	if err := partitionDistrict(db, wPart, zones); err != nil {
		return err
	}
	if err := partitionNewOrder(db, wPart, zones); err != nil {
		return err
	}
	if err := partitionOrder(db, wPart, zones); err != nil {
		return err
	}
	if err := partitionOrderLine(db, wPart, zones); err != nil {
		return err
	}
	if err := partitionStock(db, wPart, zones); err != nil {
		return err
	}
	if err := partitionCustomer(db, wPart, zones); err != nil {
		return err
	}
	if err := partitionHistory(db, wPart, zones); err != nil {
		return err
	}
	return replicateItem(db, zones)
}

func partitionCount(db *gosql.DB) (int, error) {
	var count int
	if err := db.QueryRow(`
		SELECT count(*)
		FROM crdb_internal.tables t
		JOIN crdb_internal.partitions p
		USING (table_id)
		WHERE t.name = 'warehouse'
		AND p.name ~ 'p0_\d+'
	`).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func indexExists(db *gosql.DB, table, index string) (bool, error) {
	var exists bool
	if err := db.QueryRow(`
		SELECT count(*) > 0
		FROM information_schema.statistics
		WHERE table_name = $1
		AND   index_name = $2
	`, table, index).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}
