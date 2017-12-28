package gody

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
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
		cmd.SetOutput(os.Stderr)
		cmd.Println("error to get table")
		cmd.Println(err)
		os.Exit(1)
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
			cmd.SetOutput(os.Stderr)
			cmd.Println("error to put item")
			cmd.Println(err)
			os.Exit(1)
		}
	}
}

func buildAttribute(option *PutItemOption, cmd *cobra.Command) []map[string]interface{} {
	var attrs []map[string]interface{}
	var reader *bufio.Reader

	if option.File != "" {
		file, err := os.OpenFile(option.File, os.O_RDONLY, 0600)
		if err != nil {
			cmd.SetOutput(os.Stderr)
			cmd.Println("failed to read file")
			cmd.Println(err)
			os.Exit(1)
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
	default:
		cmd.SetOutput(os.Stderr)
		cmd.Println("choice format ssv|csv|tsv|json")
		os.Exit(1)
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
		cmd.SetOutput(os.Stderr)
		cmd.Println("failed to read csv file")
		cmd.Println(err)
		os.Exit(1)
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

type LineReader interface {
	ReadBytes(delim byte) (line []byte, err error)
}

func fromJSON(option *PutItemOption, reader LineReader, cmd *cobra.Command) []map[string]interface{} {
	var line []byte
	attrs := []map[string]interface{}{}
	var v interface{}
	var err error

	for {
		if err == io.EOF {
			break
		}
		line, err = reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				cmd.SetOutput(os.Stderr)
				cmd.Println("failed to read json")
				cmd.Println(err)
				os.Exit(1)
				break
			}
		}
		if len(line) == 0 {
			continue
		}

		err = json.Unmarshal(line, &v)
		if err != nil {
			cmd.SetOutput(os.Stderr)
			cmd.Println("invalid json")
			cmd.Println(err)
			os.Exit(1)
		}

		// JSONの配列の場合は[]map[string]interface
		// JSONの場合はmap[string]interface
		// TODO:型名を文字列にして比較はダサいのでどうにかしたい
		vType := fmt.Sprint(reflect.TypeOf(v))
		if vType == "map[string]interface {}" {
			attrs = append(attrs, v.(map[string]interface{}))
		} else {
			for _, attr := range v.([]interface{}) {
				attrs = append(attrs, attr.(map[string]interface{}))
			}
		}
	}
	return attrs
}
