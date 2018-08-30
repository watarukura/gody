package gody

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func List(cmd *cobra.Command) {
	svc, err := NewService(
		viper.GetString("profile"),
		viper.GetString("region"),
		viper.GetString("endpoint"),
	)
	tables, err := svc.ListTables()
	if err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println("error to list tables")
		cmd.Println(err)
		os.Exit(1)
	}
	for _, table := range tables {
		cmd.Println(table)
	}
}
