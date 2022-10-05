package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", hello)
	// router.GET("/test", test)

	router.GET("/api/hello", hello)
	router.GET("/api/test", test)
	router.GET("/api/users", users)

	router.Run(":8080")
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello from gin framework: ")
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello from gin framework test method: ")
}

func users(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello from gin framework api users method: ")
}
