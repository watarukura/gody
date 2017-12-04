package gody

import (
	"fmt"
	"encoding/json"
	"encoding/csv"
	"os"
	"unicode/utf8"
	"sort"
	"github.com/spf13/cobra"
)

type FormatTarget struct {
	ddbresult []map[string]interface{}
	format    string
	header    bool
	fields    []string
	cmd       *cobra.Command
}

var (
	body [][]string
)

func Format(target FormatTarget) {
	switch target.format {
	case "ssv":
		toXsv(target, " ")
	case "csv":
		toXsv(target, ",")
	case "tsv":
		toXsv(target, "\t")
	case "json":
		toJson(target)
	}
}

func toXsv(target FormatTarget, delimiter string) {
	w := csv.NewWriter(os.Stdout)
	delm, _ := utf8.DecodeRuneInString(delimiter)
	w.Comma = delm

	// https://qiita.com/hi-nakamura/items/5671eae147ffa68c4466
	// headをユニークなsliceにする
	head := make([]string, 0, len(target.ddbresult))
	encountered := map[string]bool{}
	for _, v := range target.ddbresult {
		for k, _ := range v {
			if len(target.fields) > 0 {
				if !encountered[k] && Index(target.fields, k) > -1 {
					encountered[k] = true
					head = append(head, k)
				}
			} else {
				if !encountered[k] {
					encountered[k] = true
					head = append(head, k)
				}
			}
		}
	}
	// headerをsortしておおよそ同じ順に表示されるようにする
	sort.Strings(head)

	if target.header {
		w.Write(head)
		w.Flush()
	}

	var body_unit []string
	for _, v := range target.ddbresult {
		for _, h := range head {
			// 存在しないキーの場合は、値を"_"にする
			_, ok := v[h]
			if ok {
				body_unit = append(body_unit, fmt.Sprint(v[h]))
			} else {
				body_unit = append(body_unit, "_")
			}
		}
		body = append(body, body_unit)
		body_unit = make([]string, 0)
	}

	for _, b := range body {
		w.Write(b)
		w.Flush()
	}
}

func toJson(target FormatTarget) {
	jsonString, _ := json.Marshal(target.ddbresult)
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
