package main

import (
	"fmt"
    "time"
    "database/sql"
)

type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"` 
}

type Todos struct {
	db *sql.DB
}


func NewTodos(db *sql.DB) *Todos {
	return &Todos{db: db}
}

func (t *Todos) Add(title, description string) error {

	_, err := t.db.Exec(SQLAddTodo, title, description, false)
	return err
}

func (t *Todos) Delete(id int) error {

	_, err := t.db.Exec(SQLDeleteTodo, id)
	return err
}



func (t *Todos) Update(id int, title *string, desc *string, completed *bool) (Todo, error) {
	var todo Todo
	
	err := t.db.QueryRow(SQLUpdateTodo, title, desc, completed, id).Scan(
		&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.CompletedAt,
	)

	if err == sql.ErrNoRows {
		return todo, fmt.Errorf("Туду не знайдено")
	}
	return todo, err
}

func (t *Todos) List() ([]Todo, error) {

	rows, err := t.db.Query(SQLListTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.CompletedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}