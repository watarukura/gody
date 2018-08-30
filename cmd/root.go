// Copyright © 2017 Wataru Kurashima
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gody",
		Short: "call aws-sdk-go DynamoDB from cli",
		Long:  `call aws-sdk-go DynamoDB from cli`,
	}
	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVar(&profile, "profile", "default", "AWS profile")
	cmd.PersistentFlags().StringVar(&region, "region", "ap-northeast-1", "AWS region")
	cmd.PersistentFlags().StringVar(&endpoint, "endpoint", "", "DynamoDB Local Endpoint URL")

	viper.BindPFlag("profile", cmd.PersistentFlags().Lookup("profile"))
	viper.BindPFlag("region", cmd.PersistentFlags().Lookup("region"))
	viper.BindPFlag("endpoint", cmd.PersistentFlags().Lookup("endpoint"))

	cmd.AddCommand(NewCmdList())
	cmd.AddCommand(NewCmdGet())
	cmd.AddCommand(NewCmdQuery())
	cmd.AddCommand(NewCmdScan())
	cmd.AddCommand(NewCmdPut())
	cmd.AddCommand(NewCmdDel())
	cmd.AddCommand(NewCmdDesc())

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd := NewCmdRoot()
	cmd.SetOutput(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}

var (
	profile  string
	region   string
	endpoint string
)

func init() {
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gody" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gody")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
