package gody

import (
	"fmt"
	"strings"
	"encoding/json"
)

var (
	head      []string
	body      []string
	delimiter string
)

func Format(ddbresult map[string]interface{}, format string, header bool) {
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

func xsv(ddbresult map[string]interface{}, header bool, delimiter string) {
	if header {
		for k, v := range ddbresult {
			head = append(head, k)
			body = append(body, fmt.Sprint(v))
		}
		fmt.Println(strings.Join(head, delimiter))
		fmt.Println(strings.Join(body, delimiter))
	} else {
		for _, v := range ddbresult {
			body = append(body, fmt.Sprint(v))
		}
		fmt.Println(strings.Join(body, delimiter))
	}
}

func toJson(ddbresult map[string]interface{}) {
	jsonString, _ := json.Marshal(ddbresult)
	fmt.Println(string(jsonString))
}
