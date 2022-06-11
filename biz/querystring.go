package biz

import (
	"fmt"
	"ginproject/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParam(c *gin.Context) {
	path := c.Param("path")
	param := c.Param("param")

	q := c.QueryMap("ids")
	for idx, v := range q {
		fmt.Println("idx : " + string(idx) + ", v : " + string(v))
	}

	user := structs.User{"tester", "sosa-gu", 29}
	fmt.Println(user)

	c.JSON(http.StatusOK, gin.H{
		"path":  path,
		"param": param,
	})
}
