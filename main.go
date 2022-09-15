package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type SearchRequest struct {
	CustomerId string  `json:"customer_id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type SearchResult struct {
	StoreId   string `json:"store_id"`
	StoreName string `json:"store_name"`
	Result    bool   `json:"result"`
}

func main() {
	e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})

	e.POST("/search", func(c echo.Context) error {
		request := new(SearchRequest)
		err := c.Bind(request)
		if err == nil {
			result := &SearchResult{
				//Result: search(db, request.CustomerId, request.Latitude, request.Longitude),
				StoreId:   "test_id",
				StoreName: "test_name",
				Result:    true,
			}

			return c.JSON(http.StatusOK, result)
		} else {
			return c.JSON(http.StatusBadRequest, &SearchResult{Result: false})
		}
	})

	e.Logger.Fatal(e.Start(":8080"))
}
