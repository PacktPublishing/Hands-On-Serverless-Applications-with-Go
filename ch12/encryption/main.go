package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

func handler() error {
	MYSQL_USERNAME := os.Getenv("MYSQL_USERNAME")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_HOST := os.Getenv("MYSQL_HOST")

	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE)
	db, err := sql.Open("mysql", uri)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Query(`CREATE TABLE IF NOT EXISTS movies(id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(50) NOT NULL)`)
	if err != nil {
		return err
	}

	for _, movie := range []string{"Iron Man", "Thor", "Avengers", "Wonder Woman"} {
		_, err := db.Query("INSERT INTO movies(name) VALUES(?)", movie)
		if err != nil {
			return err
		}
	}

	movies, err := db.Query("SELECT id, name FROM movies")
	if err != nil {
		return err
	}

	for movies.Next() {
		var name string
		var id int
		err = movies.Scan(&id, &name)
		if err != nil {
			return err
		}

		log.Printf("ID=%d\tName=%s\n", id, name)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
