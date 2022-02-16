package helpers

import (
	"strings"
)

func DeleteTags(dish string) string {

	for {
		tagIndexStart := strings.Index(dish, "<")
		if tagIndexStart == -1 {
			return dish
		}
		tagIndexEnd := strings.Index(dish, ">")
		if tagIndexEnd == -1 {
			return dish
		}
		tagToBeRemoved := dish[tagIndexStart : tagIndexEnd+1]
		dish = strings.Replace(dish, tagToBeRemoved, "", 1)
	}

}
