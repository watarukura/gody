package gody

import (
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"log"
	"fmt"
)

type GetOption struct {
	TableName string
	HashKey   string
	RangeKey  string
}

func Get(svc *dynamodb.DynamoDB, getOption *GetOption) {
	table, err := svc.GetTable(getOption.TableName)
	if err != nil {
		log.Fatal("error to get table")
	}

	var result map[string]interface{}
	if getOption.RangeKey == "" {
		result, err = table.GetOne(getOption.HashKey)
	} else {
		result, err = table.GetOne(getOption.HashKey, getOption.RangeKey)
	}
	if err != nil {
		log.Fatal("error to get item")
		panic("error to get item")
	}
	for k,v := range result {
		fmt.Print(k + ": ")
		fmt.Println(v)
	}
}
