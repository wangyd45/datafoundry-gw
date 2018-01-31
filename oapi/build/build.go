package build

import (
	//"os"
	"fmt"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"
)

const (
	BUILD  = "/oapi/v1/builds"
	BUILDNAME = "/oapi/v1/namespaces/"
	WATCH = "/oapi/v1/watch/namespaces/"
	WATCHALL = "/oapi/v1/watch/builds"
	JSON = "application/json"
)

var log lager.Logger

//func init(){
//	log = lager.NewLogger("V1_Build.log")
//	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
//}

func CreateBuild(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST",BUILD,token, rBody)
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateBuildInNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST",BUILDNAME +namespace+ "/builds",token, rBody)
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func CreateCloneInNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST",BUILDNAME +namespace+ "/builds/" + name + "/clone",token, rBody)
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func GetBuildFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",BUILDNAME + namespace + "/builds/" + name,token, []byte{})
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllBuilds(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",BUILD,token, []byte{})
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllBuildFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",BUILDNAME + namespace + "/builds",token, []byte{})
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetLogBuildFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",BUILDNAME + namespace + "/builds/" + name +"/log",token, []byte{})
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchBuildFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",WATCH + namespace + "/builds" + name,token, []byte{})
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchAllBuilds(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",WATCHALL,token, []byte{})
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchAllBuildFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",WATCH + namespace + "/builds" ,token, []byte{})
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataBuildFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"PUT",BUILDNAME + namespace + "/builds" + name,token, rBody)
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataDetailsInNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"PUT",BUILDNAME + namespace + "/builds/" + name + "/details",token, rBody)
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchBuildFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"PATCH",BUILDNAME + namespace + "/builds" + name,token, rBody)
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteBuildFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"DELETE",BUILDNAME + namespace + "/builds" + name,token, rBody)
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllBuildFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"DELETE",BUILDNAME + namespace + "/builds",token, rBody)
	if err != nil{
		fmt.Println("CreateUser error ",err)
		//log.Error("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}














