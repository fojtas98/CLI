/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"os"

	"github.com/fojtas98/dailyMenus/data"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete restaurant by name",
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			fmt.Println("name argument is needed")
			os.Exit(0)
		}
		err := data.DeleteRestaurantByName(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		fmt.Println(args[0] + " was deleted secesffuly")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
