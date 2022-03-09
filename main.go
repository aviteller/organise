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
func addTodo(w http.ResponseWriter, r *http.Request) {
	var gi map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&gi)
	HandleError(err)
	gi["table"] = "todos"

	returnId := generalSqlInput(gi)
	var returnTodo Todo
	row := getOne("todos", int(returnId))
	row.Scan(&returnTodo.ID, &returnTodo.PersonID, &returnTodo.Content, &returnTodo.Priority, &returnTodo.Complete, &returnTodo.CreatedAt, &returnTodo.UpdatedAt, &returnTodo.Deleted)

	res := Message(true, "success")
	res["row"] = returnTodo
	Respond(w, res)
}
func addScheduleItem(w http.ResponseWriter, r *http.Request) {
	var gi map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&gi)
	HandleError(err)
	gi["table"] = "schedule"

	returnId := generalSqlInput(gi)
	var s Schedule
	row := getOne("schedule", int(returnId))
	row.Scan(&s.ID, &s.PersonID, &s.ScheduleCategoryID, &s.ScheduleFrequencyID, &s.ScheduleTime, &s.StartDate, &s.EndDate, &s.Desc, &s.CreatedAt, &s.UpdatedAt, &s.Deleted)

	res := Message(true, "success")
	res["row"] = s
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

func updateTodo(w http.ResponseWriter, r *http.Request) {
	var t map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	HandleError(err)
	fmt.Println("id", t["id"])
	if t["id"] != nil {
		// fmt.Println("todo", t)
		fmt.Println("here")
		updateOne("todos", t["id"], t)
	}
	res := Message(true, "success")
	Respond(w, res)
}
func updateScheduleItem(w http.ResponseWriter, r *http.Request) {
	var t map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	HandleError(err)
	fmt.Println("id", t["id"])
	if t["id"] != nil {
		// fmt.Println("todo", t)
		fmt.Println("here")
		updateOne("schedule", t["id"], t)
	}
	res := Message(true, "success")
	Respond(w, res)
}

func getScheduleHelpers(w http.ResponseWriter, r *http.Request) {

	var sCategories []interface{}
	var sFrequncies []interface{}

	rows := getAll("schedule_category", ArgsMap{"fields": "id, title"})

	for rows.Next() {
		s := struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		}{}
		rows.Scan(&s.ID, &s.Title)
		sCategories = append(sCategories, s)
	}
	fmt.Println("here")

	rows = getAll("schedule_frequency", ArgsMap{"fields": "id, title"})
	for rows.Next() {
		s := struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		}{}
		rows.Scan(&s.ID, &s.Title)
		sFrequncies = append(sFrequncies, s)
	}

	res := Message(true, "success")
	res["schedule_categories"] = sCategories
	res["schedule_frequencies"] = sFrequncies
	Respond(w, res)
}

func getHelpers(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	switch queryParams.Get("type") {
	case "schedules":
		getScheduleHelpers(w, r)
	default:
		res := Message(true, "failure")
		res["data"] = "To helper found"
		Respond(w, res)
	}

}

func getTodosByPersonID(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	routeID := getRouteID(r)

	// get query params
	queryParams := r.URL.Query()
	fmt.Println(queryParams)

	rows := getAllByID("todos", "person_id", routeID, ArgsMap{"orderBy": "ORDER BY complete, priority asc"})

	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.PersonID, &todo.Content, &todo.Priority, &todo.Complete, &todo.CreatedAt, &todo.UpdatedAt, &todo.Deleted)
		todos = append(todos, todo)
	}

	res := Message(true, "success")
	res["todos"] = todos
	Respond(w, res)
}

func getScheduleItemsByPersonID(w http.ResponseWriter, r *http.Request) {
	var schedule_items []Schedule
	routeID := getRouteID(r)

	// get query params
	queryParams := r.URL.Query()
	fmt.Println(queryParams)

	rows := getAllByID("schedule", "person_id", routeID, ArgsMap{"orderBy": "ORDER BY created_at asc"})

	for rows.Next() {
		var s Schedule
		rows.Scan(&s.ID, &s.PersonID, &s.ScheduleCategoryID, &s.ScheduleFrequencyID, &s.ScheduleTime, &s.StartDate, &s.EndDate, &s.Desc, &s.CreatedAt, &s.UpdatedAt, &s.Deleted)
		schedule_items = append(schedule_items, s)
	}

	res := Message(true, "success")
	res["schedule_items"] = schedule_items
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

//func that gets all Schedule from database in date range
func getScheduleByDateRange(w http.ResponseWriter, r *http.Request) {
	var schedule []Schedule
	routeID := getRouteID(r)

	// get query params
	queryParams := r.URL.Query()
	fmt.Println(queryParams)

	// add query params to args map
	args := ArgsMap{
		"orderBy": "ORDER BY created_at asc",
		"between": queryParams.Get("between"),
	}

	rows := getAllByID("schedule", "person_id", routeID, args)

	for rows.Next() {
		var s Schedule
		rows.Scan(&s.ID, &s.PersonID, &s.ScheduleCategoryID, &s.ScheduleFrequencyID, &s.ScheduleTime, &s.StartDate, &s.EndDate, &s.Desc, &s.CreatedAt, &s.UpdatedAt, &s.Deleted)
		schedule = append(schedule, s)
	}

	res := Message(true, "success")
	res["schedule"] = schedule
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
			if pathLength > 1 {
				getTodosByPersonID(w, r)
			}
		case "schedules":
			if pathLength > 1 {
				getScheduleItemsByPersonID(w, r)
			}
		case "calendar":
			if pathLength > 1 {
				getScheduleByDateRange(w, r)
			}
		case "helpers":
			getHelpers(w, r)

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
		case "todos":
			addTodo(w, r)
		case "schedules":
			addScheduleItem(w, r)

		default:
			fmt.Fprintf(w, "POST METHOD: %s DOES NOT EXIST", path)
		}
	case "PUT":
		switch pathParams[0] {
		case "todos":
			updateTodo(w, r)
		case "schedules":
			updateScheduleItem(w, r)

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
