package pkg

import (
	"github.com/gin-gonic/gin"
	"time"
	"io/ioutil"
	"encoding/json"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"
	userapi "github.com/openshift/user/api/v1"
)

var UserMap map[string]string

func init() {
	UserMap = make(map[string]string)
}

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

func SliceToken(token string) string{
	if len(token) > 7{
		return token[7:]
	}
	return ""
}

func GetUserFromToken(token string) ( string, error) {

	if len(UserMap) >100 {
		UserMap = make(map[string]string)
	}

	value,ok := UserMap[token]
	if ok {
		return value,nil
	}
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
	UserMap[token] = u.Name
	return u.Name, nil
}

func GetTimeNow()  string{
	//格式化必须是这个时间点，Go语言诞生时间？
	return time.Now().Format("2006-01-02 15:04:05.00")
}


