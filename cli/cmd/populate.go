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

// populateCmd represents the populate command
var populateCmd = &cobra.Command{
	Use:   "populate CIDR",
	Short: "Populate every address in a given CIDR",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		iu, err := cmd.Flags().GetBool("include-unusable")
		cobra.CheckErr(err)
		var ips []string

		for _, a := range args {
			c, err := cidrme.NewCIDR(a)
			cobra.CheckErr(err)
			var subips []string
			if iu {
				subips = c.GetIPs()
			} else {
				subips = c.GetUsableIPs()
			}
			ips = append(ips, subips...)
		}
		out.MustPrint(ips)
	},
}

func init() {
	rootCmd.AddCommand(populateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// populateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// populateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	populateCmd.Flags().BoolP("include-unusable", "u", false, "Also include unusable IPs in the range")
}
