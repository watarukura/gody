package gody

import (
	"log"
	"github.com/spf13/viper"
)

type GetItemOption struct {
	TableName    string
	PartitionKey string
	SortKey      string
	Format       string
	Header       bool
}

func Get(option *GetItemOption) {
	svc, err := NewService(
		viper.GetString("profile"),
		viper.GetString("region"),
	)
	table, err := svc.GetTable(option.TableName)
	if err != nil {
		log.Fatal("error to get table")
	}

	var result map[string]interface{}
	if option.SortKey == "" {
		result, err = table.GetOne(option.PartitionKey)
	} else {
		result, err = table.GetOne(option.PartitionKey, option.SortKey)
	}
	if err != nil {
		log.Fatal("error to get item")
	}

	Format(result, option.Format, option.Header)
}
