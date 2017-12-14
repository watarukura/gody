package gody

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"unicode/utf8"

	"github.com/evalphobia/aws-sdk-go-wrapper/dynamodb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type PutItemOption struct {
	TableName string `validate:"required"`
	Format    string
	File      string
}

func Put(option *PutItemOption, cmd *cobra.Command) {
	svc, err := NewService(
		viper.GetString("profile"),
		viper.GetString("region"),
	)
	table, err := svc.GetTable(option.TableName)
	if err != nil {
		cmd.Println("error to get table")
	}

	item := dynamodb.NewPutItem()
	attributes := buildAttribute(option, cmd)
	table.AddItem(item)
	for _, l := range attributes {
		for k, v := range l {
			item.AddAttribute(k, v)
		}

		err = table.Put()
		if err != nil {
			cmd.Println("error to put item")
		}
	}
}

func buildAttribute(option *PutItemOption, cmd *cobra.Command) []map[string]interface{} {
	var attrs []map[string]interface{}
	var reader *bufio.Reader

	if option.File != "" {
		file, err := os.OpenFile(option.File, os.O_RDONLY, 0600)
		if err != nil {
			cmd.Println("failed to read file")
		}
		reader = bufio.NewReader(file)
		defer file.Close()
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	switch option.Format {
	case "ssv":
		attrs = fromXsv(option, reader, " ", cmd)
	case "csv":
		attrs = fromXsv(option, reader, ",", cmd)
	case "tsv":
		attrs = fromXsv(option, reader, "\t", cmd)
	case "json":
		attrs = fromJSON(option, reader, cmd)
	}
	return attrs
}

func fromXsv(option *PutItemOption, reader *bufio.Reader, delimiter string, cmd *cobra.Command) []map[string]interface{} {
	var (
		csvReader *csv.Reader
		attrs     []map[string]interface{}
	)
	csvReader = csv.NewReader(reader)
	delm, _ := utf8.DecodeLastRuneInString(delimiter)
	csvReader.Comma = delm
	// ダブルクォートを厳密にチェックしない
	csvReader.LazyQuotes = true

	csvAll, err := csvReader.ReadAll()
	if err != nil {
		cmd.Println("failed to read csv file")
	}
	header := csvAll[0]
	body := csvAll[1:]
	attr := map[string]interface{}{}

	for _, line := range body {
		for i, field := range line {
			if field == "_" || field == "" {
				continue
			}
			attr[header[i]] = field
		}
		attrs = append(attrs, attr)
		attr = map[string]interface{}{}
	}
	return attrs
}

func fromJSON(option *PutItemOption, reader *bufio.Reader, cmd *cobra.Command) []map[string]interface{} {
	attrs := []map[string]interface{}{}
	attr := map[string]interface{}{}
	lineCount := 0
	var err error

	for {
		if err == io.EOF {
			return attrs
		}
		line, err := reader.ReadBytes('\n')
		if err != nil {
			cmd.Println("failed to read json")
		}
		lineCount++
		if len(line) == 0 {
			continue
		}

		err = json.Unmarshal(line, &attr)

		if err != nil {
			cmd.Println("invalid json")
		}
		attrs = append(attrs, attr)
	}
}
