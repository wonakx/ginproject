package biz

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

//ping
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

//pingPost
func PongPost(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(value))

	var data map[string]interface{}
	json.Unmarshal([]byte(value), &data)
	c.JSON(http.StatusOK, gin.H{
		"name": data["name"],
		"age":  data["age"],
	})
}
