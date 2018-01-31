package imagestreamtag

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
	IMAGE       = "/oapi/v1/imagestreamtags"
	IMAGENAME   = "/oapi/v1/namespaces/"
	IMAGECONFIG = "/imagestreamtags/"
	JSON        = "application/json"
)

var log lager.Logger

func init(){
	log = lager.NewLogger("oapi_v1_ImageStreamTag.log")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
}

func CreateImageStreamTag(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST", IMAGE,token, rBody)
	if err != nil{
		log.Error("CreateImageStreamTag error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateImageTagInNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST", IMAGENAME+namespace +IMAGECONFIG,token, rBody)
	if err != nil{
		log.Error("CreateImageTagInNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func GetImageTagFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET", IMAGENAME+ namespace +IMAGECONFIG+ name,token, []byte{})
	if err != nil{
		log.Error("GetBuildConfigFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllImageTag(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET", IMAGE,token, []byte{})
	if err != nil{
		log.Error("GetAllImageTag error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllImageTagFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET", IMAGENAME+ namespace +IMAGECONFIG,token, []byte{})
	if err != nil{
		log.Error("GetAllImageTagFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataImageTagFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"PUT", IMAGENAME+ namespace +IMAGECONFIG+ name,token, rBody)
	if err != nil{
		log.Error("UpdataImageTagFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchImageTagFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"PATCH", IMAGENAME+ namespace +IMAGECONFIG+ name,token, rBody)
	if err != nil{
		log.Error("PatchImageFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteImageTagFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"DELETE", IMAGENAME+ namespace +IMAGECONFIG+ name,token, rBody)
	if err != nil{
		log.Error("DeleteImagegFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

