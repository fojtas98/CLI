package scrapers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fojtas98/dailyMenus/restaurant"
	"github.com/fojtas98/dailyMenus/scrapers/helpers"
)

func AllWeekMenu(r restaurant.R) {
	skip := int(time.Now().Weekday())
	meal := ""
	menu := []string{""}
	menu = append(menu, "### \033[1m"+r.Name+"\033[0m ###")
	res, err := http.Get(r.Url)
	dataInBytes, _ := ioutil.ReadAll(res.Body)
	pageContent := string(dataInBytes)
	if err != nil {
		log.Fatal(err)
	}
	if r.ParentTag != "" {
		parentTagStartsAt := strings.Index(pageContent, r.ParentTag)
		pageContent = pageContent[parentTagStartsAt:]
	}
	for skip > 0 {
		for i := 0; i < r.Meals; i++ {
			dishInexStart := strings.Index(pageContent, r.OpenTag)
			if dishInexStart == -1 {
				menu = append(menu, " No menu found for today")
				break
			}

			dishInexStart += len(r.OpenTag)
			pageContent = pageContent[dishInexStart:]

			dishIndexEnd := strings.Index(pageContent, r.CloseTag)
			if dishIndexEnd == -1 {
				menu = append(menu, " close tag is not found please create new instace for this rtaurant with right close tag")
				break
			}
			meal = pageContent[:dishIndexEnd]

			if len(meal) <= 3 || helpers.ContainsCzWeekDay(meal) {
				i--
			} else {
				fmt.Println(len(meal))
				if skip == 1 {
					meal = strings.TrimSpace(meal)
					meal = helpers.DeleteTags(meal)
					menu = append(menu, " "+meal)
				}
			}
		}
		skip -= 1
	}
	for _, meal := range menu {
		fmt.Println(meal)
	}
}
