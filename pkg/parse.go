package pkg

import (
	"github.com/gin-gonic/gin"
)

func Parm(c *gin.Context,key string)string{
	return c.Param(key)
}

func GetToken(c *gin.Context)string{
	return c.Request.Header.Get("Authorization")
}

func GetWSToken(c *gin.Context) (ret string){
	ret = "Bearer "+c.Query("access_token")
	return ret
}

func GetRealToken(c *gin.Context) (ret string){
	ret = c.Request.Header.Get("Authorization")
	if len(ret) > 7{
		ret = ret[7:]
	}
	return ret
}

