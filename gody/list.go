package gody

import (
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"log"
	"fmt"
)

func List(svc *dynamodb.DynamoDB) {
	tables, err := svc.ListTables()
	if err != nil {
		log.Fatal("error to list tables")
	}
	fmt.Println(tables)
}
