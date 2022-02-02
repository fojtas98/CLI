package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/fojtas98/dailyMenus/data"
	"github.com/fojtas98/dailyMenus/helpers"
	"github.com/spf13/cobra"
)

var newRestaurant helpers.Restaurant

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Commad to add your restaurant",
	Long: "\033[1mTo add your restaurants menu you will need\033[0m \n" + ` 1. Url where the menu lives
 2. Name of the restaurant
 3. Html open tag(<ti>) before the meal
 4. Html clse tag </ti> after the meal
 5. Area where you can find the restaurant -> home
 6. resType, id the url contains menus for all weak the you select allWekk, if you can see just menu fot today you select justToday`,
	Run: func(cmd *cobra.Command, args []string) {
		newRestaurant.Url = getInfoFromUser("Website")
		newRestaurant.Name = getInfoFromUser("Name")
		newRestaurant.OpenTag = getInfoFromUser("OpenTag")
		newRestaurant.CloseTag = getInfoFromUser("CloseTag")
		newRestaurant.Meals, _ = strconv.Atoi(getInfoFromUser("DishCount"))
		newRestaurant.Area = getInfoFromUser("Area")
		newRestaurant.ParentTag = getInfoFromUser("ParentTag(optional)")
		newRestaurant.ResType = getInfoFromUser("Type -> allWeek/justToday")

		data.AddToRestaurants(newRestaurant)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func getInfoFromUser(name string) string {
	var line string
	for {
		fmt.Print("Enter " + name + ": ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line = scanner.Text()
		if name == "DishCount" {
			v, _ := strconv.Atoi(line)
			if v == 0 {
				fmt.Println(name + "dish count needs to be number")
			} else {
				break

			}
		} else {

			if name == "Type -> allWeek/justToday" && (line != "allWeek" && line != "justToday") {
				fmt.Println(name + "type needs to be allWeek or  justToday")
				continue

			}

			if line != "" {
				break
			}
			fmt.Println(name + "needs to be specified")
		}
	}
	return line
}
