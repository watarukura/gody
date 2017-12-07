// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/watarukura/gody/gody"
)

var putOption gody.PutItemOption

func init() {
}

func NewCmdPut() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put",
		Short: "Put record(s)",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(&putOption)
		},
		Run: func(cmd *cobra.Command, args []string) {
			gody.Put(&putOption, cmd)
		},
	}

	options := cmd.Flags()
	options.StringVarP(&putOption.TableName, "table", "T", "", "DynamoDB table name")
	options.StringVar(&putOption.Format, "format", "ssv", "Input Format ssv|csv|tsv|json")
	options.StringVar(&putOption.File, "file", "", "Path to input File (optional; default is stdin)")

	return cmd

}
