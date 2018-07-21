package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	uuid "github.com/satori/go.uuid"
)

// Movie entity
type Movie struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
}

func handler() (Movie, error) {
	// Get html page
	res, err := http.Get("https://www.themoviedb.org/movie/351286-jurassic-world-fallen-kingdom")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Parse HTML response
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find(".header .title span a h2").Text()
	description := doc.Find(".overview p").Text()
	cover, _ := doc.Find(".poster .image_content img").Attr("src")

	movie := Movie{
		ID:          uuid.Must(uuid.NewV4()).String(),
		Name:        title,
		Description: description,
		Cover:       cover,
	}

	// Save in dynamodb
	sess := session.Must(session.NewSession())
	dynamo := dynamodb.New(sess)
	req, _ := dynamo.PutItemRequest(&dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Item: map[string]*dynamodb.AttributeValue{
			"ID": &dynamodb.AttributeValue{
				S: aws.String(movie.ID),
			},
			"Name": &dynamodb.AttributeValue{
				S: aws.String(movie.Name),
			},
			"Cover": &dynamodb.AttributeValue{
				S: aws.String(movie.Cover),
			},
			"Description": &dynamodb.AttributeValue{
				S: aws.String(movie.Description),
			},
		},
	})
	err = req.Send()
	if err != nil {
		log.Fatal(err)
	}

	return movie, nil
}

func main() {
	lambda.Start(handler)
}
