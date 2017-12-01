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
	Limit     int64  `validate:"min=0""`
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
	if option.Limit > 0 {
		cond.SetLimit(option.Limit);
	}

	var query_result *dynamodb.QueryResult
	query_result, err = table.ScanWithCondition(cond)
	if err != nil {
		log.Fatal("error to scan")
	}

	var query_result_remain *dynamodb.QueryResult
	// queryの結果が1MBを超えたときはLastEvaluatedKeyがセットされて、そこで切り上げられる
	for query_result.LastEvaluatedKey != nil {
		startKey := query_result.LastEvaluatedKey
		cond.SetStartKey(startKey)
		query_result_remain, err = table.ScanWithCondition(cond)
		if err != nil {
			log.Fatal("error to scan for remain")
		}
		query_result.Items = append(query_result.Items, query_result_remain.Items...)
		query_result.LastEvaluatedKey = query_result_remain.LastEvaluatedKey
		query_result.Count += query_result.Count + query_result_remain.Count
		query_result.ScannedCount += query_result.ScannedCount + query_result_remain.ScannedCount
	}
	result := query_result.ToSliceMap()

	var fields []string
	if option.Field != "" {
		fields = strings.Split(option.Field, ",")
	}

	Format(result, option.Format, option.Header, fields)
}
