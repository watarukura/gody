package gody

import (
	"github.com/spf13/cobra"
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

func Get(option *GetItemOption, cmd *cobra.Command) {
	svc, err := NewService(
		viper.GetString("profile"),
		viper.GetString("region"),
	)
	table, err := svc.GetTable(option.TableName)
	if err != nil {
		cmd.Println("error to get table")
	}

	var result map[string]interface{}
	if option.SortKey == "" {
		result, err = table.GetOne(option.PartitionKey)
	} else {
		result, err = table.GetOne(option.PartitionKey, option.SortKey)
	}
	if err != nil {
		cmd.Println("error to get item")
	}

	var result_slice []map[string]interface{}
	result_slice = append(result_slice, result)

	var fields []string
	if option.Field != "" {
		fields = strings.Split(option.Field, ",")
	}

	var formatTarget = FormatTarget{
		ddbresult: result_slice,
		format:    option.Format,
		header:    option.Header,
		fields:    fields,
		cmd:       cmd,
	}
	Format(formatTarget)
}
