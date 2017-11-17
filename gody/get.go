package gody

import (
	"github.com/evalphobia/aws-sdk-go-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"log"
	"fmt"
	"github.com/spf13/cobra"
)

type GetFlag struct {
	TableName string
	HashKey   string
	RangeKey  string
}

func Get(*cobra.Command, []string) {
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

	//var getFlag GetFlag
	table, err := svc.GetTable("hn_item_store_data")
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
	result, err = table.GetOne("4901601353891")
	fmt.Println(result)

}
