package insta

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Photo struct {
	ID    int    `json:id`
	Title string `json:title`
}

var photos []Photo {
	{1, 'Photo1'},
	{2, 'Photo2'}
}

func app() {
	r := gin.Default()
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api/photos", getPhoto)
	r.GET("/api/photos/{id}", getPhoto)
	r.POST("/api/photos", addPhoto)
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
