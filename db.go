package main

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/jackc/pgx/v5/stdlib"
    "github.com/joho/godotenv"  
)

var DB *sql.DB

func InitDB() {

    if err := godotenv.Load(); err != nil {

		log.Println("Не знайдено файл .env") 
	}

    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL не встановлено")
    }

    var err error
    DB, err = sql.Open("pgx", dsn)
   

    if err = DB.Ping(); err != nil {
        log.Fatal("Не вдалося підключитися до БД:", err)
    }

   if _, err = DB.Exec(SQLCreateTable); err != nil {
		log.Fatal("Помилка створення таблиці:", err)
	}
   
	

    log.Println("Підключено до бази даних")
}