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

func CreateIS(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST", IMAGE,token, rBody)
	if err != nil{
		log.Error("CreateIS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateImageInNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST", IMAGENAME+namespace +"/imagestreams",token, rBody)
	if err != nil{
		log.Error("CreateBuildConfigInNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func GetImageFromNS(c *gin.Context){
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

func GetAllImage(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET", IMAGE,token, []byte{})
	if err != nil{
		log.Error("GetAllBuildConfig error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllImageFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET", IMAGENAME+ namespace +"/imagestreams",token, []byte{})
	if err != nil{
		log.Error("GetAllImageFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetSecImageFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET", IMAGENAME+ namespace +IMAGECONFIG+ name + "/secrets",token, []byte{})
	if err != nil{
		log.Error("GetSecImageFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetStaImageFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET", IMAGENAME+ namespace +IMAGECONFIG+ name + "/status",token, []byte{})
	if err != nil{
		log.Error("GetSecImageFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchImageFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	oapi.WSRequest(WATCH + namespace +IMAGECONFIG+ name, token, c.Writer,c.Request)
}

func WatchAllImage(c *gin.Context){
	token := pkg.GetWSToken(c)
	oapi.WSRequest(WATCHALL, token, c.Writer,c.Request)
}

func WatchAllImageFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	oapi.WSRequest(WATCH + namespace +"/imagestreams", token, c.Writer,c.Request)
}

func UpdataImageFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PUT", IMAGENAME+ namespace +IMAGECONFIG+ name,token, rBody)
	if err != nil{
		log.Error("UpdataImageFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataStaImageFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PUT", IMAGENAME+ namespace +IMAGECONFIG+ name + "/status",token, rBody)
	if err != nil{
		log.Error("UpdataImageFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchImageFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PATCH", IMAGENAME+ namespace +IMAGECONFIG+ name,token, rBody)
	if err != nil{
		log.Error("PatchImageFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchStaImageFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PATCH", IMAGENAME+ namespace +IMAGECONFIG+ name + "/status",token, rBody)
	if err != nil{
		log.Error("PatchStaImageFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteImageFromNS(c *gin.Context){
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

func DeleteAllImageFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("DELETE", IMAGENAME+ namespace +"/imagestreams",token, rBody)
	if err != nil{
		log.Error("DeleteAllBuildFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
