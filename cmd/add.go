package cmd

import (
	"github.com/fojtas98/dailyMenus/data"
	"github.com/fojtas98/dailyMenus/restaurant"
	"github.com/spf13/cobra"
)

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
		newRes := restaurant.R{}
		newRes.GetInfoFromUser()

		data.AddToRestaurants(newRes)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
