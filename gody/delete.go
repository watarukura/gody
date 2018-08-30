package gody

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type DeleteOption struct {
	TableName    string `validate:"required"`
	PartitionKey string `validate:"required"`
	SortKey      string
}

func Delete(option *DeleteOption, cmd *cobra.Command) {
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

	if option.SortKey == "" {
		err = table.Delete(option.PartitionKey)
	} else {
		err = table.Delete(option.PartitionKey, option.SortKey)
	}
	if err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println("error to delete item")
		cmd.Println(err)
		os.Exit(1)
	}
}
