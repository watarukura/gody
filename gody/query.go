package gody

import (
	"log"
	"github.com/spf13/viper"
	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
)

type QueryOption struct {
	TableName    string
	PartitionKey string
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

	result := query_result.ToSliceMap()
	Format(result, option.Format, option.Header)
}

func buildCondition(table *dynamodb.Table, option *QueryOption) *dynamodb.ConditionList {
	cond := table.NewConditionList();
	design, err := table.Design();
	if err != nil {
		log.Fatal("error to describe table")
	}

	cond.SetLimit(option.Limit)
	pkey := design.GetHashKeyName();
	cond.AndEQ(pkey, option.PartitionKey)

	if design.HasRangeKey() {
		skey := design.GetRangeKeyName();
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
