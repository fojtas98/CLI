package restaurant

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func (r *R) GetInfoFromUser() {
	r.addWebSite().
		addCountOfDishes().
		addName().
		addOpenTag().
		addCloseTag().
		addArea().
		addType().
		addParentTag()
}

func (r *R) addWebSite() *R {
	fmt.Printf("enter Website: ")
	r.Url = readLine(false)
	fmt.Println("---------------------")
	return r
}
func (r *R) addName() *R {
	fmt.Printf("enter Name: ")
	r.Name = readLine(false)
	fmt.Println("---------------------")
	return r
}
func (r *R) addOpenTag() *R {
	fmt.Printf("enter Open tag: ")
	r.OpenTag = readLine(false)
	fmt.Println("---------------------")
	return r
}
func (r *R) addCloseTag() *R {
	fmt.Printf("enter close tag: ")

	r.CloseTag = readLine(false)
	fmt.Println("---------------------")
	return r
}
func (r *R) addArea() *R {
	fmt.Printf("enter Area: ")

	r.Area = readLine(false)
	fmt.Println("---------------------")
	return r
}
func (r *R) addType() *R {
	fmt.Printf("enter Type(must be justToday or allWeek): ")

	for {
		temp := readLine(false)
		if temp == "justTodat" || temp == "allWeek" {
			r.ResType = temp
			fmt.Println("---------------------")
			break

		}
		fmt.Println("type must be justToday or allWeek")
	}

	return r
}
func (r *R) addCountOfDishes() *R {
	fmt.Printf("enter count of dishes in menu: ")
	for {
		v, err := strconv.Atoi(readLine(false))
		if err == nil && v > 0 {
			r.Meals = v
			fmt.Println("---------------------")
			break
		}

		fmt.Println("this filed needs to be a number bigger then 0")
	}

	return r
}
func (r *R) addParentTag() *R {
	fmt.Printf("enter parent tag(OPTIONAL): ")
	r.ParentTag = readLine(true)
	fmt.Println("---------------------")
	return r
}

func readLine(opt bool) (line string) {
	for {
		scanner.Scan()
		line = scanner.Text()
		if line != "" || opt {
			break
		}
		fmt.Println("!this filed is not optional!")

	}
	return
}
