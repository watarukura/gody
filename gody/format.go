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
	if header {
		for k,_ := range ddbresult[0] {
			head = append(head, k)
		}
		for _,b := range ddbresult {
			for _,v := range b {
				body_unit = append(body_unit, fmt.Sprint(v))
			}
			body = append(body, body_unit)
			body_unit = make([]string, 0)
		}
		fmt.Println(strings.Join(head, delimiter))
		for _, b2 :=  range body {
			fmt.Println(strings.Join(b2, delimiter))
		}
	} else {
		for _,b := range ddbresult {
			for _,v := range b {
				body_unit = append(body_unit, fmt.Sprint(v))
			}
			body = append(body, body_unit)
		}
		for _, b2 :=  range body {
			fmt.Println(strings.Join(b2, delimiter))
		}
	}
}

func toJson(ddbresult []map[string]interface{}) {
	jsonString, _ := json.Marshal(ddbresult)
	fmt.Println(string(jsonString))
}
