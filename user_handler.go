package main

import (
	"io/ioutil"
	"log"

	// "cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	// middlewareで認証をして成功すると、ここでユーザーIDを取得できる
	userId := c.MustGet("userId").(string)

	userInfo, err := GetUser(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, userInfo)
}

func PutUserHandler(c *gin.Context) {
	var reqb PutUserRequestBody
	// userId := c.MustGet("userId").(string)
	err := c.ShouldBindJSON(&reqb)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// client, err := storage.NewClient(c, option)

	jsonFromFile, err := ioutil.ReadFile("./token.json")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Printf("secret=%s", string(jsonFromFile))
}
