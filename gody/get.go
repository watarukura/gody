package gody

import (
	"github.com/evalphobia/aws-sdk-go-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"log"
	"fmt"
)

type GetOption struct {
	TableName string
	HashKey   string
	RangeKey  string
}

func Get(getOption *GetOption) {
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

	fmt.Println(getOption)

	//var getFlag GetFlag
	table, err := svc.GetTable(getOption.TableName)
	if err != nil {
		log.Fatal("error to get table")
	}
	//fmt.Println(table)

	var result map[string]interface{}
	//if getFlag.RangeKey == "" {
	//	result, err = table.GetOne(getFlag.HashKey)
	//} else {
	//	result, err = table.GetOne(getFlag.HashKey, getFlag.RangeKey)
	//}
	//if err != nil {
	//	panic("error to get item")
	//}
	result, err = table.GetOne(getFlag.HashKey)
	fmt.Println(result)

}
