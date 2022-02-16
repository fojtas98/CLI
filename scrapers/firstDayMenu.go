package scrapers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fojtas98/dailyMenus/restaurant"
	"github.com/fojtas98/dailyMenus/scrapers/helpers"
)

func FirstDayMenu(r restaurant.R) {
	menu := []string{""}
	response, err := http.Get(r.Url)
	if err != nil {
		log.Fatal(err)
	}
	dataInBytes, _ := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)
	if r.ParentTag != "" {
		parentTagStartsAt := strings.Index(pageContent, r.ParentTag)
		pageContent = pageContent[parentTagStartsAt:]
	}
	menu = append(menu, "### \033[1m"+r.Name+"\033[0m ###")
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
			menu = append(menu, " close tag is not found please create new instace for this restaurant with right close tag")
			break
		}
		meal := pageContent[:dishIndexEnd]
		meal = strings.TrimSpace(meal)
		meal = helpers.DeleteTags(meal)
		menu = append(menu, " "+meal)
	}
	for _, meal := range menu {
		fmt.Println(meal)
	}
}
