package db

import (
	"database/sql"
	"fmt"
	"strings"
)

type DB_Class struct {
	db *sql.DB
}

type Select_Class struct {
	db   *sql.DB
	rows []string
}

type Select_Query_Class struct {
	db    *sql.DB
	rows  []string
	table string
}

func (db *DB_Class) Select(rows ...string) *Select_Class {
	return &Select_Class{db.db, rows}
}

func (query *Select_Class) From(table string) *Select_Query_Class {
	return &Select_Query_Class{query.db, query.rows, table}
}

func (query *Select_Query_Class) Query() []map[string]string {
	var rows_str []string
	for i := 0; i < len(query.rows); i++ {
		rows_str = append(rows_str, "\""+query.rows[i]+"\"")
	}
	var exec_query string = "select " + strings.Join(rows_str, ", ") + " from \"" + query.table + "\""
	fmt.Println(exec_query)
	rows, err := query.db.Query(exec_query)
	if err != nil {
		fmt.Println("Query error")
		return nil
	}
	cols, _ := rows.Columns()

	var result []map[string]string

	for rows.Next() {
		data := make(map[string]string)
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
		}
		result = append(result, data)
	}
	return result
}
