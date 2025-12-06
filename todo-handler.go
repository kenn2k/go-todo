package main

import (
	"encoding/json"
	"net/http"
)


func listTodos(w http.ResponseWriter, todos *Todos) {
	list, err := todos.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}


func addTodo(w http.ResponseWriter, r *http.Request, todos *Todos) {

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Невалідний JSON", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "Title обов'язкове поле", http.StatusBadRequest)
		return
	}


	if err := todos.Add(req.Title, req.Description); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}


func updateTodo(w http.ResponseWriter, r *http.Request, todos *Todos, id int) {
	var req struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Completed   *bool   `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Невалідний JSON", http.StatusBadRequest)
		return
	}


	todo, err := todos.Update(id, req.Title, req.Description, req.Completed)
	if err != nil {
		if err.Error() == "todo not found" {
			http.Error(w, "Завдання не знайдено", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}


func deleteTodo(w http.ResponseWriter, todos *Todos, id int) {
	if err := todos.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}