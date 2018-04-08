package pkg

import (
	"github.com/gin-gonic/gin"
	"time"
	"io/ioutil"
	"encoding/json"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"
	userapi "github.com/openshift/user/api/v1"
)

func GetToken(c *gin.Context)string{
	return c.Request.Header.Get("Authorization")
}

func GetWSToken(c *gin.Context) (ret string){
	ret = "Bearer "+c.Query("access_token")
	return ret
}

func IsWebsocket(c *gin.Context) (bret bool){
	bret = false
	value :=c.Request.Header.Get("Upgrade")
	if value == "websocket"{
		bret = true
	}else {
		bret = false
	}
	return bret
}

func GetUserFromToken(token string) ( string, error) {
	u := &userapi.User{}
	req,err := oapi.GenRequest("GET","/oapi/v1/users/~",token,[]byte{})
	if err != nil{
		return "",err
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	err = json.Unmarshal(result,u)
	if err != nil{
		return "",err
	}
	return u.Name, nil
}

func GetTimeNow()  string{
	//格式化必须是这个时间点，Go语言诞生时间？
	return time.Now().Format("2006-01-02 15:04:05.00")
}


