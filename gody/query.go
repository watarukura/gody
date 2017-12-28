package gody

import (
	"os"
	"strings"

	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type QueryOption struct {
	TableName    string `validate:"required"`
	PartitionKey string `validate:"required"`
	SortKey      string
	Format       string
	Header       bool
	Limit        int64 `validate:"min=0"`
	Index        string
	Eq           bool
	Lt           bool
	Le           bool
	Gt           bool
	Ge           bool
	Field        string
}

func Query(option *QueryOption, cmd *cobra.Command) {
	svc, err := NewService(
		viper.GetString("profile"),
		viper.GetString("region"),
	)
	table, err := svc.GetTable(option.TableName)
	if err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println("error to get table")
		cmd.Println(err)
		os.Exit(1)
	}

	cond := buildCondition(table, option, cmd)
	if option.Limit > 0 {
		cond.SetLimit(option.Limit)
	}

	var queryResult *dynamodb.QueryResult
	queryResult, err = table.Query(cond)
	if err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println("error to query")
		cmd.Println(err)
		os.Exit(1)
	}

	var queryResultRemain *dynamodb.QueryResult
	// queryの結果が1MBを超えたときはLastEvaluatedKeyがセットされて、そこで切り上げられる
	for queryResult.LastEvaluatedKey != nil {
		startKey := queryResult.LastEvaluatedKey
		cond.SetStartKey(startKey)
		queryResultRemain, err = table.ScanWithCondition(cond)
		if err != nil {
			cmd.SetOutput(os.Stderr)
			cmd.Println("error to query for remain")
			cmd.Println(err)
			os.Exit(1)
		}
		queryResult.Items = append(queryResult.Items, queryResultRemain.Items...)
		queryResult.LastEvaluatedKey = queryResultRemain.LastEvaluatedKey
		queryResult.Count += queryResult.Count + queryResultRemain.Count
		queryResult.ScannedCount += queryResult.ScannedCount + queryResultRemain.ScannedCount
	}
	result := queryResult.ToSliceMap()

	var fields []string
	if option.Field != "" {
		fields = strings.Split(option.Field, ",")
	}

	var formatTarget = FormatTarget{
		ddbresult: result,
		format:    option.Format,
		header:    option.Header,
		fields:    fields,
		cmd:       cmd,
	}
	Format(formatTarget)
}

func buildCondition(table *dynamodb.Table, option *QueryOption, cmd *cobra.Command) *dynamodb.ConditionList {
	cond := table.NewConditionList()
	design, err := table.Design()
	if err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println("error to describe table")
		cmd.Println(err)
		os.Exit(1)
	}

	cond.SetLimit(option.Limit)
	var (
		pkey string
		skey string
	)
	if option.Index != "" {
		cond.SetIndex(option.Index)
		gsi := design.GSI
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
		pkey = design.GetHashKeyName()
		skey = design.GetRangeKeyName()
	}

	cond.AndEQ(pkey, option.PartitionKey)

	if skey != "" && option.SortKey != "" {
		switch {
		case option.Eq == true:
			cond.AndEQ(skey, option.SortKey)
			break
		case option.Lt == true:
			cond.AndLT(skey, option.SortKey)
			break
		case option.Le == true:
			cond.AndLE(skey, option.SortKey)
			break
		case option.Gt == true:
			cond.AndGT(skey, option.SortKey)
			break
		case option.Ge == true:
			cond.AndGE(skey, option.SortKey)
			break
		default:
			break
		}
	}

	return cond
}
