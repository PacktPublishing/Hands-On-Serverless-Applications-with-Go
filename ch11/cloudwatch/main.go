package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go/aws"
)

func insert() {
	// insert logic

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	}

	svc := cloudwatch.New(cfg)
	req := svc.PutMetricDataRequest(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("InsertMovie"),
		MetricData: []cloudwatch.MetricDatum{
			cloudwatch.MetricDatum{
				Dimensions: []cloudwatch.Dimension{
					cloudwatch.Dimension{
						Name:  aws.String("Environment"),
						Value: aws.String("production"),
					},
				},
				MetricName: aws.String("ActionMovies"),
				Value:      aws.Float64(1.0),
				Unit:       cloudwatch.StandardUnitCount,
			},
		},
	})
	_, err = req.Send()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	lambda.Start(insert)
}
