package main

import (
	"encoding/json"
	"net/http"
)

type Task struct{
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Done    bool   `json:"done"`
}

var tasks = []Task{}
var nextID = 1

func (app application) createTask(w http.ResponseWriter, r *http.Request)  {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
}

func (app application) getAllTask(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	
	err:= json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (app application) getTask(w http.ResponseWriter, r *http.Request)  {
	
}
func (app application) updateTask(w http.ResponseWriter, r *http.Request)  {
	
}
func (app application) deleteTask(w http.ResponseWriter, r *http.Request)  {
	
}