package user

import (
	//"os"
	"fmt"
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

//func init(){
//	log = lager.NewLogger("V1_User.log")
//	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
//}

func CreateUser(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST",USER,token, rBody)
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON ,result)
}

func GetUser(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.Request(10,"GET",USER + name,token,[]byte{})
	if err != nil{
		//log.Error("GetUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

func GetAllUser(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",USER,token,[]byte{})
	if err != nil{
		//log.Error("CetAllUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

func WatchUser(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.Request(10,"GET",WATCH + name,token,[]byte{})
	if err != nil{
		//log.Error("WatchAllUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

func WatchAllUser(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",WATCH,token,[]byte{})
	if err != nil{
		//log.Error("WatchAllUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

func UpdataUser(c *gin.Context){
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _:= ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"PUT",USER + name,token,rBody)
	if err != nil{
		//log.Error("DeleteAllUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

func PatchUser(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.Request(10,"PATCH",USER + name,token,[]byte{})
	if err != nil{
		//log.Error("PatchUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

func DeleteUser(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.Request(10,"DELETE",USER + name,token,[]byte{})
	if err != nil{
		//log.Error("DeleteUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

func DeleteAllUser(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"DELETE",USER,token,[]byte{})
	if err != nil{
		//log.Error("DeleteAllUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}
