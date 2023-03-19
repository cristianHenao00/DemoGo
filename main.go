package main

import (
	"net/http"

	"github.com/cristian409/DemoGo/db"
	"github.com/cristian409/DemoGo/models"
	"github.com/cristian409/DemoGo/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBconnection()
	db.DB.AutoMigrate(models.Tasks{})
	db.DB.AutoMigrate(models.User{})

	route := mux.NewRouter()

	route.HandleFunc("/", routes.HomeHandler)

	route.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	route.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	route.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	route.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	route.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	route.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	route.HandleFunc("/tasks", routes.PostTasksHandler).Methods("POST")
	route.HandleFunc("/tasks/{id}", routes.DeleteTasksHandler).Methods("DELETE")

	http.ListenAndServe(":3000", route)
}
