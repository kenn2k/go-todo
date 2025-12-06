package main

import (
	"net/http"
	"path"
	"strconv"
)

func SetupRoutes(todos *Todos) {

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listTodos(w, todos)
		case http.MethodPost:
			addTodo(w, r, todos)
		default:
			http.Error(w, "Неправильний метод", http.StatusMethodNotAllowed)
		}
	})


	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {

		idStr := path.Base(r.URL.Path)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Невалідний ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodPatch:
			updateTodo(w, r, todos, id)


		case http.MethodDelete:
			deleteTodo(w, todos, id)

		default:
			http.Error(w, "Неправильний метод", http.StatusMethodNotAllowed)
		}
	})
}