package gody

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestFormat(t *testing.T) {
	marr := getTestFormatTargetDdbResult()
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

	var formatTarget5 = FormatTarget{
		ddbresult: marr,
		header:    true,
		format:    "json",
		fields:    []string{"name", "jan", "dummy"},
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
		{input: formatTarget5, want: "[{\"dummy\":\"_\",\"jan\":\"4937751121103\",\"name\":\"つぼキーク\"}]"},
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

func getTestFormatTargetDdbResult() []map[string]interface{} {
	var ddbResult []map[string]interface{}
	m := map[string]interface{}{"jan": "4937751121103", "name": "つぼキーク", "price": 2000}
	ddbResult = append(ddbResult, m)
	return ddbResult
}

func getTestFormatTarget(ddbResult []map[string]interface{}) FormatTarget {
	cmd := new(cobra.Command)
	target := FormatTarget{
		ddbresult: ddbResult,
		header:    true,
		format:    "json",
		fields:    []string{"name", "jan"},
		cmd:       cmd,
	}
	return target
}

var testCreateBodyExpectedResult = [][]string{{"つぼキーク", "4937751121103"}}

func Test_createBody(t *testing.T) {
	target := getTestFormatTarget(getTestFormatTargetDdbResult())
	head := []string{"name", "jan"}
	body := createBody(target, head)
	if !reflect.DeepEqual(body, testCreateBodyExpectedResult) {
		t.Errorf("expected: %s, actual: %s", testCreateBodyExpectedResult, body)
	}
}

func Test_deprecatedCreateBody(t *testing.T) {
	target := getTestFormatTarget(getTestFormatTargetDdbResult())
	head := []string{"name", "jan"}
	body := deprecatedCreateBody(target, head)
	if !reflect.DeepEqual(body, testCreateBodyExpectedResult) {
		t.Errorf("expected: %s, actual: %s", testCreateBodyExpectedResult, body)
	}
}

func Benchmark_deprecatedCreateBody(b *testing.B) {
	target := getTestFormatTarget(getTestFormatTargetDdbResult())
	head := []string{"name", "jan"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deprecatedCreateBody(target, head)
	}
}

func Benchmark_createBody(b *testing.B) {
	target := getTestFormatTarget(getTestFormatTargetDdbResult())
	head := []string{"name", "jan"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createBody(target, head)
	}
}
