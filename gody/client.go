package gody

import (
	"io"
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"github.com/evalphobia/aws-sdk-go-wrapper/config"
	"log"
)

type Client struct {
	stdin io.Reader
	stdout io.Writer
	stderr io.Writer
	profile string
	region string
}

func NewClient(stdin io.Reader, stdout io.Writer, stderr io.Writer, profile string, region string) (*dynamodb.DynamoDB, error) {
	svc, err := dynamodb.New(config.Config{
		Region:  region,
		Profile: profile,
	})
	if err != nil {
		log.Fatal("create service error")
	}
	return svc, nil
}
