package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	_ "github.com/go-sql-driver/mysql"
)

var encryptedMysqlUsername string = os.Getenv("MYSQL_USERNAME")
var encryptedMysqlPassword string = os.Getenv("MYSQL_PASSWORD")
var mysqlDatabase string = os.Getenv("MYSQL_DATABASE")
var mysqlPort string = os.Getenv("MYSQL_PORT")
var mysqlHost string = os.Getenv("MYSQL_HOST")
var decryptedMysqlUsername, decryptedMysqlPassword string

func decrypt(encrypted string) (string, error) {
	kmsClient := kms.New(session.New())
	decodedBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	input := &kms.DecryptInput{
		CiphertextBlob: decodedBytes,
	}
	response, err := kmsClient.Decrypt(input)
	if err != nil {
		return "", err
	}
	return string(response.Plaintext[:]), nil
}

func init() {
	decryptedMysqlUsername, _ = decrypt(encryptedMysqlUsername)
	decryptedMysqlPassword, _ = decrypt(encryptedMysqlPassword)
}

func handler() error {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", decryptedMysqlUsername, decryptedMysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)
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
