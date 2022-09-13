/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/drewstinnett/cidrme/cidrme"
	"github.com/spf13/cobra"
)

// containsCmd represents the contains command
var containsCmd = &cobra.Command{
	Use:   "contains NETWORK IP",
	Short: "Check if an IP is included in a subnet",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cidrme.ContainsWithStrings(args[0], args[1])
		cobra.CheckErr(err)
		out.MustPrint(res)
	},
}

func init() {
	rootCmd.AddCommand(containsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// containsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// containsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
