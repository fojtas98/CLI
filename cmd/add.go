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

var addToMap struct {
	Funcs string
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		newRestaurant.Url = getInfoFromUser("Website")
		newRestaurant.Name = getInfoFromUser("Name")
		newRestaurant.OpenTag = getInfoFromUser("OpenTag")
		newRestaurant.CloseTag = getInfoFromUser("CloseTag")
		newRestaurant.Meals, _ = strconv.Atoi(getInfoFromUser("DishCount"))
		newRestaurant.Area = getInfoFromUser("Area")
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
