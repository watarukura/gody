package gody

import (
	"fmt"
	"encoding/json"
	"encoding/csv"
	"os"
	"unicode/utf8"
	"sort"
)

var (
	body [][]string
)

func Format(ddbresult []map[string]interface{}, format string, header bool) {
	switch format {
	case "ssv":
		toXsv(ddbresult, header, " ")
	case "csv":
		toXsv(ddbresult, header, ",")
	case "tsv":
		toXsv(ddbresult, header, "\t")
	case "json":
		toJson(ddbresult)
	}
}

func toXsv(ddbresult []map[string]interface{}, header bool, delimiter string) {
	w := csv.NewWriter(os.Stdout)
	delm,_ := utf8.DecodeRuneInString(delimiter)
	w.Comma = delm

	// https://qiita.com/hi-nakamura/items/5671eae147ffa68c4466
	// headをユニークなsliceにする
	head := make([]string, 0, len(ddbresult))
	encountered := map[string]bool{}
	for _, v := range ddbresult {
		for k, _ := range v {
			if !encountered[k] {
				encountered[k] = true
				head = append(head, k)
			}
		}
	}
	// headerをsortしておおよそ同じ順に表示されるようにする
	sort.Strings(head)

	if header {
		w.Write(head)
		w.Flush()
	}

	var body_unit []string
	for _, v := range ddbresult {
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

func toJson(ddbresult []map[string]interface{}) {
	jsonString, _ := json.Marshal(ddbresult)
	fmt.Println(string(jsonString))
}
