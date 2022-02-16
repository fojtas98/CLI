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
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		FGGreen := promptui.BGRed

		promptui.Styler(FGGreen)("✔")

		prompt := promptui.Prompt{
			Label: "Name",
		}
		result, _ := prompt.Run()

		fmt.Printf("You choose %q\n", result)
		prompt2 := promptui.Select{
			Label: "Select what do you want to update",
			Items: []string{"Website", "Name", "Open tag", "Close Tag", "Number of meals in menu", "Parent tag", "Area"},
			Size:  7,
		}

		_, result2, err := prompt2.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result2)
		// 	if args[0] == "" {
		// 		fmt.Println("name argument is needed")
		// 		os.Exit(0)
		// 	}
		// 	err := data.UpdateRestaurantByName(args[0])
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		os.Exit(0)
		// 	}
		// 	fmt.Println(args[0] + " was deleted secesffuly")
		// },
	}}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
