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
	log = lager.NewLogger("oapi_v1_DeploymentConfig")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
}

func CreateDC(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("CreateDC Read Request.Body error",err)
	}
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST", DEP,token, rBody)
	if err != nil{
		log.Error("CreateDC error ",err)
	}
	log.Info("Create DC",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("CreateDC Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateDCInNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("CreateDCInNS Read Request.Body error",err)
	}
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST", DEPNAME+namespace +"/deploymentconfigs",token, rBody)
	if err != nil{
		log.Error("CreateDCInNS error ",err)
	}
	log.Info("Create DC In NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("CreateDCInNS Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func CreateInsInNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("CreateInsInNS Read Request.Body error",err)
	}
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST", DEPNAME+namespace+DEPCONFIG+ name + INS,token, rBody)
	if err != nil{
		log.Error("CreateInsInNS error ",err)
	}
	log.Info("Create Ins In NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("CreateInsInNS Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func CreateRollBackInNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("CreateRollBackInNS Read Request.Body error",err)
	}
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("POST", DEPNAME+namespace+DEPCONFIG+ name + "/rollback",token, rBody)
	if err != nil{
		log.Error("CreateInstInNameSpace error ",err)
	}
	log.Info("Create Roll Back In NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("CreateRollBackInNS Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON , result)
}

func GetDCFromNS(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchDCFromNS(c)
	}else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
		req, err := oapi.GenRequest("GET", DEPNAME+namespace+DEPCONFIG+name, token, []byte{})
		if err != nil {
			log.Error("GetDCFromNS error ", err)
		}
		log.Info("Cet DC From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil{
			log.Error("GetDCFromNS Read req.Body error",err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllDC(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchAllDC(c)
	}else {
		token := pkg.GetToken(c)
		req, err := oapi.GenRequest("GET", DEP, token, []byte{})
		if err != nil {
			log.Error("GetAllDC error ", err)
		}
		log.Info("List DC ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil{
			log.Error("GetAllDC Read req.Body error",err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllDCFromNS(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchAllDCFromNS(c)
	}else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
		req, err := oapi.GenRequest("GET", DEPNAME+namespace+"/deploymentconfigs", token, []byte{})
		if err != nil {
			log.Error("GetAllDCFromNS error ", err)
		}
		log.Info("List DC From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil{
			log.Error("GetAllDCFromNS Read req.Body error",err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetLogDepFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET",DEPNAME + namespace + DEPCONFIG + name +"/log",token, []byte{})
	if err != nil{
		log.Error("GetLogDepFromNS error ",err)
	}
	log.Info("Get Log Dep From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("GetLogDepFromNS Read req.Body error",err)
	}
	defer req.Body.Close()
	jstring := "{ \"message\": \""+string(result)+"\"}"
	c.Data(req.StatusCode, JSON, []byte(jstring))
}

func GetScaleDepFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET",DEPNAME + namespace + DEPCONFIG + name +"/scale",token, []byte{})
	if err != nil{
		log.Error("GetScaleDepFromNS error ",err)
	}
	log.Info("Get Scale Dep From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("GetScaleDepFromNS Read req.Body error",err)
	}
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
	log.Info("Get Status Dep From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("GetStatusDepFromNS Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func watchDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	log.Info("Watch A DC From NameSpace",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":"start watch"})
	oapi.WSRequest(WATCH + namespace +DEPCONFIG+ name, token, c.Writer,c.Request)
	log.Info("Watch A DC From NameSpace",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":"end watch"})
}

func watchAllDC(c *gin.Context){
	token := pkg.GetWSToken(c)
	log.Info("Watch collection DC",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":"start watch"})
	oapi.WSRequest(WATCHALL, token, c.Writer,c.Request)
	log.Info("Watch collection DC",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":"end watch"})
}

func watchAllDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	log.Info("Watch collectionA DC From NameSpace",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":"start watch"})
	oapi.WSRequest(WATCH + namespace + "/deploymentconfigs", token, c.Writer,c.Request)
	log.Info("Watch collection DC From NameSpace",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":"end watch"})
}

func UpdataDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("UpdataDCFromNS Read Request.Body error",err)
	}
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PUT", DEPNAME+ namespace +DEPCONFIG+ name,token, rBody)
	if err != nil{
		log.Error("UpdataDCFromNS error ",err)
	}
	log.Info("Upata DC From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("UpdataDCFromNS Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataScaleDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("UpdataScaleDCFromNS Read Request.Body error",err)
	}
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PUT", DEPNAME+ namespace +DEPCONFIG+ name + "/scale",token, rBody)
	if err != nil{
		log.Error("UpdataScaleDCFromNS error ",err)
	}
	log.Info("Upata Scale DC From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("UpdataScaleDCFromNS Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataStatusDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("UpdataStatusDCFromNS Read Request.Body error",err)
	}
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PUT", DEPNAME+ namespace +DEPCONFIG+ name + "/status",token, rBody)
	if err != nil{
		log.Error("UpdataStatusDCFromNS error ",err)
	}
	log.Info("Upata Status DC From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("UpdataStatusDCFromNS Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("PatchDCFromNS Read Request.Body error",err)
	}
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PATCH", DEPNAME+ namespace +DEPCONFIG+ name,token, rBody)
	if err != nil{
		log.Error("PatchDCFromNS error ",err)
	}
	log.Info("Patch DC From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("PatchDCFromNS Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchScaleDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("PatchScaleDCFromNS Read Request.Body error",err)
	}
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PATCH", DEPNAME+ namespace +DEPCONFIG+ name + "/scale",token, rBody)
	if err != nil{
		log.Error("PatchScaleDCFromNS error ",err)
	}
	log.Info("Patch Scale DC From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("PatchScaleDCFromNS Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchStatusDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("PatchStatusDCFromNS Read Request.Body error",err)
	}
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("PATCH", DEPNAME+ namespace +DEPCONFIG+ name + "/status",token, rBody)
	if err != nil{
		log.Error("PatchStatusDCFromNS error ",err)
	}
	log.Info("Patch Status DC From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("PatchStatusDCFromNS Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteDCFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("DeleteDCFromNS Read Request.Body error",err)
	}
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("DELETE", DEPNAME+ namespace +DEPCONFIG+ name,token, rBody)
	if err != nil{
		log.Error("DeleteDCFromNS error ",err)
	}
	log.Info("Delete DC From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("DeleteDCFromNS Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllDepFromNS(c *gin.Context){
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		log.Error("DeleteAllDepFromNS Read Request.Body error",err)
	}
	defer c.Request.Body.Close()
	req,err := oapi.GenRequest("DELETE", DEPNAME+ namespace +"/deploymentconfigs",token, rBody)
	if err != nil{
		log.Error("DeleteAllDepFromNS error ",err)
	}
	log.Info("Delete Collection DC From NameSpace ",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{
		log.Error("DeleteAllDepFromNS Read req.Body error",err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

