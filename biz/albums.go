package biz

import (
	"fmt"
	"ginproject/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAlbums(c *gin.Context) {

	for idx, a := range albums {
		fmt.Println("idx: " + strconv.Itoa(idx) + ", ID :" + a.ID + ", Title : " + a.Title + ", Artist : " + a.Artist + ", Price : " + strconv.Itoa(int(a.Price)))
	}

	c.IndentedJSON(http.StatusOK, albums)
}

var albums = []structs.Album{
	{ID: "1", Title: "Blue Train", Artist: "john Coltrance", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Clifford Brown", Artist: "Cliiford Brown", Price: 29.99},
}
