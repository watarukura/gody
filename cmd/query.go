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

var queryOption gody.QueryOption

func init() {
}

func NewCmdQuery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query",
		Short: "query table",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(&queryOption)
		},
		Run: func(cmd *cobra.Command, args []string) {
			gody.Query(&queryOption, cmd)
		},
	}

	options := cmd.Flags()
	options.StringVarP(&queryOption.TableName, "table", "T", "", "DynamoDB table name")
	options.StringVar(&queryOption.PartitionKey, "pkey", "", "Partition Key")
	options.StringVar(&queryOption.SortKey, "skey", "", "Sort Key")
	options.StringVar(&queryOption.Format, "format", "ssv", "Output Format ssv|csv|tsv|json")
	options.BoolVar(&queryOption.Header, "header", false, "With Header")
	options.Int64Var(&queryOption.Limit, "limit", 0, "Output limit (0 means all)")
	options.StringVar(&queryOption.Index, "index", "", "GSI Name")
	options.BoolVar(&queryOption.Eq, "eq", false, "RangeKey Query Parameter EQ")
	options.BoolVar(&queryOption.Lt, "lt", false, "RangeKey Query Parameter LT")
	options.BoolVar(&queryOption.Le, "le", false, "RangeKey Query Parameter LE")
	options.BoolVar(&queryOption.Gt, "gt", false, "RangeKey Query Parameter GT")
	options.BoolVar(&queryOption.Ge, "ge", false, "RangeKey Query Parameter GE")
	options.StringVar(&queryOption.Field, "field", "", "Select Fields comma separated ex)field1,field2...")

	return cmd

}
