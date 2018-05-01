package gody

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

const emptySymbol = "_"

type FormatTarget struct {
	ddbresult []map[string]interface{}
	format    string
	header    bool
	fields    []string
	cmd       *cobra.Command
}

func Format(target FormatTarget) {
	switch target.format {
	case "ssv":
		toXsv(target, " ")
	case "csv":
		toXsv(target, ",")
	case "tsv":
		toXsv(target, "\t")
	case "json":
		toJSON(target)
	default:
		target.cmd.SetOutput(os.Stderr)
		target.cmd.Println("choice format ssv|csv|tsv|json")
		os.Exit(1)
	}
}

func toXsv(target FormatTarget, delimiter string) {
	w := csv.NewWriter(target.cmd.OutOrStdout())
	delm, _ := utf8.DecodeRuneInString(delimiter)
	w.Comma = delm

	head := make([]string, 0, len(target.ddbresult))
	if len(target.fields) > 0 {
		head = target.fields
	} else {
		// https://qiita.com/hi-nakamura/items/5671eae147ffa68c4466
		// headをユニークなsliceにする
		encountered := map[string]bool{}
		for _, v := range target.ddbresult {
			for k := range v {
				if !encountered[k] {
					encountered[k] = true
					head = append(head, k)
				}
			}
		}
		// headerをsortしておおよそ同じ順に表示されるようにする
		sort.Strings(head)
	}

	if target.header {
		w.Write(head)
		w.Flush()
	}

	body := createBody(target, head)

	for _, b := range body {
		w.Write(b)
		w.Flush()
	}
}

func deprecatedCreateBody(target FormatTarget, head []string) [][]string {
	var (
		body [][]string
	)
	bodyUnit := []string{}
	for _, v := range target.ddbresult {
		for _, h := range head {
			// 存在しないキーの場合は、値を"_"にする
			_, ok := v[h]
			if ok {
				bodyUnit = append(bodyUnit, fmt.Sprint(v[h]))
			} else {
				bodyUnit = append(bodyUnit, "_")
			}
		}
		body = append(body, bodyUnit)
		bodyUnit = []string{}
	}
	return body
}

func createBody(target FormatTarget, head []string) [][]string {
	body := make([][]string, 0, len(target.ddbresult))

	for _, v := range target.ddbresult {
		bodyUnit := make([]string, 0, len(head))
		for _, h := range head {
			// 存在しないキーの場合は、値を"_"にする
			if _, ok := v[h]; ok {
				bodyUnit = append(bodyUnit, fmt.Sprint(v[h]))
			} else {
				bodyUnit = append(bodyUnit, emptySymbol)
			}
		}
		body = append(body, bodyUnit)
	}
	return body
}

func toJSON(target FormatTarget) {
	var jsonString []byte
	if len(target.fields) > 0 {
		m := make(map[string]interface{}, len(target.fields))
		marr := []map[string]interface{}{}
		for _, v := range target.ddbresult {
			for _, f := range target.fields {
				// 存在しないキーの場合は、値を"_"にする
				if _, ok := v[f]; ok {
					m[f] = v[f]
				} else {
					m[f] = emptySymbol
				}
			}
			marr = append(marr, m)
			m = map[string]interface{}{}
		}
		jsonString, _ = json.Marshal(marr)
	} else {
		jsonString, _ = json.Marshal(target.ddbresult)
	}
	target.cmd.Println(string(jsonString))
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
