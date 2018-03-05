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

func CreateBC(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST",BUILD,token, rBody)
	if err != nil{
		log.Error("CreateBC error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateBCInNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST",BUILDNAME +namespace + "/buildconfigs",token, rBody)
	if err != nil{
		log.Error("CreateBCInNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func CreateInsInNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST",BUILDNAME +namespace+ BUILDCONFIG + name + "/instantiate",token, rBody)
	if err != nil{
		log.Error("CreateInsInNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func CreateInstInNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST",BUILDNAME +namespace+ BUILDCONFIG + name + "/instantiatebinary",token, rBody)
	if err != nil{
		log.Error("CreateInstInNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func CreateWebInNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST",BUILDNAME +namespace+ BUILDCONFIG + name + "/webhooks",token, rBody)
	if err != nil{
		log.Error("CreateWebInNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func CreateWebInNSP(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST",BUILDNAME +namespace+ BUILDCONFIG + name + "/webhooks/" + path,token, rBody)
	if err != nil{
		log.Error("CreateWebInNSP error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func GetBCFromNS(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchBCFromNS(c)
	}else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
		req, err := oapi.GenRequest("GET", BUILDNAME+namespace+BUILDCONFIG+name, token, []byte{})
		if err != nil {
			log.Error("GetBCFromNS error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllBC(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchAllBC(c)
	}else {
		token := pkg.GetToken(c)
		req, err := oapi.GenRequest("GET", BUILD, token, []byte{})
		if err != nil {
			log.Error("GetAllBC error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllBCFromNS(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchAllBCFromNS(c)
	}else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
		req, err := oapi.GenRequest("GET", BUILDNAME+namespace+"/buildconfigs", token, []byte{})
		if err != nil {
			log.Error("GetAllBCFromNS error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func watchBCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	oapi.WSRequest(WATCH + namespace + BUILDCONFIG + name, token, c.Writer,c.Request)
}

func watchAllBC(c *gin.Context){
	token := pkg.GetWSToken(c)
	oapi.WSRequest(WATCHALL, token, c.Writer,c.Request)
}

func watchAllBCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	oapi.WSRequest(WATCH + namespace + "/buildconfigs", token, c.Writer,c.Request)
}

func UpdataBCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PUT",BUILDNAME + namespace + BUILDCONFIG + name,token, rBody)
	if err != nil{
		log.Error("UpdataBCFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchBCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PATCH",BUILDNAME + namespace + BUILDCONFIG + name,token, rBody)
	if err != nil{
		log.Error("PatchBCFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteBCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("DELETE",BUILDNAME + namespace + BUILDCONFIG + name,token, rBody)
	if err != nil{
		log.Error("DeleteBCFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllBuildFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("DELETE",BUILDNAME + namespace + BUILDCONFIG,token, rBody)
	if err != nil{
		log.Error("DeleteAllBuildFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
