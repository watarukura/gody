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
	var body_unit []string
	for k, _ := range ddbresult[0] {
		head = append(head, k)
	}
	for i, h := range head {
		for _, v := range ddbresult {
			body_unit[i] = fmt.Sprint(v[h])
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
