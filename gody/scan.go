package gody

import (
	"log"
	"github.com/spf13/viper"
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"strings"
)

type ScanOption struct {
	TableName string `validate:"required"`
	Format    string
	Header    bool
	Limit     int64
	Field     string
}

func Scan(option *ScanOption) {
	svc, err := NewService(
		viper.GetString("profile"),
		viper.GetString("region"),
	)
	table, err := svc.GetTable(option.TableName)
	if err != nil {
		log.Fatal("error to get table")
	}

	cond := table.NewConditionList();
	cond.SetLimit(option.Limit);
	var query_result *dynamodb.QueryResult
	query_result, err = table.ScanWithCondition(cond)
	if err != nil {
		log.Fatal("error to scan")
	}

	result := query_result.ToSliceMap()

	var fields []string
	if option.Field != "" {
		fields = strings.Split(option.Field, ",")
	}

	Format(result, option.Format, option.Header, fields)
}
