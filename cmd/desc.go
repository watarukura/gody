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

var descOption gody.DescOption

func init() {
}

func NewCmdDesc() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "desc",
		Short: "Describe table",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(&descOption)
		},
		Run: func(cmd *cobra.Command, args []string) {
			gody.Desc(&descOption, cmd)
		},
	}

	options := cmd.Flags()
	options.StringVarP(&descOption.TableName, "table", "T", "", "DynamoDB table name")
	options.StringVar(&descOption.Format, "format", "ssv", "Output Format ssv|csv|tsv|json")
	options.BoolVar(&descOption.Header, "header", false, "With Header")
	options.StringVar(&descOption.Field, "field", "", "Select Fields comma separated ex)field1,field2...")

	return cmd

}
