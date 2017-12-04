package gody

import (
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

func List(cmd *cobra.Command) {
	svc, err := NewService(
		viper.GetString("profile"),
		viper.GetString("region"),
	)
	tables, err := svc.ListTables()
	if err != nil {
		cmd.Println("error to list tables")
	}
	for _, table := range tables {
		cmd.Println(table)
	}
}
