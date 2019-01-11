package project

import (
	"fmt"
	oapi "github.com/asiainfoldp/datafoundry-gw/apirequest"
	"github.com/asiainfoldp/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"os"
)

var logger lager.Logger

func init() {
	logger = lager.NewLogger("oapi_v1_Project")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

//创建project-OK
func CreateProject(c *gin.Context) {
	//获取前端传递的Token，无需拼接"Bearer XXXXXXXXXX"
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	//获取前端参数
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", "/oapi/v1/projectrequests"+urlParas, token, rBody)
	if err != nil {
		logger.Error("Create A Project Fail", err)
	}
	logger.Info("Create project", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	//返回结果JSON格式
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

//获取project-OK
func getProject(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/oapi/v1/projects/"+name+urlParas, token, nil)
	if err != nil {
		logger.Error("Get A Project Fail", err)
	}
	logger.Info("Get projects/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GorWProject(c *gin.Context) {

	if pkg.IsWebsocket(c) {
		watchAProject(c)
	} else {
		getProject(c)
	}
}

func GorWAllProjects(c *gin.Context) {

	if pkg.IsWebsocket(c) {
		watchAllProjects(c)
	} else {
		getAllProjects(c)
	}
}

//获取project列表-OK
func getAllProjects(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/oapi/v1/projects"+urlParas, token, nil)
	if err != nil {
		logger.Error("Get All Projects Fail", err)
	}
	logger.Info("List projects", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func watchAProject(c *gin.Context) {

	token := pkg.GetWSToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch projects/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/oapi/v1/watch/projects/"+name+urlParas, token, c.Writer, c.Request)
	logger.Info("Watch projects/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})
}

func watchAllProjects(c *gin.Context) {
	token := pkg.GetWSToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch collection projects", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/oapi/v1/watch/projects"+urlParas, token, c.Writer, c.Request)
	logger.Info("Watch collection projects", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})
}

//更新project
func UpdateProject(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", "/oapi/v1/projects/"+name+urlParas, token, rBody)
	if err != nil {
		logger.Error("Update A Project Fail", err)
	}
	logger.Info("Update projects/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchAProject(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", "/oapi/v1/projects/"+name+urlParas, token, rBody)
	if err != nil {
		fmt.Println("Patch A Project :%s Fail", name, err)
	}
	logger.Info("Patch projects/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.JSON(req.StatusCode, gin.H{"application/json": result})
}

//删除project-OK
func DeleteProject(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("DELETE", "/oapi/v1/projects/"+name+urlParas, token, nil)
	if err != nil {
		logger.Error("Delete A Project Fail", err)
	}
	logger.Info("Delete projects/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
