package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestCmdGet(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "gody get --pkey aaa", want: "Error: Parameter error: TableName is required"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdRoot()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		cmd.Execute()

		get := buf.String()
		lineOne := strings.Split(get, "\n")[0]
		if c.want != lineOne {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
