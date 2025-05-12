package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	json.NewEncoder(w).Encode(tasks)
}

func (app application) getTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "taskID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func (app application) updateTask(w http.ResponseWriter, r *http.Request) {
	var updated Task
	err := json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == updated.ID {
			tasks[i] = updated
			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func (app application) deleteTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == input.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}