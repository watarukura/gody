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

var deleteOption gody.DeleteOption

func init() {
}

func NewCmdDel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete one record",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(&deleteOption)
		},
		Run: func(cmd *cobra.Command, args []string) {
			gody.Delete(&deleteOption, cmd)
		},
	}

	options := cmd.Flags()
	options.StringVarP(&deleteOption.TableName, "table", "T", "", "DynamoDB table name")
	options.StringVar(&deleteOption.PartitionKey, "pkey", "", "Partition Key")
	options.StringVar(&deleteOption.SortKey, "skey", "", "Sort Key")

	return cmd

}
