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
	RootCmd.AddCommand(queryCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func queryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query",
		Short: "query table",
		Run: func(*cobra.Command, []string) {
			gody.Query(&queryOption)
		},
	}

	options := cmd.Flags()
	options.StringVarP(&queryOption.TableName, "table", "T", "", "DynamoDB table name")
	options.StringVar(&queryOption.PartitionKey, "pkey", "", "Partition Key")
	options.StringVar(&queryOption.SortKey, "skey", "", "Sort Key")
	options.StringVar(&queryOption.Format, "format", "ssv", "Output Format ssv|csv|tsv|json")
	options.BoolVar(&queryOption.Header, "header", false, "With Header")
	options.Int64Var(&queryOption.Limit, "limit", 100, "Output limit")
	options.StringVar(&queryOption.Index, "index", "", "GSI Name")
	options.BoolVar(&queryOption.Eq, "eq", true, "GSI Query Parameter EQ")
	options.BoolVar(&queryOption.Lt, "lt", false, "GSI Query Parameter LT")
	options.BoolVar(&queryOption.Le, "le", false, "GSI Query Parameter LE")
	options.BoolVar(&queryOption.Gt, "gt", false, "GSI Query Parameter GT")
	options.BoolVar(&queryOption.Ge, "ge", false, "GSI Query Parameter GE")

	return cmd

}
