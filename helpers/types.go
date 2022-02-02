package helpers

type Restaurant struct {
	Name      string
	Url       string
	ResType   string
	Meals     int
	OpenTag   string
	CloseTag  string
	ParentTag string
	Area      string
}

type TodaysMenu []string

type Restaurants []Restaurant
