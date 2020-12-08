package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


func Authorize(c *gin.Context) {
	clientID := c.Query("client_id")
	redirectUrl := c.Query("redirect_url")
	session, err := c.Cookie("session")
	if err != nil || session == ""{
		c.Redirect(302, 
			// 重定向到登录界面
			fmt.Sprintf("http://127.0.0.1:9301/oauth/page/login?from=oauth_login&client=%s&redirect_url=%s", clientID, redirectUrl),
		)
	} else {
		c.Redirect(302, 
			// 重定向到确认权限界面
			fmt.Sprintf("http://127.0.0.1:9301/oauth/page/confirm?from=oauth_login&client=%s&redirect_url=%s", clientID, redirectUrl),
		)
	}
}

func Confirm(c *gin.Context) {
	clientID := c.Query("client_id")
	redirectUrl := c.Query("redirect_url")
	session, err := c.Cookie("session")
	if err != nil || session == ""{
		c.Redirect(302, 
			// 重定向到登录界面
			fmt.Sprintf("http://127.0.0.1:9301/oauth/page/login?from=oauth_login&client=%s&redirect_url=%s", clientID, redirectUrl),
		)
	} else {
		code := "123456"
		c.Redirect(302, 
			// 重定向到客户端提供的redirect_url，并带上code
			fmt.Sprintf("%s&code=%s", redirectUrl, code),
		)
	}
}