package project

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"

)

func CreateProject(c *gin.Context){

	token := pkg.GetToken(c)
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"POST","/oapi/v1/projects",token,rBody)
	if err != nil{
		fmt.Println("Create A Project Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.JSON(http.StatusOK, gin.H{"result": result})


}

func GetProject(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.Request(10,"GET","/oapi/v1/projects/"+name,token,nil)
	if err != nil{
		fmt.Println("Get A Project %s Fail",name,err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func GetAllProjects(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/projects",token,nil)
	if err != nil{
		fmt.Println("Get All Project Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.JSON(http.StatusOK, gin.H{"result": result})
}
/*
func WatchAProject(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.Request(10,"GET","/oapi/v1/watch/projects/"+name,token,nil)
	if err != nil{
		fmt.Println("Watch A Project %s Fail",name,err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.JSON(http.StatusOK, gin.H{"result": result})
}

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
func UpdateProject(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"PUT","/oapi/v1/projects"+name,token,rBody)
	if err != nil{
		fmt.Println("Update A Project :%s Fail",name,err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.JSON(http.StatusOK, gin.H{"result": result})
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
func DeleteProject(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.Request(10,"DELETE","/oapi/v1/projects"+name,token,nil)
	if err != nil{
		fmt.Println("Delete A Project :%s Fail",name,err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.JSON(http.StatusOK, gin.H{"result": result})
}

