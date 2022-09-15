package main

import (
	"net/http"
	"log"
	"fmt"
	// "docomoHack_back/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	// "os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

)

type SearchRequest struct {
	CustomerId string  `json:"customer_id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type SearchResult struct {
	StoreIds   []string `json:"store_ids"`
	StoreNames []string `json:"store_names"`
	Result     bool     `json:"result"`
}

func sqlConnect() (database *gorm.DB, err error) {

	DBMS := "mysql"
	USER := "samp"
	PASS := "samp"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "db"

	log.Print(USER)
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
	

}


func connectDB() *gorm.DB {
	// DB接続
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Log: DBへの接続に成功しました in connectDB\n")
	}

	return db
}


func search(db *gorm.DB, CustomerId string, Latitude float64, Longitude float64) ([]string, []string, bool) {
	var (
		users      = make([]User, 0, 10)
		storeIds   = make([]string, 0, 10)
		storeNames = make([]string, 0, 10)
	)
	price := 800
	if db.Find(&users, "price = ?", price); len(users) == 0 {
		fmt.Printf("Error: 条件に一致するユーザがいません: %d in search\n", price)
		return []string{}, []string{}, false
	}

	for _, s := range users {
		storeIds = append(storeIds, s.UserId)
		storeNames = append(storeNames, s.UserName)
	}

	return storeIds, storeNames, true
}

type StoreIdAndStoreName struct {
	StoreId   string
	StoreName string
}
type User struct {
	UserId   string
	UserName string
	Price    int
}


func main() {
	e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	godotenv.Load()
	db := connectDB()
	
	e.POST("/search", func(c echo.Context) error {
		request := new(SearchRequest)
		err := c.Bind(request)
		if err == nil {
			//result := &SearchResult{
			//Result: search(db, request.CustomerId, request.Latitude, request.Longitude),
			//	StoreIds:   []string{"test_id"},
			//	StoreNames: []string{"test_name"},
			//	Result:    true,
			//}
			StoreIds, StoreNames, Result := search(db, request.CustomerId, request.Latitude, request.Longitude)
			result := &SearchResult{
				StoreIds:   StoreIds,
				StoreNames: StoreNames,
				Result:     Result,
			}

			return c.JSON(http.StatusOK, result)
		} else {
			return c.JSON(http.StatusBadRequest, &SearchResult{Result: false})
		}
	})

	e.GET("debug" , func(c echo.Context) error {
		request := new(SearchRequest)
		err := c.Bind(request)
		if err == nil {
			StoreIds, StoreNames, Result := search(db, request.CustomerId, request.Latitude, request.Longitude)
			result := &SearchResult{
				StoreIds:   StoreIds,
				StoreNames: StoreNames,
				Result:     Result,
			}

			return c.JSON(http.StatusOK, result)
		} else {
			return c.JSON(http.StatusBadRequest, &SearchResult{Result: false})
		}
	})

	e.Logger.Fatal(e.Start(":8080"))
}