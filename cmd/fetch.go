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
	"sync"

	"github.com/fojtas98/dailyMenus/data"
	"github.com/fojtas98/dailyMenus/restaurant"
	"github.com/fojtas98/dailyMenus/scrapers"
	"github.com/spf13/cobra"
)

var byName bool

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch <name>",
	Short: "Gets daily menu from restaurants",
	Example: `  dailyMenus fetch myArea
	dailyMenus fetch myFavoriteRestaurant -r`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		rSlice := []restaurant.R{}
		var err error

		if !byName {
			rSlice, err = data.GetRestaurantsByArea(args[0])

		} else {
			rSlice, err = data.GetRestaurantsByRestaurant(args[0])

		}

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		for _, res := range rSlice {
			wg.Add(1)
			go func(r restaurant.R) {
				if r.ResType == "justToday" {

					scrapers.FirstDayMenu(r)
				} else {
					scrapers.AllWeekMenu(r)
				}
				wg.Done()
			}(res)

		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().BoolVarP(&byName, "restaruant", "r", false, "fetch restaurant for given name")
	// fetchCmd.Flags().MarkHidden("restaruant")
}
