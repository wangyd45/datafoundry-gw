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
	req,err := oapi.Request(10,"POST","/oapi/v1/projectrequests",token,rBody)
	if err != nil{
		logger.Error("Create A Project Fail",err)
	}
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

//获取project-OK
func GetProject(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.Request(10,"GET","/oapi/v1/projects/"+name,token,nil)
	if err != nil{
		logger.Error("Get A Project Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

//获取project列表-OK
func GetAllProjects(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/projects",token,nil)
	if err != nil{
		logger.Error("Get All Projects Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func WatchAProject(c *gin.Context) {

	//go func(){
	token := pkg.GetToken(c)
	name := c.Param("name")
	//for i := 0; i < 16; i++ {
	//	time.Sleep(2 * time.Second)
		req, err := oapi.Request(0, "GET", "/oapi/v1/watch/projects/"+name, token, nil)
		if err != nil {
			logger.Error("Get All Projects Fail", err)
		}


		fmt.Println("----- code ", req.StatusCode)
		fmt.Println("***** status ", req.Status)
		fmt.Println("---**** ", req.Body)

		//aa:
		result,err := ioutil.ReadAll(req.Body)


		//			goto aa
		fmt.Println("------ length body ", len(result))
		defer req.Body.Close()
		c.Data(req.StatusCode, "application/json", result)
	//}

	//}()
}
/*
func WatchAllProjects(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/watch/projects",token,nil)
	if err != nil{
		fmt.Println("Watch All Project Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.JSON(http.StatusOK, gin.H{"result": result})
}
*/

//更新project
func UpdateProject(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"PUT","/oapi/v1/projects/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A Project Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}
/*
func PatchAProject(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"PATCH","/oapi/v1/projects"+name,token,rBody)
	if err != nil{
		fmt.Println("Patch A Project :%s Fail",name,err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.JSON(http.StatusOK, gin.H{"result": result})
}
*/

//删除project-OK
func DeleteProject(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.Request(10,"DELETE","/oapi/v1/projects/"+name,token,nil)
	if err != nil{
		logger.Error("Delete A Project Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

