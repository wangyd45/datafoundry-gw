package deploymentconfig

import (
	"os"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"
)

const (
	DEP       = "/oapi/v1/deploymentconfigs"
	DEPNAME   = "/oapi/v1/namespaces/"
	DEPCONFIG = "/deploymentconfigs/"
	WATCH     = "/oapi/v1/watch/namespaces/"
	WATCHALL  = "/oapi/v1/watch/buildconfigs"
	INS       = "/instantiate"
	JSON      = "application/json"
)

var log lager.Logger

func init(){
	log = lager.NewLogger("oapi_v1_DeploymentConfig.log")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
}

func CreateDC(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST", DEP,token, rBody)
	if err != nil{
		log.Error("CreateDC error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateDCInNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST", DEPNAME+namespace +DEPCONFIG,token, rBody)
	if err != nil{
		log.Error("CreateDCInNS error ",err)
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
	req,err := oapi.GenRequest("POST", DEPNAME+namespace+DEPCONFIG+ name + INS,token, rBody)
	if err != nil{
		log.Error("CreateInsInNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func CreateRollBackInNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST", DEPNAME+namespace+DEPCONFIG+ name + "/rollback",token, rBody)
	if err != nil{
		log.Error("CreateInstInNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func GetDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET", DEPNAME+ namespace +DEPCONFIG+ name,token, []byte{})
	if err != nil{
		log.Error("GetDCFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllDC(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET", DEP,token, []byte{})
	if err != nil{
		log.Error("GetAllDC error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET", DEPNAME+ namespace +DEPCONFIG,token, []byte{})
	if err != nil{
		log.Error("GetAllDCFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetLogDepFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET",DEPNAME + namespace + DEPCONFIG + name +"/log",token, []byte{})
	if err != nil{
		log.Error("GetLogBuildFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetScaleDepFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET",DEPNAME + namespace + DEPCONFIG + name +"/scale",token, []byte{})
	if err != nil{
		log.Error("GetScaleDepFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetStatusDepFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET",DEPNAME + namespace + DEPCONFIG + name +"/status",token, []byte{})
	if err != nil{
		log.Error("GetStatusDepFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	oapi.WSRequest(WATCH + namespace +DEPCONFIG+ name, token, c.Writer,c.Request)
}

func WatchAllDC(c *gin.Context){
	token := pkg.GetToken(c)
	oapi.WSRequest(WATCHALL, token, c.Writer,c.Request)
}

func WatchAllDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	oapi.WSRequest(WATCH + namespace + DEPCONFIG, token, c.Writer,c.Request)
}

func UpdataDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PUT", DEPNAME+ namespace +DEPCONFIG+ name,token, rBody)
	if err != nil{
		log.Error("UpdataDCFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataScaleDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PUT", DEPNAME+ namespace +DEPCONFIG+ name + "/scale",token, rBody)
	if err != nil{
		log.Error("UpdataDCFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataStatusDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PUT", DEPNAME+ namespace +DEPCONFIG+ name + "/status",token, rBody)
	if err != nil{
		log.Error("UpdataDCFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PATCH", DEPNAME+ namespace +DEPCONFIG+ name,token, rBody)
	if err != nil{
		log.Error("PatchDeploymentdConfigFromNameSpace error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchScaleDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PATCH", DEPNAME+ namespace +DEPCONFIG+ name + "/scale",token, rBody)
	if err != nil{
		log.Error("PatchScaleDCFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchStatusDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PATCH", DEPNAME+ namespace +DEPCONFIG+ name + "/status",token, rBody)
	if err != nil{
		log.Error("PatchScaleDCFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("DELETE", DEPNAME+ namespace +DEPCONFIG+ name,token, rBody)
	if err != nil{
		log.Error("DeleteDCFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllDepFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("DELETE", DEPNAME+ namespace +DEPCONFIG,token, rBody)
	if err != nil{
		log.Error("DeleteAllDepFromNS error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

