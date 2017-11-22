package gody

import (
	"fmt"
	"strings"
)

var (
	head []string
	body []string
)

func Format(ddbresult map[string]interface{}, format string, header bool) {
	switch format {
	case "ssv":
		ssv(ddbresult, header)
		//case "csv":
		//	csv(ddbresult, header)
		//case "tsv":
		//	tsv(ddbresult, header)
		//case "json":
		//	json(ddbresult, header)
	}
}

func ssv(ddbresult map[string]interface{}, header bool) {
	if header {
		for k, v := range ddbresult {
			head = append(head, k)
			body = append(body, fmt.Sprintf("%s", v))
		}
		fmt.Println(strings.Join(head, " "))
		fmt.Println(strings.Join(body, " "))
	} else {
		for _, v := range ddbresult {
			body = append(body, fmt.Sprintf("%s", v))
		}
		fmt.Println(strings.Join(body, " "))
	}
}
