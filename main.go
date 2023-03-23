package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine
var connectionString = "host = 127.0.0.1 port=5432 user = postgres password = postgres dbname  = WebAppDB sslmode = disable"
var connection *sql.DB

func main() {
	var e error
	connection, e = sql.Open("postgres", connectionString)
	if e != nil {
		fmt.Println(e)
		return
	}

	router = gin.Default()
	router.LoadHTMLGlob("html/*.html")
	router.GET("/", HandlerIndex)

	_ = router.Run("127.0.0.1:8080")
}

func HandlerIndex(c *gin.Context) {

	rows, e := connection.Query(`SELECT "Name" FROM "User" WHERE "Login"=$1 AND "Password"=$2`, "admin", "admin")
	if e != nil {
		fmt.Println(e)
		c.HTML(400, "400.html", gin.H{
			"Error": e.Error(),
		})
		return
	}

	var name string

	for rows.Next() {
		e = rows.Scan(&name)
		if e != nil {
			fmt.Println(e)
			c.HTML(400, "400.html", gin.H{
				"Error": e.Error(),
			})
			return
		}
	}

	c.HTML(200, "index.html", gin.H{
		"Name": name,
	})
}
