package gody

import (
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type DescOption struct {
	TableName string `validate:"required"`
	Format    string
	Header    bool
	Field     string
}

func Desc(option *DescOption, cmd *cobra.Command) {
	svc, err := NewService(
		viper.GetString("profile"),
		viper.GetString("region"),
	)
	table, err := svc.GetTable(option.TableName)
	if err != nil {
		cmd.Println("error to get table")
	}

	design := table.Design
	result := StructToMap(&design)
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

// https://qiita.com/keitaj/items/440a50a53c8980ee338f
func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		result[field] = value
	}

	return result
}
