package pkg

import (
	"encoding/json"
	oapi "github.com/asiainfoldp/datafoundry-gw/apirequest"
	"github.com/gin-gonic/gin"
	userapi "github.com/openshift/user/api/v1"
	"io/ioutil"
	"strings"
	"time"
)

type LogData struct {
	Message string `json:"message"`
}

var UserMap map[string]string

type Data struct {
	ApiVersion string      `json:"apiVersion"`
	Kind       string      `json:"kind"`
	GroupNames interface{} `json:"groupNames"`
	Metadata   interface{} `json:"metadata"`
	Spec       interface{} `json:"spec"`
	RoleRef    interface{} `json:"roleRef"`
	Subjects   interface{} `json:"subjects"`
}

type Object struct {
	Objects []Data `json:"objects"`
}

func init() {
	UserMap = make(map[string]string)
}

//http请求获取Token "Beaer ......"
func GetToken(c *gin.Context) string {
	return c.Request.Header.Get("Authorization")
}

func GetHost(c *gin.Context) string {
	return c.Request.Header.Get("Cluster")
}

func GetHawHost(c *gin.Context) string{
	return c.Request.Header.Get("HawkularUrl")
}

//WS请求获取Token "Beaer ......"
func GetWSToken(c *gin.Context) (ret string) {
	ret = "Bearer " + c.Query("access_token")
	return ret
}

func GetWsHost(c *gin.Context) string {
	host := c.Request.URL.Query().Get("Cluster")
	return host
}

//判断请求是否WS请求
func IsWebsocket(c *gin.Context) (bret bool) {
	bret = false
	value := c.Request.Header.Get("Upgrade")
	if value == "websocket" {
		bret = true
	} else {
		bret = false
	}
	return bret
}

//截取Token
func SliceToken(token string) string {
	if len(token) > 7 {
		return token[7:]
	}
	return ""
}

//截取URL
func SliceURL(url string) string {
	sList := strings.Split(url, "?")
	if len(sList) == 2 {
		return "?" + sList[1]
	} else {
		return ""
	}
}

//根据Token获取用户名
func GetUserFromToken(token string) string {

	if len(UserMap) > 100 {
		UserMap = make(map[string]string)
	}

	value, ok := UserMap[token]
	if ok {
		return value
	}
	u := &userapi.User{}
	req, err := oapi.GenRequest("GET", "/oapi/v1/users/~", "Bearer "+token, []byte{})
	if err != nil {
		return ""
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	err = json.Unmarshal(result, u)
	if err != nil {
		return ""
	}
	UserMap[token] = u.Name
	return u.Name
}

//获取实时时间
func GetTimeNow() string {
	//格式化必须是这个时间点，Go语言诞生时间？
	return time.Now().Format("2006-01-02 15:04:05.00")
}

//解析并返回http请求体
func BreakBody(body []byte) (ret map[string][]byte, err error) {
	ret = make(map[string][]byte)
	obj := Object{}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(obj.Objects); i++ {
		body, err := json.Marshal(obj.Objects[i])
		if err != nil {
			return nil, err
		}
		ret[obj.Objects[i].Kind] = body
	}
	return ret, nil
}
