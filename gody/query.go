package gody

import (
	"log"
	"github.com/spf13/viper"
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"strings"
)

type QueryOption struct {
	TableName    string `validate:"required"`
	PartitionKey string `validate:"required"`
	SortKey      string
	Format       string
	Header       bool
	Limit        int64
	Index        string
	Eq           bool
	Lt           bool
	Le           bool
	Gt           bool
	Ge           bool
	Field        string
}

func Query(option *QueryOption) {
	svc, err := NewService(
		viper.GetString("profile"),
		viper.GetString("region"),
	)
	table, err := svc.GetTable(option.TableName)
	if err != nil {
		log.Fatal("error to get table")
	}

	cond := buildCondition(table, option)
	var query_result *dynamodb.QueryResult
	query_result, err = table.Query(cond)
	if err != nil {
		log.Fatal("error to query")
	}

	var query_result_remain *dynamodb.QueryResult
	for query_result.LastEvaluatedKey != nil {
		startKey := query_result.LastEvaluatedKey
		cond.SetStartKey(startKey)
		query_result_remain,err = table.ScanWithCondition(cond)
		if err != nil {
			log.Fatal("error to query for remain")
		}
		query_result.Items = append(query_result.Items, query_result_remain.Items...)
		query_result.LastEvaluatedKey = query_result_remain.LastEvaluatedKey
		query_result.Count += query_result.Count + query_result_remain.Count
		query_result.ScannedCount += query_result.ScannedCount+ query_result_remain.ScannedCount
	}
	result := query_result.ToSliceMap()

	var fields []string
	if option.Field != "" {
		fields = strings.Split(option.Field, ",")
	}

	Format(result, option.Format, option.Header, fields)
}

func buildCondition(table *dynamodb.Table, option *QueryOption) *dynamodb.ConditionList {
	cond := table.NewConditionList();
	design, err := table.Design();
	if err != nil {
		log.Fatal("error to describe table")
	}

	cond.SetLimit(option.Limit);
	var (
		pkey string
		skey string
	)
	if option.Index != "" {
		cond.SetIndex(option.Index)
		gsi := design.GSI;
		for _, v := range gsi {
			if *v.IndexName == option.Index {
				keys := v.KeySchema
				for _, vv := range keys {
					if *vv.KeyType == "HASH" {
						pkey = *vv.AttributeName
					}
					if *vv.KeyType == "RANGE" {
						skey = *vv.AttributeName
					}
				}
			}
		}
	} else {
		pkey = design.GetHashKeyName();
		skey = design.GetRangeKeyName();
	}

	cond.AndEQ(pkey, option.PartitionKey)

	if skey != "" && option.SortKey != "" {
		switch {
		case option.Eq == true:
			cond.AndEQ(skey, option.SortKey);
			break;
		case option.Lt == true:
			cond.AndLT(skey, option.SortKey);
			break;
		case option.Le == true:
			cond.AndLE(skey, option.SortKey);
			break;
		case option.Gt == true:
			cond.AndGT(skey, option.SortKey);
			break;
		case option.Ge == true:
			cond.AndGE(skey, option.SortKey);
			break;
		default:
			break;
		}
	}

	return cond
}
