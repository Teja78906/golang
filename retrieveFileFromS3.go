package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client

func handler(w http.ResponseWriter, r *http.Request) {

	// Get the bucket name and file key from the URL
	bucketName := r.URL.Query().Get("bucket")
	fileKey := r.URL.Query().Get("key")

	// Retrieve the file from the bucket
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileKey),
	}
	result, err := s3Client.GetObject(context.TODO(), input)
	if err != nil {
		fmt.Fprintln(w, "unable to retrieve file")
		log.Println(err)
		return
	}

	// Write the file contents to the response
	_, err = io.Copy(w, result.Body)
	if err != nil {
		fmt.Fprintln(w, "unable to write file contents to response:")
		log.Println(err)
		return
	}
	defer result.Body.Close()
}

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile("profile_name"),
		config.WithRegion("Region!!"))
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	s3Client = s3.NewFromConfig(cfg)

	http.HandleFunc("/s3", handler)

	http.ListenAndServe(":8080", nil)
}
