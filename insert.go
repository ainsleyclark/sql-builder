// Copyright 2020 The GoQueryBuilder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builder

import (
	"fmt"
	"reflect"
	"strings"
)

// Insert
//
// Sets the insert table statement to the table name.
func (s *Sqlbuilder) Insert(table string) *Sqlbuilder {
	s.insertStmt = table
	return s
}

// Update
//
// Sets the update table statement to the table name.
func (s *Sqlbuilder) Update(table string) *Sqlbuilder {
	s.updateStmt = table
	return s
}

// Column
//
// Adds a string column for INSERT & UPDATE statements.
func (s *Sqlbuilder) Column(column string, value interface{}) *Sqlbuilder {
	val, ok := printInterface(reflect.ValueOf(value))
	if !ok {
		val = "NULL"
	}
	col := [2]string{column, val}
	s.columns = append(s.columns, col)
	return s
}

// Increment
//
// Adds 1 to a column.
func (s *Sqlbuilder) Increment(column string) *Sqlbuilder {
	col := [2]string{column, fmt.Sprintf("%s + 1", column)}
	s.columns = append(s.columns, col)
	return s
}

// Decrement
//
// Subtracts 1 to a column.
func (s *Sqlbuilder) Decrement(column string) *Sqlbuilder {
	col := [2]string{column, fmt.Sprintf("%s - 1", column)}
	s.columns = append(s.columns, col)
	return s
}

func (s *Sqlbuilder) Returning(field string) *Sqlbuilder {
	s.returning = " RETURNING " + field
	return s
}

// buildUpdate
//
// Creates the insert statement when `Build()` is called.
func (s *Sqlbuilder) buildUpdate() string {
	var set string
	for _, v := range s.columns {
		set += s.formatSchema(v[0]) + " = " + v[1] + ", "
	}

	set = strings.TrimSuffix(set, ", ")

	// TODO: This should be within a `Where` function.
	var where string
	if s.whereStmt != `` {
		where = ` WHERE ` + strings.TrimSuffix(s.whereStmt, ` AND `) + ` `
	}

	returning := ""
	if s.returning != "" {
		returning = s.returning
	}

	if s.Dialect == "postgres" {
		return "UPDATE " + s.updateStmt + " SET " + set + where + returning
	}

	return "UPDATE `" + s.updateStmt + "` SET " + set + where + returning
}

// buildInsert
//
// Creates the update statement when `Build()` is called.
func (s *Sqlbuilder) buildInsert() string {
	var cols string
	var values string

	for _, v := range s.columns {
		cols += s.formatSchema(v[0]) + ", "
		values += v[1] + ", "
	}

	cols = strings.TrimSuffix(cols, ", ")
	values = strings.TrimSuffix(values, ", ")

	returning := ""
	if s.returning != "" {
		returning = s.returning
	}

	if s.Dialect == "postgres" {
		return "INSERT INTO " + s.insertStmt + " (" + cols + ") VALUES (" + values + ")" + returning
	}

	return "INSERT INTO `" + s.insertStmt + "` (" + cols + ") VALUES (" + values + ")" + returning
}
