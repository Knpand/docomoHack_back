package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type StoreIdAndStoreName struct {
	StoreId   string
	StoreName string
}
type User struct {
	UserId   string
	UserName string
	Price    int
}

func sqlConnect() (database *gorm.DB, err error) {
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("USER")
	PASS := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DBNAME")

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