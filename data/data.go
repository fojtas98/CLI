package data

import (
	"database/sql"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/fojtas98/dailyMenus/helpers"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	_, b, _, _ := runtime.Caller(0)
	path := path.Dir(filepath.Dir(b))

	var err error
	db, err = sql.Open("sqlite3", path+"/sqlite-database")
	if err != nil {
		return err
	}
	return db.Ping()
}

func CreateTable() {

	err := db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	createRestaurantsTable := `CREATE TABLE restaurants (
		"restaurantId" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"url" TEXT,
		"type" TEXT,
		"numberOfMealsInMenu" INTEGER,
		"openTag" TEXT,
		"closeTag" TEXT,
		"area" TEXT 
	  );`

	resStatement, err := db.Prepare(createRestaurantsTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	resStatement.Exec()
}

func AddToRestaurants(restaurant helpers.Restaurant) {
	insertNoteSQL := "INSERT INTO restaurants(name,url,type,numberOfMealsInMenu, openTag, closeTag, area) VALUES (?, ?, ?, ? ,?, ?, ?)"
	rv := reflect.ValueOf(restaurant)
	var args []interface{}
	for i := 0; i < rv.NumField(); i++ {
		args = append(args, rv.Field(i).Interface())
	}
	_, err := db.Exec(insertNoteSQL, args...)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Restaurant with name " + restaurant.Name + " has been successfully added")
}

func GetRestaurantsByArea(area string) (helpers.Restaurants, error) {
	var restaurants helpers.Restaurants
	rows, err := db.Query("SELECT * from restaurants where area = ?", area)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var restaurant helpers.Restaurant
		if err := rows.Scan(&id, &restaurant.Name, &restaurant.Url, &restaurant.ResType,
			&restaurant.Meals, &restaurant.OpenTag, &restaurant.CloseTag, &restaurant.Area); err != nil {
			return restaurants, err
		}
		restaurants = append(restaurants, restaurant)
	}
	if err = rows.Err(); err != nil {
		return restaurants, err

	}
	if len(restaurants) == 0 {
		return restaurants, fmt.Errorf("this area dont have any restaurants, try different area")
	}
	return restaurants, nil
}

func DeleteRestaurantByName(name string) error {
	insertNoteSQL := "DELETE FROM restaurants where name = ?"

	result, err := db.Exec(insertNoteSQL, name)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("restaurant with this name wasnt found")
	}

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
