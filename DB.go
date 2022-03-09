package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func getDB() *sql.DB {
	database, err := sql.Open("sqlite3", "./org.db")
	if err != nil {
		panic(err)
	}
	return database
}
func initSQLDB() {
	createTable("CREATE TABLE IF NOT EXISTS schedule (id INTEGER PRIMARY KEY,person_id INTEGER,schedule_cat_id INTEGER,schedule_frequency_id INTEGER,schedule_time TEXT, start_date TEXT, end_date TEXT, desc TEXT, created_at TEXT, updated_at TEXT NULL,deleted INTEGER DEFAULT 0)")

	createTable("CREATE TABLE IF NOT EXISTS schedule_category (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, created_at TEXT, updated_at TEXT NULL,deleted INTEGER DEFAULT 0)")

	createTable("CREATE TABLE IF NOT EXISTS schedule_frequency (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, created_at TEXT, updated_at TEXT NULL,deleted INTEGER DEFAULT 0)")
	// PersonParentLink create table

	// add link type next
	createTable("CREATE TABLE IF NOT EXISTS person_parent_link (id INTEGER PRIMARY KEY AUTOINCREMENT, person_id INTEGER, parent_id INTEGER, created_at TEXT, updated_at TEXT NULL,deleted INTEGER DEFAULT 0)")

	createTable("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, dob TEXT ,created_at TEXT, updated_at TEXT NULL,deleted INTEGER DEFAULT 0)")

	createTable("CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY AUTOINCREMENT, table_name TEXT, table_id INTEGER,  content TEXT, created_at TEXT, updated_at TEXT NULL,deleted INTEGER DEFAULT 0)")
	createTable("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY AUTOINCREMENT, person_id INTEGER, content TEXT, priority INTEGER, complete INTEGER DEFAULT 0, created_at TEXT, updated_at TEXT NULL,deleted INTEGER DEFAULT 0)")

}
func createTable(createTableString string) {
	stmt, err := getDB().Prepare(createTableString)
	HandleError(err)
	stmt.Exec()

}

func getTotalRowCount(query string) int {
	var count int
	err := getDB().QueryRow("SELECT COUNT(*) FROM (" + query + ")").Scan(&count)

	HandleError(err)
	return count
}

func getOne(table string, id int) *sql.Row {
	// fmt.Println("getOne", table, id)
	sqlStatment := "SELECT * FROM " + table + " WHERE id = ? AND deleted = 0"
	row := getDB().QueryRow(sqlStatment, id)
	return row

}

func updateOne(table string, id interface{}, args map[string]interface{}) int64 {
	keys := make([]string, 0, len(args)+1)
	values := make([]interface{}, 0, len(args)+1)

	for k, v := range args {
		if k == "id" {
			continue
		}
		keys = append(keys, k)
		values = append(values, v)

	}

	s := make([]interface{}, len(values))
	for i, v := range values {
		s[i] = fmt.Sprint(v)
	}

	keys = append(keys, "updated_at")
	s = append(s, time.Now().Format(time.RFC3339))
	s = append(s, id)

	sqlStatment := "UPDATE " + table + " SET " + strings.Join(keys, "=?,") + "=? WHERE id = ?"

	// fmt.Println(sqlStatment)
	// fmt.Println(s)

	stmt, err := getDB().Prepare(sqlStatment)

	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(s...)

	if err != nil {
		panic(err)
	}

	lastID, err := res.LastInsertId()

	return lastID
}

func getValues(args map[string]interface{}) []interface{} {
	var values []interface{}
	for _, value := range args {
		values = append(values, value)
	}
	return values
}

func getAllByID(table string, field string, id int, args map[string]interface{}) *sql.Rows {

	orderBy := ""
	between := ""
	// fmt.Println(args)

	if args["between"] != nil {
		betweenSlice := strings.Split(args["between"].(string), ",")
		between = " AND " + betweenSlice[0] + " BETWEEN '" + betweenSlice[1] + "' AND '" + betweenSlice[2] + "' "
		// fmt.Println("between", between)
		// between = args["orderBy"].(string)
	}

	if args["orderBy"] != nil {

		orderBy = args["orderBy"].(string)
	}

	sqlStatment := "SELECT * FROM " + table + " WHERE " + field + " = ? AND deleted = 0 " + between + orderBy

	fmt.Println(sqlStatment)

	// fmt.Println(sqlStatment, id)
	rows, err := getDB().Query(sqlStatment, id)

	HandleError(err)

	return rows
}

func generalQuery(sqlStatment string, withCount bool) []map[string]interface{} {
	// fmt.Println("generalQuery", sqlStatment)
	// fmt.Println(sqlStatment)
	if withCount {
		count := getTotalRowCount(sqlStatment)
		fmt.Println(count)
	}
	rows, err := getDB().Query(sqlStatment)
	HandleError(err)

	cols, err := rows.Columns()
	HandleError(err)

	var allMaps []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		pointers := make([]interface{}, len(cols))
		for i, _ := range values {
			pointers[i] = &values[i]
		}
		err := rows.Scan(pointers...)
		HandleError(err)
		resultMap := make(map[string]interface{})
		for i, val := range values {
			// fmt.Printf("Adding key=%s val=%v\n", cols[i], val)
			resultMap[cols[i]] = val
		}
		allMaps = append(allMaps, resultMap)
	}

	return allMaps
}

func getTables() []string {
	rows, err := getDB().Query("SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'")
	HandleError(err)
	var tables []string
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		HandleError(err)
		tables = append(tables, name)
	}
	return tables
}

//get table schema by table name
func getTableColumns(table string) []string {

	rows, err := getDB().Query("SELECT name FROM PRAGMA_TABLE_INFO('" + table + "')")
	HandleError(err)
	var keys []string
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		HandleError(err)
		keys = append(keys, name)
	}
	return keys
}

func getAll(table string, args map[string]interface{}) *sql.Rows {

	orderBy := ""

	if args["orderBy"] != nil {
		orderBy = args["orderBy"].(string)
	}
	selectFields := "*"
	if args["fields"] != nil {
		selectFields = args["fields"].(string)
	}

	sqlStatment := "SELECT " + selectFields + " FROM " + table + " WHERE deleted = 0 " + orderBy

	fmt.Println(sqlStatment)
	rows, err := getDB().Query(sqlStatment)

	HandleError(err)

	return rows
}

func deleteRow(rowId int, table_name string) string {
	// sqlStatment := "UPDATE" + table_name + "SET deleted = 1 WHERE id = ?"
	sqlStatment := fmt.Sprintf("UPDATE "+table_name+" SET deleted = 1 WHERE id = %d", rowId)
	// fmt.Println(sqlStatment)
	stmt, err := getDB().Prepare(sqlStatment)
	HandleError(err)
	res, err := stmt.Exec()
	fmt.Println(res)
	HandleError(err)
	return "ROW " + strconv.Itoa(rowId) + " DELETED FROM TABLE " + table_name
}

// super dangourous function that allows input into any table quite easily
func generalSqlInput(body map[string]interface{}) int64 {

	keys := make([]string, 0, len(body)+1)
	values := make([]interface{}, 0, len(body)+1)
	var table interface{}
	for k, v := range body {
		if k == "table" {
			table = v
		} else {
			keys = append(keys, k)
			values = append(values, v)
		}
	}

	s := make([]string, len(values))
	for i, v := range values {
		s[i] = fmt.Sprint(v)
	}

	keys = append(keys, "created_at")
	s = append(s, time.Now().Format(time.RFC3339))

	sqlStatment := "INSERT INTO " + table.(string) + "(" + strings.Join(keys, ",") + ") VALUES ('" + strings.Join(s, "','") + "')"

	fmt.Println(sqlStatment)

	stmt, err := getDB().Prepare(sqlStatment)

	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec()

	if err != nil {
		panic(err)
	}

	lastID, err := res.LastInsertId()

	return lastID

}
