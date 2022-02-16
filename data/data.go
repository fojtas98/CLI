package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/fojtas98/dailyMenus/restaurant"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {

	_, b, _, _ := runtime.Caller(0)
	path := path.Dir(filepath.Dir(b))
	var err error
	if _, err := os.Stat(path + "/sqlite-database.db"); err != nil {
		os.Chmod(path, 0777)
		f, err := os.Create(path + "/sqlite-database.db")
		f.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	db, err = sql.Open("sqlite3", path+"/sqlite-database.db")

	if err != nil {
		return err
	}
	CreateTable()
	return db.Ping()
}

func CreateTable() {

	err := db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	createRestaurantsTable := `CREATE TABLE IF NOT EXISTS restaurants(
		"restaurantId" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"url" TEXT,
		"type" TEXT,
		"numberOfMealsInMenu" INTEGER,
		"openTag" TEXT,
		"closeTag" TEXT,
		"parentTag" TEXT,
		"area" TEXT 
	  );`

	resStatement, err := db.Prepare(createRestaurantsTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	resStatement.Exec()
}

func AddToRestaurants(restaurant restaurant.R) {
	insertNoteSQL := "INSERT INTO restaurants(name, url ,type ,numberOfMealsInMenu, openTag, closeTag ,parentTag , area) VALUES (?, ?, ?, ? ,?, ?, ?, ?)"
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

func GetRestaurantsByArea(area string) ([]restaurant.R, error) {
	rSlice := []restaurant.R{}

	rows, err := db.Query("SELECT * from restaurants WHERE LOWER( restaurants.area )= ?", strings.ToLower(area))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var r restaurant.R
		if err := rows.Scan(&id, &r.Name, &r.Url, &r.ResType,
			&r.Meals, &r.OpenTag, &r.CloseTag, &r.ParentTag, &r.Area); err != nil {
			return rSlice, err
		}
		rSlice = append(rSlice, r)
	}
	if err = rows.Err(); err != nil {
		return rSlice, err

	}
	if len(rSlice) == 0 {
		return rSlice, fmt.Errorf("this area dont have any restaurants, try different area")
	}
	return rSlice, nil
}
func GetRestaurantsByRestaurant(name string) ([]restaurant.R, error) {
	sliceR := []restaurant.R{}
	rows, err := db.Query("	SELECT * FROM restaurants WHERE LOWER( restaurants.name ) = ?", strings.ToLower(name))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var r restaurant.R
		if err := rows.Scan(&id, &r.Name, &r.Url, &r.ResType,
			&r.Meals, &r.OpenTag, &r.CloseTag, &r.Area); err != nil {
			return sliceR, err
		}
		sliceR = append(sliceR, r)
	}
	if err = rows.Err(); err != nil {
		return sliceR, err

	}
	if len(sliceR) == 0 {
		return sliceR, fmt.Errorf("this area dont have any restaurants, try different area")
	}
	return sliceR, nil
}

func DeleteRestaurantByName(name string) error {
	insertNoteSQL := "DELETE FROM restaurants WHERE LOWER( restaurants.name ) = ?"

	result, err := db.Exec(insertNoteSQL, strings.ToLower(name))
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("restaurant with this name wasnt found")
	}

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func UpdateRestaurantByName(name string) error {
	insertNoteSQL := "UPDATE restaurants WHERE LOWER( restaurants.name ) = ?"

	result, err := db.Exec(insertNoteSQL, strings.ToLower(name))
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("restaurant with this name wasnt found")
	}

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
