package main

import (
    "log"
    "net/http"
    "os"
    
)

func main() {
    InitDB()
    defer DB.Close()

    todos := NewTodos(DB)
    SetupRoutes(todos)

    port := os.Getenv("PORT")
   
    
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
