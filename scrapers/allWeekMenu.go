package scrapers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fojtas98/dailyMenus/helpers"
)

func AllWeekMenu(res helpers.Restaurant) {
	skip := int(time.Now().Weekday())
	var dish string
	var result helpers.TodaysMenu
	result = append(result, "### \033[1m"+res.Name+"\033[0m ###")
	response, err := http.Get(res.Url)
	dataInBytes, _ := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)
	if err != nil {
		log.Fatal(err)
	}
	if res.ParentTag != "" {
		parentTagStartsAt := strings.Index(pageContent, res.ParentTag)
		pageContent = pageContent[parentTagStartsAt:]
	}
	for skip > 0 {
		for i := 0; i < res.Meals; i++ {
			dishInexStart := strings.Index(pageContent, res.OpenTag)
			if dishInexStart == -1 {
				result = append(result, " No menu found for today")
				break
			}

			dishInexStart += len(res.OpenTag)
			pageContent = pageContent[dishInexStart:]

			dishIndexEnd := strings.Index(pageContent, res.CloseTag)
			if dishIndexEnd == -1 {
				result = append(result, " close tag is not found please create new instace for this restaurant with right close tag")
				break
			}
			dish = pageContent[:dishIndexEnd]
			dish = helpers.DeleteWeekDay(dish)
			if len(dish) == 0 {
				i--
			} else {
				if skip == 1 {
					dish = strings.TrimSpace(dish)
					dish = helpers.DeleteTags(dish)
					result = append(result, " "+dish)
				}
			}

		}
		skip -= 1
	}
	for _, meal := range result {
		fmt.Println(meal)
	}
}
