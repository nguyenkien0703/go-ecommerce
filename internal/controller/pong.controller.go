package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)




type PongController struct {

}


func NewPongController() *PongController {
	return &PongController{}
}





func (pc *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "anonymous")
	// c.shouldBindJSON()
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"name":    name,
		"uid":     uid,
		"users":   []string{"user1", "user2"},
	})
}
