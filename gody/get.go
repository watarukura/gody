package gody

import (
	"log"
	"github.com/spf13/viper"
	"strings"
)

type GetItemOption struct {
	TableName    string `validate:"required"`
	PartitionKey string `validate:"required"`
	SortKey      string
	Format       string
	Header       bool
	Field        string
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

	var result_slice []map[string]interface{}
	result_slice = append(result_slice, result)

	var fields []string
	if option.Field != "" {
		fields = strings.Split(option.Field, ",")
	}

	Format(result_slice, option.Format, option.Header, fields)
}
