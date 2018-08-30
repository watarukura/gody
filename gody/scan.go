package gody

import (
	"os"
	"strings"

	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ScanOption struct {
	TableName string `validate:"required"`
	Format    string
	Header    bool
	Limit     int64 `validate:"min=0"`
	Field     string
}

func Scan(option *ScanOption, cmd *cobra.Command) {
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

	cond := table.NewConditionList()
	if option.Limit > 0 {
		cond.SetLimit(option.Limit)
	}

	var queryResult *dynamodb.QueryResult
	queryResult, err = table.ScanWithCondition(cond)
	if err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println("error to scan")
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
			cmd.Println("error to scan for remain")
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
