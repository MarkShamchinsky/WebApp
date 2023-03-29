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
	router.Static("/assets/", "front/")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", HandlerIndex)
	router.POST("/user/reg", HandlerUserRegistration)
	_ = router.Run("127.0.0.1:8080")
}

func HandlerIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func HandlerUserRegistration(c *gin.Context) {
	var user User
	e := c.BindJSON(&user)
	if e != nil {
		c.JSONP(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
	e = user.Create()
	if e != nil {
		c.JSONP(200, gin.H{
			"Error": "Не удалось зарегестрировать пользователя",
		})
		return
	}
	c.JSON(200, gin.H{
		"Error": nil,
	})
}
