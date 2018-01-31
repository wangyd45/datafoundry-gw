package buildconfig

import (
	"os"
	//"fmt"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"
)

const (
	BUILD  = "/oapi/v1/buildconfigs"
	BUILDNAME = "/oapi/v1/namespaces/"
	BUILDCONFIG = "/buildconfigs/"
	WATCH = "/oapi/v1/watch/namespaces/"
	WATCHALL = "/oapi/v1/watch/buildconfigs"
	JSON = "application/json"
)

var log lager.Logger

func init(){
	log = lager.NewLogger("oapi_v1_BuildConfig.log")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
}

func CreateBuildConfig(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST",BUILD,token, rBody)
	if err != nil{
		log.Error("CreateBuildConfig error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateBuildConfigInNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST",BUILDNAME +namespace + BUILDCONFIG,token, rBody)
	if err != nil{
		log.Error("CreateBuildConfigInNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func CreateInsInNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST",BUILDNAME +namespace+ BUILDCONFIG + name + "/instantiate",token, rBody)
	if err != nil{
		log.Error("CreateInsInNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func CreateInstInNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST",BUILDNAME +namespace+ BUILDCONFIG + name + "/instantiatebinary",token, rBody)
	if err != nil{
		log.Error("CreateInstInNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func CreateWebInNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST",BUILDNAME +namespace+ BUILDCONFIG + name + "/webhooks",token, rBody)
	if err != nil{
		log.Error("CreateWebInNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func CreateWebInNameSpacePath(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST",BUILDNAME +namespace+ BUILDCONFIG + name + "/webhooks/" + path,token, rBody)
	if err != nil{
		log.Error("CreateWebInNameSpacePath error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func GetBuildConfigFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",BUILDNAME + namespace + BUILDCONFIG + name,token, []byte{})
	if err != nil{
		log.Error("GetBuildConfigFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllBuildConfig(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",BUILD,token, []byte{})
	if err != nil{
		log.Error("GetAllBuildConfig error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllBuildConfigFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",BUILDNAME + namespace + BUILDCONFIG,token, []byte{})
	if err != nil{
		log.Error("GetAllBuildConfigFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchBuildConfigFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",WATCH + namespace + BUILDCONFIG + name,token, []byte{})
	if err != nil{
		log.Error("WatchBuildConfigFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchAllBuildConfig(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",WATCHALL,token, []byte{})
	if err != nil{
		log.Error("WatchAllBuildConfig error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchAllBuildConfigFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",WATCH + namespace + BUILDCONFIG ,token, []byte{})
	if err != nil{
		log.Error("WatchAllBuildConfigFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataBuildConfigFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"PUT",BUILDNAME + namespace + BUILDCONFIG + name,token, rBody)
	if err != nil{
		log.Error("UpdataBuildConfigFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchBuildConfigFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"PATCH",BUILDNAME + namespace + BUILDCONFIG + name,token, rBody)
	if err != nil{
		log.Error("PatchBuildConfigFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteBuildConfigFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"DELETE",BUILDNAME + namespace + BUILDCONFIG + name,token, rBody)
	if err != nil{
		log.Error("DeleteBuildConfigFromNameSpace error ",err)
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
	req,err := oapi.Request(10,"DELETE",BUILDNAME + namespace + BUILDCONFIG,token, rBody)
	if err != nil{
		log.Error("DeleteAllBuildFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
