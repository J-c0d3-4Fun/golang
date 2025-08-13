package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	file, err := os.Open("/Users/jbrown/golang/realProjects/glogger/tester.txt")
	if err != nil {
		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", file, err)
	} else {
		defer file.Close()
		// Putting the object in the bucket
		output, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String("testbucket"),
			Key:    aws.String("foo"),
			Body:   file,
		})
		// I need to add a validation check for Files so we dont override them

		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Successfully uploaded object. ETag: %s\n", aws.ToString(output.ETag))

	}
}
