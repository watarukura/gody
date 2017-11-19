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
	"log"
)

var getOption gody.GetOption

func init() {
	RootCmd.AddCommand(getCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getCmd() *cobra.Command {
	svc, err := newService()
	if err != nil {
		log.Fatal("create service failed.")
	}

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get one record",
		Run: func(*cobra.Command, []string) {
			gody.Get(svc, &getOption)
		},
	}

	options := cmd.Flags()
	options.StringVarP(&getOption.TableName, "table", "T", "", "DynamoDB table name")
	options.StringVar(&getOption.HashKey, "hash", "", "Hash Key")
	options.StringVar(&getOption.RangeKey, "range", "", "Range Key")

	return cmd

}
