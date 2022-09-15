package main

import (
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

type SearchRequest struct {
	CustomerId string  `json:"customer_id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type SearchResult struct {
	StoreIds   []string `json:"store_ids"`
	StoreNames []string `json:"store_names"`
	Categories []string `json:"categories"`
	Prices     []int    `json:"prices"`
	Result     bool     `json:"result"`
}

type BookRequest struct {
	CustomerId string `json:"customer_id"`
	StoreId    string `json:"store_id"`
}

type BookResult struct {
	Result bool `json:"result"`
}

type CustomerSignupRequest struct {
	CustomerId   string    `json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	Gender       string    `json:"gender"`
	Address      string    `json:"address"`
	Birth        time.Time `json:"birth"`
	Role         string    `json:"rool"`
	Password     string    `json:"password"`
}

type CustomerSignupResult struct {
	Result bool `json:"result"`
}

type StoreSignupRequest struct {
	StoreId   string  `json:"store_id"`
	StoreName string  `json:"store_name"`
	Category  string  `json:"category"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Price     int     `json:"price"`
	Password  string  `json:"password"`
	Gender    string  `json:"gender"`
	MinAge    int     `json:"min_age"`
	MaxAge    int     `json:"max_age"`
	Role      string  `json:"role"`
}

type StoreSignupResult struct {
	Result bool `json:"result"`
}

type TermUpdateRequest struct {
	StoreId string `json:"store_id"`
	Gender  string `json:"gender"`
	MinAge  int    `json:"min_age"`
	MaxAge  int    `json:"max_age"`
	Role    string `json:"role"`
}

type TermUpdateResult struct {
	Result bool `json:"result"`
}

type LoginRequest struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type LoginResult struct {
	Result bool `json:"result"`
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
			StoreIds, StoreNames, Categories, Prices, Result := search(db, request.CustomerId, request.Latitude, request.Longitude)
			result := &SearchResult{
				StoreIds:   StoreIds,
				StoreNames: StoreNames,
				Categories: Categories,
				Prices:     Prices,
				Result:     Result,
			}

			return c.JSON(http.StatusOK, result)
		} else {
			return c.JSON(http.StatusBadRequest, &SearchResult{Result: false})
		}
	})

	e.POST("/book", func(c echo.Context) error {
		request := new(BookRequest)
		err := c.Bind(request)
		if err == nil {
			Result := book(db, request.CustomerId, request.StoreId)
			result := &BookResult{
				Result: Result,
			}

			return c.JSON(http.StatusOK, result)
		} else {
			return c.JSON(http.StatusBadRequest, &BookResult{Result: false})
		}
	})

	e.POST("/customersignup", func(c echo.Context) error {
		request := new(CustomerSignupRequest)
		err := c.Bind(request)
		if err == nil {
			Result := customerSignup(db, request.CustomerId, request.CustomerName, request.Gender, request.Address, request.Birth, request.Role, request.Password)
			result := &CustomerSignupResult{
				Result: Result,
			}

			return c.JSON(http.StatusOK, result)
		} else {
			return c.JSON(http.StatusBadRequest, &CustomerSignupResult{Result: false})
		}
	})

	e.POST("/storesignup", func(c echo.Context) error {
		request := new(StoreSignupRequest)
		err := c.Bind(request)
		if err == nil {
			Result := storeSignup(db, request.StoreId, request.StoreName, request.Category, request.Address, request.Latitude, request.Longitude, request.Price, request.Password, request.Gender, request.MinAge, request.MaxAge, request.Role)
			result := &StoreSignupResult{
				Result: Result,
			}

			return c.JSON(http.StatusOK, result)
		} else {
			return c.JSON(http.StatusBadRequest, &StoreSignupResult{Result: false})
		}
	})

	e.POST("/termupdate", func(c echo.Context) error {
		request := new(TermUpdateRequest)
		err := c.Bind(request)
		if err == nil {
			Result := termUpdate(db, request.StoreId, request.Gender, request.MinAge, request.MaxAge, request.Role)
			result := &TermUpdateResult{
				Result: Result,
			}

			return c.JSON(http.StatusOK, result)
		} else {
			return c.JSON(http.StatusBadRequest, &TermUpdateResult{Result: false})
		}
	})

	e.POST("/login", func(c echo.Context) error {
		request := new(LoginRequest)
		err := c.Bind(request)
		if err == nil {
			Result := login(db, request.UserId, request.Password)
			result := &LoginResult{
				Result: Result,
			}

			return c.JSON(http.StatusOK, result)
		} else {
			return c.JSON(http.StatusBadRequest, &LoginResult{Result: false})
		}
	})

	e.Logger.Fatal(e.Start(":8080"))
}
