package imagestream

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
	IMAGE       = "/oapi/v1/imagestreams"
	IMAGENAME   = "/oapi/v1/namespaces/"
	IMAGECONFIG = "/imagestreams/"
	WATCH       = "/oapi/v1/watch/namespaces/"
	WATCHALL    = "/oapi/v1/watch/imagestreams"
	JSON        = "application/json"
)

var log lager.Logger

func init(){
	log = lager.NewLogger("oapi_v1_ImageStream.log")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
}

func CreateImageStream(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST", IMAGE,token, rBody)
	if err != nil{
		log.Error("CreateImageStream error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateImageInNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"POST", IMAGENAME+namespace +IMAGECONFIG,token, rBody)
	if err != nil{
		log.Error("CreateBuildConfigInNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func GetImageFromNameSpace(c *gin.Context){
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

func GetAllImage(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET", IMAGE,token, []byte{})
	if err != nil{
		log.Error("GetAllBuildConfig error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllImageFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET", IMAGENAME+ namespace +IMAGECONFIG,token, []byte{})
	if err != nil{
		log.Error("GetAllImageFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetSecretsImageFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET", IMAGENAME+ namespace +IMAGECONFIG+ name + "/secrets",token, []byte{})
	if err != nil{
		log.Error("GetSecretsImageFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetStatusImageFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET", IMAGENAME+ namespace +IMAGECONFIG+ name + "/status",token, []byte{})
	if err != nil{
		log.Error("GetSecretsImageFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchImageFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",WATCH + namespace +IMAGECONFIG+ name,token, []byte{})
	if err != nil{
		log.Error("WatchImageFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchAllImage(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",WATCHALL,token, []byte{})
	if err != nil{
		log.Error("WatchAllImage error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchAllImageFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET",WATCH + namespace +IMAGECONFIG,token, []byte{})
	if err != nil{
		log.Error("WatchAllImageFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataImageFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"PUT", IMAGENAME+ namespace +IMAGECONFIG+ name,token, rBody)
	if err != nil{
		log.Error("UpdataImageFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataStatusImageFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"PUT", IMAGENAME+ namespace +IMAGECONFIG+ name + "/status",token, rBody)
	if err != nil{
		log.Error("UpdataImageFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchImageFromNameSpace(c *gin.Context){
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

func PatchStatusImageFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"PATCH", IMAGENAME+ namespace +IMAGECONFIG+ name + "/status",token, rBody)
	if err != nil{
		log.Error("PatchStatusImageFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteImageFromNameSpace(c *gin.Context){
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

func DeleteAllImageFromNameSpace(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.Request(10,"DELETE", IMAGENAME+ namespace +IMAGECONFIG,token, rBody)
	if err != nil{
		log.Error("DeleteAllBuildFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
