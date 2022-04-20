package main

import (
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func GetUserHandler(c *gin.Context) {
	// middlewareで認証をして成功すると、ここでユーザーIDを取得できる
	userId := c.MustGet("userId").(string)
	if userId == "" {
		c.JSON(401, gin.H{"error": "Invalid session ID."})
		return
	}

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

	// jsonFromFile, err := ioutil.ReadFile("./admin_gcs_token.json")
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	// log.Printf("secret=%s", string(jsonFromFile))

	client, err := storage.NewClient(c, option.WithCredentialsFile("/Users/aopontan/Desktop/team-i/qin-todo/admin_gcs_token.json"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer client.Close()

	// バケットを操作するため、バケットハンドルを作成する
	bkt := client.Bucket("team-l-qin-todo")
	attrs, err := bkt.Attrs(c)
	if err != nil {
		c.JSON(501, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{ "attrs": attrs})

	// バケットの作成
	// if err := bkt.Create(c, "team-l-qin-todo", nil); err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// }
}

func DeleteUserHandler(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	if userId == "" {
		c.JSON(401, gin.H{"error": "Invalid session ID."})
		return
	}

	err := DeleteUser(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(200)
}
