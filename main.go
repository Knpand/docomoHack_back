package main

import (
	"net/http"

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
	Result     bool     `json:"result"`
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

	e.Logger.Fatal(e.Start(":8080"))
}
