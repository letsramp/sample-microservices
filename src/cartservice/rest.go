package main

import (
	pb "cartservice/genproto"
	"fmt"

	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {
	user_id := c.Param("user_id")
	cart := GetCart(user_id)
	if cart != nil {
		c.JSON(200, cart)
	} else {
		c.JSON(404, gin.H{"error": fmt.Sprintf("cart not found for user_id [%s]", user_id)})
	}
}

func post(c *gin.Context) {
	user_id := c.Param("user_id")
	if len(user_id) < 1 {
		c.JSON(403, gin.H{"error": fmt.Sprintf("invalid user id [%s]", user_id)})
		return
	}
	var item pb.CartItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(501, gin.H{"error": "failed updating cart"})
		return
	}
	AddItem(user_id, item.ProductId, item.Quantity)
	c.JSON(200, gin.H{"success": "200 OK"})
}

func delete(c *gin.Context) {
	user_id := c.Param("user_id")
	if len(user_id) < 1 {
		c.JSON(403, gin.H{"error": fmt.Sprintf("invalid user id [%s]", user_id)})
		return
	}
	EmtyCart(user_id)
	c.JSON(200, gin.H{"success": "200 OK"})
}

func runRest(port string) {
	go func() {
		log.Infof("rest server started on port %s", port)
		router := gin.Default()
		router.GET("/cart/:user_id", get)
		router.POST("/cart/:user_id", post)
		router.DELETE("/cart/:user_id", delete)
		router.Run(fmt.Sprintf("0.0.0.0:%s", port))
	}()
}
