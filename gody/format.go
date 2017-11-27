package gody

import (
	"fmt"
	"strings"
	"encoding/json"
)

var (
	head      []string
	body      [][]string
	body_unit []string
	delimiter string
)

func Format(ddbresult []map[string]interface{}, format string, header bool) {
	switch format {
	case "ssv":
		xsv(ddbresult, header, " ")
	case "csv":
		xsv(ddbresult, header, ",")
	case "tsv":
		xsv(ddbresult, header, "\t")
	case "json":
		toJson(ddbresult)
	}
}

func xsv(ddbresult []map[string]interface{}, header bool, delimiter string) {
	var head_dup []string
	for _, h := range ddbresult {
		for k, _ := range h {
			head_dup = append(head_dup, k)
		}
	}
	head := removeDuplicate(head_dup)

	var body_unit []string
	for _, v := range ddbresult {
		for _, h := range head {
			_,ok := v[h]
			if ok {
				body_unit = append(body_unit, fmt.Sprint(v[h]))
			} else {
				body_unit = append(body_unit, "_")
			}
		}
		body = append(body, body_unit)
		body_unit = make([]string, 0)
	}
	if header {
		fmt.Println(strings.Join(head, delimiter))
	}
	for _, b2 := range body {
		fmt.Println(strings.Join(b2, delimiter))
	}
}

func toJson(ddbresult []map[string]interface{}) {
	jsonString, _ := json.Marshal(ddbresult)
	fmt.Println(string(jsonString))
}

// https://qiita.com/hi-nakamura/items/5671eae147ffa68c4466
func removeDuplicate(args []string) []string {
	results := make([]string, 0, len(args))
	encountered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !encountered[args[i]] {
			encountered[args[i]] = true
			results = append(results, args[i])
		}
	}
	return results
}
