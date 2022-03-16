package routes

import (
	"log"
	"net/http"
	database "todos/api/config"
	todo "todos/api/pkg/todo"
)

/**
 * All API points for the application is present here.
 * @function Route
 */
func Route() {
	// Goes to config package to setup database connection to mongodb and also sets the collection
	database.Setup()

	// To manage routing in go.
	mux := http.NewServeMux()
	mux.HandleFunc("/todo", todo.Index)
	mux.HandleFunc("/todo/create", todo.Store)
	mux.HandleFunc("/todo/edit", todo.Update)
	mux.HandleFunc("/todo/delete", todo.Destory)
	mux.HandleFunc("/todo/mark-all", todo.MarkAll)

	// To start server at the host = localhost and port = 8080 with the given API endpoints.
	log.Println(http.ListenAndServe(":8080", mux))
}
