package gody

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	m := map[string]interface{}{"jan": "4937751121103", "name": "つぼキーク", "price": 2000}
	var marr []map[string]interface{}
	marr = append(marr, m)
	cmd := new(cobra.Command)

	var formatTarget1 = FormatTarget{
		ddbresult: marr,
		header:    true,
		format:    "ssv",
		fields:    []string{"name", "jan"},
		cmd:       cmd,
	}

	var formatTarget2 = FormatTarget{
		ddbresult: marr,
		header:    true,
		format:    "ssv",
		fields:    []string{},
		cmd:       cmd,
	}

	var formatTarget3 = FormatTarget{
		ddbresult: marr,
		header:    true,
		format:    "json",
		fields:    []string{},
		cmd:       cmd,
	}

	var formatTarget4 = FormatTarget{
		ddbresult: marr,
		header:    true,
		format:    "json",
		fields:    []string{"name", "jan"},
		cmd:       cmd,
	}

	cases := []struct {
		input FormatTarget
		want  string
	}{
		{input: formatTarget1, want: "name jan"},
		{input: formatTarget2, want: "jan name price"},
		{input: formatTarget3, want: "[{\"jan\":\"4937751121103\",\"name\":\"つぼキーク\",\"price\":2000}]"},
		{input: formatTarget4, want: "[{\"jan\":\"4937751121103\",\"name\":\"つぼキーク\"}]"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd.SetOutput(buf)
		fmt.Printf("input %+v\n", c.input)
		Format(c.input)

		get := buf.String()
		lineOne := strings.Split(get, "\n")[0]
		if c.want != lineOne {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
