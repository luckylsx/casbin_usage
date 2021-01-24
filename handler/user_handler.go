package handler

import (
	"casbin_usage/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

func Login(c *gin.Context) {
	username, password := c.PostForm("username"), c.PostForm("password")
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(password)
		log.Fatal(err)
	}
	sessionId := fmt.Sprintf("%s-%s", u.String(), username)
	_ = pkg.GlobalCache.Set(sessionId, []byte(username))
	c.SetCookie("current_subject", sessionId, 30*60, "/resource", "", false, true)
	c.JSON(200, pkg.Response{
		Code:    0,
		Message: username + " login in successfully",
		Data:    nil,
	})
}

func ReadResource(c *gin.Context) {
	c.JSON(200, pkg.Response{
		Code:    0,
		Message: "read resource successfully",
		Data:    "resource",
	})
}

func WriteResource(c *gin.Context) {
	c.JSON(200, pkg.Response{Code: 1, Message: "write resource successfully", Data: "resource"})
}
