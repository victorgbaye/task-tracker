package main

import "net/http"

func (app application) createTask(w http.ResponseWriter, r *http.Request)  {
	
}
func (app application) getAllTask(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello world"))
}
func (app application) getTask(w http.ResponseWriter, r *http.Request)  {
	
}
func (app application) updateTask(w http.ResponseWriter, r *http.Request)  {
	
}
func (app application) deleteTask(w http.ResponseWriter, r *http.Request)  {
	
}