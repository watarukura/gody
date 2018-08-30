package gody

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		viper.GetString("endpoint"),
	)
	table, err := svc.GetTable(option.TableName)
	if err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println("error to get table")
		cmd.Println(err)
		os.Exit(1)
	}

	var result map[string]interface{}
	if option.SortKey == "" {
		result, err = table.GetOne(option.PartitionKey)
	} else {
		result, err = table.GetOne(option.PartitionKey, option.SortKey)
	}
	if err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println("error to get item")
		cmd.Println(err)
		os.Exit(1)
	}

	var resultSlice []map[string]interface{}
	resultSlice = append(resultSlice, result)

	var fields []string
	if option.Field != "" {
		fields = strings.Split(option.Field, ",")
	}

	var formatTarget = FormatTarget{
		ddbresult: resultSlice,
		format:    option.Format,
		header:    option.Header,
		fields:    fields,
		cmd:       cmd,
	}
	Format(formatTarget)
}
