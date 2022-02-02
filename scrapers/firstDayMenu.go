package scrapers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fojtas98/dailyMenus/helpers"
)

func FirstDayMenu(res helpers.Restaurant) {
	var result helpers.TodaysMenu
	response, err := http.Get(res.Url)
	if err != nil {
		log.Fatal(err)
	}
	dataInBytes, _ := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)
	if res.ParentTag != "" {
		parentTagStartsAt := strings.Index(pageContent, res.ParentTag)
		pageContent = pageContent[parentTagStartsAt:]
	}
	result = append(result, "### \033[1m"+res.Name+"\033[0m ###")
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
		dish := pageContent[:dishIndexEnd]
		dish = strings.TrimSpace(dish)
		dish = helpers.DeleteTags(dish)
		result = append(result, " "+dish)
	}
	for _, meal := range result {
		fmt.Println(meal)
	}
}
