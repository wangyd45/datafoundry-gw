package user

import (
	"os"
	//"fmt"
	"net/http"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"
)

const (
	USER  = "/oapi/v1/users/"
	WATCH = "/oapi/v1/watch/users/"
	JSON = "application/json"
)

var log lager.Logger

func init(){
	log = lager.NewLogger("oapi_v1_User.log")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
}

//创建用户
func CreateUser(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST",USER,token, rBody)
	if err != nil{
		log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON ,result)
}

//获取用户
func GetUser(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("GET",USER + name,token,[]byte{})
	if err != nil{
		log.Error("GetUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

//获取所有用户
func GetAllUser(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET",USER,token,[]byte{})
	if err != nil{
		log.Error("CetAllUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

func WatchUser(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	oapi.WSRequest(WATCH + name, token, c.Writer,c.Request)
}

//删除所有用户
func DeleteAllUser(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("DELETE",USER,token,[]byte{})
	if err != nil{
		log.Error("DeleteAllUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}
