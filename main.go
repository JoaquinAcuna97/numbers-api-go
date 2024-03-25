package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type number struct {
	Number   int    `json:number`
	Category string `json:category`
}

var numbers = []number{}

func getNumbers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, numbers)
}

func postNumbers(context *gin.Context) {
	var input struct {
		Number int `json:"number" binding:"required"`
	}
	err := context.BindJSON(&input)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "There was an error processing the request!"})
	} else {
		var newNumber number = auxCalculateCategory(input.Number)
		numbers = append(numbers, newNumber)
		context.Writer.Header().Set("Content-Type", "application/json")
		context.JSON(http.StatusCreated, numbers)
	}

}

func auxCalculateCategory(input int) number {
	// Core logic: Utilizing only one if statement
	divisibilityTypes := map[int]string{
		3:  "Type 1",
		5:  "Type 2",
		15: "Type 3",
	}

	var result string

	for divisor, typeName := range divisibilityTypes {
		if input%divisor == 0 {
			result = typeName
		}
	}
	if result == "" {
		result = strconv.Itoa(input)
	}

	var newNumber = number{Number: input, Category: result}
	return newNumber
}

func getCategoryOfNumber(context *gin.Context) {
	number, err := strconv.Atoi(context.Param("number"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't convert this to an int!"})
	} else {
		for _, n := range numbers {
			if n.Number == number {
				context.JSON(http.StatusOK, n)
				return

			}
		}
		context.JSON(http.StatusNotFound, gin.H{"message": "number not found"})
	}
}

func main() {

	router := gin.Default()
	router.GET("/numbers", getNumbers)
	router.POST("/numbers", postNumbers)
	router.GET("/numbers/:number", getCategoryOfNumber)
	router.Run("localhost:8080")
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
