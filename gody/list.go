package gody

import (
	"github.com/spf13/cobra"
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"log"
	"fmt"
	"github.com/evalphobia/aws-sdk-go-wrapper/config"
)

func List(*cobra.Command, []string) {
	svc, err := dynamodb.New(config.Config{
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
