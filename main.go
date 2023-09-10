package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/json", returnJson)
	router.GET("/json2", returnJson2)

	router.Run(":8081")
}

func returnJson(c *gin.Context) {
	m := map[string]string{"status": "ok"}
	j, _ := json.Marshal(m)
	c.Data(http.StatusOK, "application/json", j)
}

func returnJson2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"狀態": "ok",
	})
}
