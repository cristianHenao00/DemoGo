package routes

import (
	"encoding/json"
	"net/http"

	"github.com/cristian409/DemoGo/db"
	"github.com/cristian409/DemoGo/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Tasks
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
	w.WriteHeader(http.StatusOK)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
	w.WriteHeader(http.StatusOK)
}

func PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&task)
	w.WriteHeader(http.StatusOK)
}

func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusOK)
}
