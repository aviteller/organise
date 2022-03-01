package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Message is exported
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond is exported
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	// w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(data)
}

// Hab
func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
