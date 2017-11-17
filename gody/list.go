package gody

import (
	"github.com/spf13/cobra"
	"github.com/evalphobia/aws-sdk-go-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"log"
	"fmt"
)

func List(*cobra.Command, []string) {
	svc, err := dynamodb.New(config.Config{
		//AccessKey: "access",
		//SecretKey: "secret",
		//Region:    "ap-northeast-1",
		//Endpoint:  "http://localhost:8000", // option for DynamoDB Local
		//Filename: "~/.aws/credentials",
		//Profile:  "CRM-STG",
		Region:  "ap-northeast-1",
		Profile: "default",
	})
	if err != nil {
		log.Fatal("error to create client")
	}

	tables, err := svc.ListTables()
	if err != nil {
		log.Fatal("error to list tables")
	}
	fmt.Println(tables)

}
