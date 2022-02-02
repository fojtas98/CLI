package helpers

import (
	"strings"
)

func DeleteWeekDay(dish string) string {

	czWeekDays := []string{"pondělí", "úterý", "středa", "čtvrtek", "pátek"}

	for _, day := range czWeekDays {
		if strings.Contains(strings.ToLower(dish), day) {
			return ""
		}
	}
	return dish
}
