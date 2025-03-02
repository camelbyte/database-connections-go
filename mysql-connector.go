package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Tutorial struct {
	Title      string
	Subtitle   string
	Content    string
	CodeBlocks string
	Author     string
	Date       time.Time
	Views      int
}

func main() {

	dbi := "root:Sophia@tcp(localhost:3306)/tutorials"

	db, err := sql.Open("mysql", dbi)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// initialize the database (mariaDB), create the data to insert following the struct format.
	// db.Exec("CREATE DATABASE IF NOT EXISTS tutorials")
	// db.Exec("USE tutorials")
	// db.Exec("CREATE TABLE IF NOT EXISTS tutorials (id INT AUTO_INCREMENT PRIMARY KEY, title TEXT, content TEXT)")
	// db.Exec("INSERT INTO tutorials (title, content) VALUES ('First Post', 'Content of the first post.')")

	tutorial := Tutorial{
		Title:      "First Post",
		Subtitle:   "Subtitle of the first post",
		Content:    "Content of the first post.",
		CodeBlocks: "Code blocks of the first post.",
		Author:     "Keith Thomson",
		Date:       time.Now(),
	}

	// Insert the tutorial into the database
	query := "INSERT INTO tutorials (title, subtitle, content, codeBlock, Author, Date) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, tutorial.Title, tutorial.Subtitle, tutorial.Content, tutorial.CodeBlocks, tutorial.Author, tutorial.Date)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Object added to database.")

}
