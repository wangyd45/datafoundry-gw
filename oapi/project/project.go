package project

import (
	"os"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"
	"fmt"
)

var logger lager.Logger

func init() {
	logger = lager.NewLogger("oapi_v1_Project")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

//创建project-OK
func CreateProject(c *gin.Context){
	//获取前端传递的Token，无需拼接"Bearer XXXXXXXXXX"
	token := pkg.GetToken(c)
	//获取前端参数
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req,err := oapi.GenRequest("POST","/oapi/v1/projectrequests",token,rBody)
	if err != nil{
		logger.Error("Create A Project Fail",err)
	}
	logger.Info("Create project",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

//获取project-OK
func getProject(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/oapi/v1/projects/"+name,token,nil)
	if err != nil{
		logger.Error("Get A Project Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}


func GorWProject(c *gin.Context) {

	if pkg.IsWebsocket(c){
		watchAProject(c)
	}else{
		getProject(c)
	}
}

func GorWAllProjects(c *gin.Context) {

	if pkg.IsWebsocket(c){
		watchAllProjects(c)
	}else{
		getAllProjects(c)
	}
}

//获取project列表-OK
func getAllProjects(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET","/oapi/v1/projects",token,nil)
	if err != nil{
		logger.Error("Get All Projects Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func watchAProject(c *gin.Context) {

	token := pkg.GetWSToken(c)
	name := c.Param("name")
	oapi.WSRequest("/oapi/v1/watch/projects/"+name,token,c.Writer,c.Request)
}

func watchAllProjects(c *gin.Context){
	token := pkg.GetWSToken(c)
	oapi.WSRequest("/oapi/v1/watch/projects",token,c.Writer,c.Request)
}


//更新project
func UpdateProject(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/oapi/v1/projects/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A Project Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchAProject(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/oapi/v1/projects/"+name,token,rBody)
	if err != nil{
		fmt.Println("Patch A Project :%s Fail",name,err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.JSON(req.StatusCode, gin.H{"application/json": result})
}

//删除project-OK
func DeleteProject(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("DELETE","/oapi/v1/projects/"+name,token,nil)
	if err != nil{
		logger.Error("Delete A Project Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

