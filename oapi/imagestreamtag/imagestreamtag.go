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

func CreateIST(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST", IMAGE,token, rBody)
	if err != nil{
		log.Error("CreateIST error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateImageTagInNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST", IMAGENAME+namespace +IMAGECONFIG,token, rBody)
	if err != nil{
		log.Error("CreateImageTagInNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func GetImageTagFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET", IMAGENAME+ namespace +IMAGECONFIG+ name,token, []byte{})
	if err != nil{
		log.Error("GetBuildConfigFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllImageTag(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET", IMAGE,token, []byte{})
	if err != nil{
		log.Error("GetAllImageTag error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllImageTagFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET", IMAGENAME+ namespace +IMAGECONFIG,token, []byte{})
	if err != nil{
		log.Error("GetAllImageTagFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataImageTagFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PUT", IMAGENAME+ namespace +IMAGECONFIG+ name,token, rBody)
	if err != nil{
		log.Error("UpdataImageTagFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchImageTagFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PATCH", IMAGENAME+ namespace +IMAGECONFIG+ name,token, rBody)
	if err != nil{
		log.Error("PatchImageFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteImageTagFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("DELETE", IMAGENAME+ namespace +IMAGECONFIG+ name,token, rBody)
	if err != nil{
		log.Error("DeleteImagegFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

