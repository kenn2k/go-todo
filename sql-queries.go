package main

const (
	
	SQLCreateTable = `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT DEFAULT '',
		completed BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		completed_at TIMESTAMP
	)`

	
	SQLListTodos = `
	SELECT id, title, description, completed, created_at, completed_at 
	FROM todos 
	ORDER BY id`

	
	SQLAddTodo = `
	INSERT INTO todos (title, description, completed) 
	VALUES ($1, $2, $3)`

	
	SQLDeleteTodo = `
	DELETE FROM todos 
	WHERE id = $1`

	
	SQLUpdateTodo = `
	UPDATE todos 
	SET 
		title = COALESCE($1, title),
		description = COALESCE($2, description),
		completed = COALESCE($3, completed),
		completed_at = CASE 
			WHEN COALESCE($3, completed) = TRUE AND completed = FALSE THEN CURRENT_TIMESTAMP 
			WHEN COALESCE($3, completed) = FALSE THEN NULL 
			ELSE completed_at 
		END
	WHERE id = $4
	RETURNING id, title, description, completed, created_at, completed_at`
)