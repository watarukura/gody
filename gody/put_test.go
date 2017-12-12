package gody

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestPut(t *testing.T) {
	cmd := cobra.Command{}

	c1 := "jan,name,price\n4515438304003,茶こし共柄,500"
	curdir, _ := os.Getwd()
	tmpdir, _ := ioutil.TempDir(curdir, "put-test")
	defer os.RemoveAll(tmpdir)
	tmpc1 := filepath.Join(tmpdir, "put-test1.csv")
	ioutil.WriteFile(
		tmpc1,
		[]byte(c1),
		0777,
	)

	m1 := map[string]interface{}{"jan": "4515438304003", "name": "茶こし共柄", "price": 500}
	var marr1 []map[string]interface{}
	marr1 = append(marr1, m1)

	var putItemOption1 = PutItemOption{
		TableName: "item",
		Format:    "csv",
		File:      tmpc1,
	}

	cases := []struct {
		input *PutItemOption
		want  []map[string]interface{}
	}{
		{input: &putItemOption1, want: marr1},
	}

	for _, c := range cases {
		fmt.Printf("input %+v\n", c.input)
		get := buildAttribute(c.input, &cmd)

		fmt.Printf("get %+v\n", get)

		if reflect.DeepEqual(c.want[0], get[0]) {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}

}
