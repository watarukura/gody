package gody

import (
	"log"
	"fmt"
	"github.com/spf13/viper"
)

func List() {
	fmt.Println("eee" + viper.GetString("profile"))
	svc, err := NewService(
		viper.GetString("profile"),
		viper.GetString("region"),
	)
	tables, err := svc.ListTables()
	if err != nil {
		log.Fatal("error to list tables")
	}
	for _, table := range tables {
		fmt.Println(table)
	}
}
