package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	UserId   string
	UserName string
	Price    int
}

type Customer struct {
	CustomerId   string
	Password     string
	Gender       string
	CustomerName string
	Address      string
	Birth        time.Time
	Rool         string
}

type Store struct {
	StoreId   string
	Password  string
	StoreName string
	Category  string
	Latitude  float64
	Longitude float64
	Address   string
	Price     int
}

type Term struct {
	StoreId string
	Gender  string
	MinAge  int
	MaxAge  int
	Role    string
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

func book(db *gorm.DB, CustomerId string, StoreId string) bool {
	return true
}

func customerSignup(db *gorm.DB, CustomerId string, CustomerName string, Gender string, Address string, Birth time.Time, Rool string, Password string) bool {
	var customer Customer

	if d_err := db.First(&customer, "customer_id = ?", CustomerId).Error; d_err != nil {
		customer = Customer{
			CustomerId:   CustomerId,
			CustomerName: CustomerName,
			Gender:       Gender,
			Address:      Address,
			Birth:        Birth,
			Rool:         Rool,
			Password:     Password,
		}
		if customer_err := db.Create(&customer).Error; customer_err == nil {
			fmt.Printf("Log: 顧客情報を登録: %s in customerSignup\n", CustomerId)
			return true
		} else {
			fmt.Printf("Error: 顧客情報の登録に失敗しました: %s in customerSignup\n", CustomerId)
			return false
		}
	} else {
		fmt.Printf("Error: そのIDは既に使われています: %s in customerSignup\n", CustomerId)
		return false
	}
}

func storeSignup(db *gorm.DB, StoreId string, StoreName string, Category string, Address string, Latitude float64, Longitude float64, Price int, Password string, Gender string, MinAge int, MaxAge int, Role string) bool {
	var (
		store Store
		term  Term
	)

	if d_err := db.First(&store, "store_id = ?", StoreId).Error; d_err != nil {
		store = Store{
			StoreId:   StoreId,
			Password:  Password,
			StoreName: StoreName,
			Category:  Category,
			Latitude:  Latitude,
			Longitude: Longitude,
			Address:   Address,
			Price:     Price,
		}
		if store_err := db.Create(&store).Error; store_err == nil {
			fmt.Printf("Log: 飲食店の情報を登録: %s in storeSignup\n", StoreId)
			term = Term{
				StoreId: StoreId,
				Gender:  Gender,
				MinAge:  MinAge,
				MaxAge:  MaxAge,
				Role:    Role,
			}
			if term_err := db.Create(&term).Error; term_err == nil {
				fmt.Printf("Log: 飲食店の顧客条件情報を登録: %s in storeSignup\n", StoreId)
				return true
			} else {
				fmt.Printf("Error: 飲食店の顧客条件情報の登録に失敗しました: %s in storeSignup\n", StoreId)
				return false
			}
		} else {
			fmt.Printf("Error: 飲食店の情報の登録に失敗しました: %s in storeSignup\n", StoreId)
			return false
		}
	} else {
		fmt.Printf("Error: そのIDは既に使われています: %s in storeSignup\n", StoreId)
		return false
	}
}

func termUpdate(db *gorm.DB, StoreId string, Gender string, MinAge int, MaxAge int, Role string) bool {
	var (
		term Term
	)

	if err := db.Model(&term).Where("store_id = ?", StoreId).Updates(map[string]interface{}{"gender": Gender, "min_age": MinAge, "max_age": MaxAge, "role": Role}).Error; err != nil {
		fmt.Printf("Error:  飲食店の顧客条件情報の更新に失敗しました: %s in termUpdate\n", StoreId)
		return false
	}

	fmt.Printf("Log: 飲食店の顧客条件情報を更新: %s in termUpdate\n", StoreId)
	return true
}

func login(db *gorm.DB, UserId string, Password string) bool {
	var customer Customer
	if customer_err := db.First(&customer, "customer_id = ? AND password = ?", UserId, Password).Error; customer_err != nil {
		var store Store
		if store_err := db.First(&store, "store_id = ? AND password = ?", UserId, Password).Error; store_err != nil {
			fmt.Printf("Error: ログインに失敗しました: %s in login\n", UserId)
			return false
		}
	}

	return true
}
