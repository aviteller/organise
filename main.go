package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	// u "./utils"
)

// super dangourous function that allows input into any table quite easily
func generalSqlInputRoute(w http.ResponseWriter, r *http.Request) {
	var gi map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&gi)
	HandleError(err)

	returnId := generalSqlInput(gi)
	res := Message(true, "success")
	res["id"] = returnId
	Respond(w, res)

}

func addTask(w http.ResponseWriter, r *http.Request) {
	var gi map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&gi)
	HandleError(err)
	gi["table"] = "tasks"

	gi["due_date"] = gi["due_datetime"].(string)[:10]
	gi["due_time"] = gi["due_datetime"].(string)[11:]
	delete(gi, "due_datetime")

	fmt.Println(gi)
	returnId := generalSqlInput(gi)
	res := Message(true, "success")
	res["id"] = returnId
	Respond(w, res)
}

func logger(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")

}

func getRoute(r *http.Request) string {
	path := r.URL.Path
	path = strings.Replace(path, "/api/", "", -1)
	pathParams := strings.Split(path, "/")

	return pathParams[0]
}
func getRouteID(r *http.Request) int {
	path := r.URL.Path
	path = strings.Replace(path, "/api/", "", -1)
	pathParams := strings.Split(path, "/")

	id, err := strconv.Atoi(pathParams[1])

	HandleError(err)

	return id
}

// func getTask(w http.ResponseWriter, r *http.Request) {
// 	routeID := getRouteID(r)
// 	var t Task
// 	row := getOne("tasks", routeID)
// 	row.Scan(&t.ID, &t.PersonID, &t.TaskCategoryID, &t.DueDate, &t.DueTime, &t.Desc, &t.CreatedAt, &t.UpdatedAt, &t.Deleted)
// 	res := Message(true, "success")

// 	res["task"] = t
// 	Respond(w, res)
// }
func getPerson(w http.ResponseWriter, r *http.Request) {

	var p Person
	var children []Person
	// var tasks []Task
	routeID := getRouteID(r)

	row := getOne("people", routeID)
	row.Scan(&p.ID, &p.Name, &p.DOB, &p.CreatedAt, &p.UpdatedAt, &p.Deleted)

	// if p.ParentID == 0 {
	children = getChildren(p.ID)
	// }

	// tasks = getTasks(p.ID)
	// fmt.Println(tasks)

	res := Message(true, "success")
	res["person"] = p
	res["children"] = children
	// res["tasks"] = tasks
	Respond(w, res)

}

type ArgsMap map[string]interface{}

func getTableNames(w http.ResponseWriter, r *http.Request) {
	tables := getTables()
	res := Message(true, "success")
	res["tables"] = tables
	Respond(w, res)
}

func getTableDetalis(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.Replace(path, "/api/", "", -1)
	pathParams := strings.Split(path, "/")
	columns := getTableColumns(pathParams[1])
	res := Message(true, "success")
	res["columns"] = columns
	Respond(w, res)
}

func generalQueryRoute(w http.ResponseWriter, r *http.Request) {
	var gi map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&gi)

	HandleError(err)
	if gi["query"] != nil {
		query := gi["query"].(string)
		fmt.Println(query)
		rows := generalQuery(query, true)
		res := Message(true, "success")
		res["rows"] = rows
		Respond(w, res)
	}
}

func getAllTaskTypes(w http.ResponseWriter, r *http.Request) {
	// var tasktypes []TaskCategory
	// rows := getAll("task_types", ArgsMap{})

	// for rows.Next() {
	// 	var tt TaskCategory
	// 	rows.Scan(&tt.ID, &tt.Name, &tt.CreatedAt, &tt.UpdatedAt, &tt.Deleted)
	// 	tasktypes = append(tasktypes, tt)
	// }
	// res := Message(true, "success")

	// res["task_types"] = tasktypes
	// Respond(w, res)
}

// func getTasks(id int) []Task {
// 	var tasks []Task
// 	rows := getAllByID("tasks", "person_id", id, ArgsMap{"orderBy": "ORDER BY due_date desc"})
// 	for rows.Next() {
// 		var t Task
// 		rows.Scan(&t.ID, &t.PersonID, &t.TaskCategoryID, &t.DueDate, &t.DueTime, &t.Desc, &t.CreatedAt, &t.UpdatedAt, &t.Deleted)
// 		tasks = append(tasks, t)
// 	}

// 	return tasks
// }

func getParentPersonLink(id string) []int {
	var parentPersonIds []int
	ids := generalQuery("SELECT person_id FROM person_parent_link WHERE parent_id = "+id, false)
	// fmt.Println("after genral query", ids)
	for _, id := range ids {
		intid := int(id["person_id"].(int64))
		parentPersonIds = append(parentPersonIds, intid)
	}
	return parentPersonIds
}

func getChildren(parent_id int) []Person {
	var children []Person
	childrenIds := getParentPersonLink(strconv.Itoa(parent_id))
	// fmt.Println("after get person parent", childrenIds)
	for _, id := range childrenIds {

		var child Person

		row := getOne("people", id)

		row.Scan(&child.ID, &child.Name, &child.DOB, &child.CreatedAt, &child.UpdatedAt, &child.Deleted)

		children = append(children, child)
	}

	return children
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	routeID := getRouteID(r)
	rows := getAllByID("todos", "person_id", routeID, ArgsMap{"orderBy": "ORDER BY priority asc"})

	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.PersonID, &todo.Content, &todo.Priority, &todo.Complete, &todo.CreatedAt, &todo.UpdatedAt, &todo.Deleted)
		todos = append(todos, todo)
	}

	res := Message(true, "success")
	res["todos"] = todos
	Respond(w, res)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	var people []Person
	fmt.Println(r.URL.Query())
	rows := getAll("people", ArgsMap{"orderBy": "ORDER BY dob asc"})

	for rows.Next() {
		var p Person
		rows.Scan(&p.ID, &p.Name, &p.DOB, &p.CreatedAt, &p.UpdatedAt, &p.Deleted)
		people = append(people, p)
	}

	res := Message(true, "success")
	res["people"] = people
	Respond(w, res)
}

func router(w http.ResponseWriter, r *http.Request) {
	logger(w, r)
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	method := r.Method
	path := r.URL.Path
	path = strings.Replace(path, "/api/", "", -1)
	pathParams := strings.Split(path, "/")
	pathLength := len(pathParams)
	fmt.Println(pathParams, pathLength)
	switch method {
	case "GET":
		switch pathParams[0] {

		case "gettables":
			getTableNames(w, r)
		case "gettableschema":
			getTableDetalis(w, r)
		case "greet":
			greet(w, r)
		case "todos":
			getTodos(w, r)

		case "people":
			if pathLength > 1 {
				getPerson(w, r)
			} else {
				getPeople(w, r)
			}
		case "generalquery":
			generalQueryRoute(w, r)
		case "generalinput":
			generalSqlInputRoute(w, r)
		case "tasks":
			if pathLength > 1 {
				//getTask(w, r)
			}

		case "tasktypes":

			getAllTaskTypes(w, r)

		default:
			fmt.Fprintf(w, "GET METHOD: %s DOES NOT EXIST", path)
		}

	case "POST":
		switch pathParams[0] {
		case "generalquery":
			generalQueryRoute(w, r)
		case "generalinput":
			generalSqlInputRoute(w, r)
		case "tasks":
			addTask(w, r)

		default:
			fmt.Fprintf(w, "POST METHOD: %s DOES NOT EXIST", path)
		}
	case "DELETE":
		if pathLength > 1 {
			deleteRowRoute(w, r)
		}

	default:
		fmt.Fprintf(w, "METHOD: %s IS NOT ALLOWED", method)
	}

}

func deleteRowRoute(w http.ResponseWriter, r *http.Request) {
	routeID := getRouteID(r)
	routeName := getRoute(r)

	rstring := deleteRow(routeID, routeName)
	res := Message(true, "success")
	res["msg"] = rstring
	Respond(w, res)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	static := http.FileServer(http.Dir("./static/public/"))
	http.Handle("/", static)
	initSQLDB()

	http.HandleFunc("/api/", router)
	http.ListenAndServe(":8080", nil)
}
