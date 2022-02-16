package helpers

import (
	"strings"
)

func ContainsCzWeekDay(dish string) bool {

	czWeekDays := []string{"pondělí", "úterý", "středa", "čtvrtek", "pátek"}

	for _, day := range czWeekDays {
		if strings.Contains(strings.ToLower(dish), day) {
			return true
		}
	}
	return false
}
