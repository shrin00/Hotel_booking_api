package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Table struct {
	ID         int    `json:"id"`
	Seats      int    `json:"seats"`
	Booked     bool   `json:"booked"`
	ReservedBy Person `json:"reserved_by"`
}

var tables = []Table{
	{ID: 1, Seats: 4},
	{ID: 2, Seats: 4},
	{ID: 3, Seats: 6},
	{ID: 4, Seats: 6},
	{ID: 5, Seats: 8},
	{ID: 6, Seats: 8},
}

func getTables(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tables)
}

func getAvailableTables(c *gin.Context) {
	var availableTables []Table
	for i := 0; i < len(tables); i++ {
		table := tables[i]
		if !table.Booked {
			availableTables = append(availableTables, table)
		}
	}

	c.IndentedJSON(http.StatusOK, availableTables)
}

func getBookedTables(c *gin.Context) {
	var bookedTables []Table
	for i := 0; i < len(tables); i++ {
		table := tables[i]
		if table.Booked {
			bookedTables = append(bookedTables, table)
		}
	}
	c.IndentedJSON(http.StatusOK, bookedTables)
}

func bookTable(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var person Person

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Sorry!! Not a number."})
		return
	}

	if err := c.BindJSON(&person); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No person"})
		return
	}

	for i := 0; i < len(tables); i++ {
		table := tables[i]
		if table.ID == id {
			if table.Booked {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Sorry table is already booked!"})
				return
			}

			table.Booked = true
			table.ReservedBy = person

			tables[i] = table

			c.IndentedJSON(http.StatusOK, table)
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Sorry!! Not found."})
}

func addTable(c *gin.Context) {
	var newTable Table

	if err := c.BindJSON(&newTable); err != nil {
		return
	}

	tables = append(tables, newTable)
	c.IndentedJSON(http.StatusOK, newTable)
}

func getTableInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Sorry!! Not a number."})
		return
	}

	for i := 0; i < len(tables); i++ {
		table := tables[i]
		if table.ID == id {
			c.IndentedJSON(http.StatusOK, table)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Sorry!! Not found."})
}

func main() {
	route := gin.Default()
	route.GET("/tables", getTables)
	route.GET("/available_tables", getAvailableTables)
	route.GET("/booked_tables", getBookedTables)
	route.POST("/book/:id", bookTable)
	route.POST("/tables", addTable)
	route.GET("/tables/:id", getTableInfo)

	route.Run("localhost:8080")
}