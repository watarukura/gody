package main

import (
	"github.com/evalphobia/aws-sdk-go-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"fmt"
	"os"
	"flag"
	"log"
)

func main() {
	var (
		method   string
		table    string
		hashkey  string
		rangekey string
		index    string
		profile  string
		region   string
	)

	flag.StringVar(&table, "table", "", "table name")
	flag.StringVar(&hashkey, "hashkey", "", "hash key")
	flag.StringVar(&rangekey, "rangekey", "", "range key")
	flag.StringVar(&index, "index", "", "GSI")
	flag.StringVar(&profile, "profile", "", "AWS profile")
	flag.Parse();
	method = flag.Args()[0]

	if profile == "" {
		profile = "default"
	}
	if region == "" {
		region = "ap-northeast-1"
	}

	// Create DynamoDB service
	svc, err := dynamodb.New(config.Config{
		//AccessKey: "access",
		//SecretKey: "secret",
		//Region:    "ap-northeast-1",
		//Endpoint:  "http://localhost:8000", // option for DynamoDB Local
		//Filename: "~/.aws/credentials",
		//Profile:  "CRM-STG",
		Region:  region,
		Profile: profile,
	})
	if err != nil {
		log.Fatal("error to create client")
	}

	switch method {
	case "list":
		{
			tables, err := svc.ListTables()
			if err != nil {
				log.Fatal("error to list tables")
			}
			fmt.Println(tables)
		}
	case "get":
		{
			target, err := svc.GetTable(table)
			if err != nil {
				panic("error to get table")
			}

			var result map[string]interface{}
			if rangekey == "" {
				result, err = target.GetOne(hashkey)
			} else {
				result, err = target.GetOne(hashkey, rangekey)
			}
			if err != nil {
				panic("error to get item")
			}
			fmt.Println(result)
		}
	case "query":
		{
			//target, err := svc.GetTable(table)
			//if err != nil {
			//	panic("error to get table")
			//}
			//
			//var result map[string]interface{}
			//if rangekey == "" {
			//	result, err = target.Query(ContionsList)
			//} else {
			//	result, err = target.GetOne(hashkey, rangekey)
			//}
			//if err != nil {
			//	panic("error to get item")
			//}
			//fmt.Println(result)
		}
	case "put":
		{

		}
	case "delete":
		{

		}
	}

	os.Exit(0)

	//// Create new DynamoDB item (row on RDBMS)
	//item := dynamodb.NewPutItem()
	//item.AddAttribute("user_id", 999)
	//item.AddAttribute("status", 1)
	//
	//// Add item to the put spool
	//table.AddItem(item)
	//
	//item2 := dynamodb.NewItem()
	//item.AddAttribute("user_id", 1000)
	//item.AddAttribute("status", 2)
	//item.AddConditionEQ("status", 3) // Add condition for write
	//table.AddItem(item2)
	//
	//// Put all items in the put spool
	//err = table.Put()
	//
	//// Use svc.PutAll() to put all of the tables,
	//// `err = svc.PutAll()`
	//
	//// Scan items
	//cond = table.NewConditionList()
	//cond.SetLimit(1000)
	//cond.FilterEQ("status", 2)
	//result, err = table.ScanWithCondition(cond)
	//data := result.ToSliceMap() // `result.ToSliceMap()` returns []map[string]interface{}
	//
	////Scan from last key
	//cond.SetStartKey(result.LastEvaluatedKey)
	//result, err = table.ScanWithCondition(cond)
	//data = append(data, result.ToSliceMap())
	//
	//// Query items
	//cond := table.NewConditionList()
	//cond.AndEQ("user_id", 999)
	//cond.FilterLT("age", 20)
	//cond.SetLimit(100)
	//result, err := table.Query(cond)
	//if err != nil {
	//	panic("error to query")
	//}
	//
	//// mapping result data to the struct
	//type User struct {
	//	ID     int64 `dynamodb:"user_id"`
	//	Age    int   `dynamodb:"age"`
	//	Status int   `dynamodb:"status"`
	//}
	//var list []*User
	//err = result.Unmarshal(&list)
	//if err != nil {
	//	panic("error to unmarshal")
	//}
	//
	//if len(list) == int(result.Count) {
	//	fmt.Println("success to get items")
	//}
}
