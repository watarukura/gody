package gody

import (
	"github.com/spf13/cobra"
	"github.com/evalphobia/aws-sdk-go-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"log"
	"fmt"
)

type getFlag struct {
	tableName string
	hashKey   string
	rangeKey  string
}

func Get(*cobra.Command, []string) {
	svc, err := dynamodb.New(config.Config{
		//AccessKey: "access",
		//SecretKey: "secret",
		//Region:    "ap-northeast-1",
		//Endpoint:  "http://localhost:8000", // option for DynamoDB Local
		//Filename: "~/.aws/credentials",
		//Profile:  "CRM-STG",
		//Region:  region,
		//Profile: profile,
	})
	if err != nil {
		log.Fatal("error to create client")
	}

	table, err := svc.GetTable(tableName)
	if err != nil {
		log.Fatal("error to list tables")
	}
	fmt.Println(table)

	var result map[string]interface{}
	if rangeKey == "" {
		result, err = target.GetOne(hashKey)
	} else {
		result, err = target.GetOne(hashKey, rangeKey)
	}
	if err != nil {
		panic("error to get item")
	}
	fmt.Println(result)

}
