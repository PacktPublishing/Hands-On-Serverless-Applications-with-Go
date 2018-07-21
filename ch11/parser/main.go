package main

import (
	"context"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-xray-sdk-go/xray"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context/ctxhttp"
)

// Movie entity
type Movie struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
}

func handler(ctx context.Context, url string) (Movie, error) {
	xray.Configure(xray.Config{
		LogLevel:       "info",
		ServiceVersion: "1.2.3",
	})

	// Get html page
	res, err := ctxhttp.Get(ctx, xray.Client(nil), url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Parse HTML response
	movie := Movie{}
	xray.Capture(ctx, "Parsing", func(ctx1 context.Context) error {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return err
		}

		title := doc.Find(".header .title span a h2").Text()
		description := doc.Find(".overview p").Text()
		cover, _ := doc.Find(".poster .image_content img").Attr("src")

		movie = Movie{
			ID:          uuid.Must(uuid.NewV4()).String(),
			Name:        title,
			Description: description,
			Cover:       cover,
		}

		xray.AddMetadata(ctx1, "movie.title", title)
		xray.AddMetadata(ctx1, "movie.description", description)
		xray.AddMetadata(ctx1, "movie.cover", cover)

		return nil
	})

	// Save in dynamodb
	sess := session.Must(session.NewSession())
	dynamo := dynamodb.New(sess)
	xray.AWS(dynamo.Client)
	_, err = dynamo.PutItemWithContext(ctx, &dynamodb.PutItemInput{
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
	if err != nil {
		log.Fatal(err)
	}

	return movie, nil
}

func main() {
	lambda.Start(handler)
}
