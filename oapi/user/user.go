package user

import (
	"os"
	//"fmt"
	oapi "github.com/asiainfoldp/datafoundry-gw/apirequest"
	"github.com/asiainfoldp/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"net/http"
)

const (
	USER  = "/oapi/v1/users"
	WATCH = "/oapi/v1/watch/users/"
	JSON  = "application/json"
)

var log lager.Logger

func init() {
	log = lager.NewLogger("oapi_v1_User.log")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
}

//创建用户
func CreateUser(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("POST", USER+urlParas, token, rBody)
	if err != nil {
		log.Error("CreateUser error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

//获取用户
func GetUser(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", USER+"/"+name+urlParas, token, []byte{})
	if err != nil {
		log.Error("GetUser error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

//获取用户
func GetSelf(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", USER+"/~"+urlParas, token, []byte{})
	if err != nil {
		log.Error("GetSelf error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

//获取所有用户
func GetAllUser(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", USER+urlParas, token, []byte{})
	if err != nil {
		log.Error("CetAllUser error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

func WatchUser(c *gin.Context) {
	token := pkg.GetWSToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	oapi.WSRequest(WATCH+name+urlParas, token, c.Writer, c.Request)
}

func WatchAllUser(c *gin.Context) {
	token := pkg.GetWSToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	oapi.WSRequest(WATCH+urlParas, token, c.Writer, c.Request)
}

//更新用户
func UpdataUser(c *gin.Context) {
	name := c.Param("name")
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("PUT", USER+"/"+name+urlParas, token, rBody)
	if err != nil {
		log.Error("DeleteAllUser error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

func PatchUser(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("PATCH", USER+"/"+name+urlParas, token, []byte{})
	if err != nil {
		log.Error("PatchUser error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

//删除单个用户
func DeleteUser(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("DELETE", USER+"/"+name+urlParas, token, []byte{})
	if err != nil {
		log.Error("DeleteUser error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}

//删除所有用户
func DeleteAllUser(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("DELETE", USER+urlParas, token, []byte{})
	if err != nil {
		log.Error("DeleteAllUser error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(http.StatusOK, JSON, result)
}
