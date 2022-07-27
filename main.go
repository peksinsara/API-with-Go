package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type car struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Quantity int    `json:"quantity"`
}

var cars = []car{
	{ID: "1", Name: "Mercedes Benz G Class", Year: 2021, Quantity: 2},
	{ID: "2", Name: "Audi A3 Sportsback", Year: 2019, Quantity: 5},
	{ID: "3", Name: "Brabus", Year: 2020, Quantity: 6},
}

//returning JSON version of all cars
func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

func addCar(c *gin.Context) {
	var newCar car
	if err := c.BindJSON(&newCar); err != nil {
		return
	}
	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, newCar)
}

func main() {
	router := gin.Default()
	router.GET("/cars", getCars)
	router.POST("/cars", addCar)
	router.Run("localhost:8080")
}
