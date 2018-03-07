package replicationcontroller

import (
	"os"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	api "github.com/asiainfoLDP/datafoundry-gw/apirequest"
)

var log lager.Logger

const (
	SERVICE     = "/api/v1/replicationcontrollers"
	SERVICENAME = "/api/v1/namespaces"
	WATCH       = "/api/v1/watch/namespaces"
	WATCHALL    = "/api/v1/watch/replicationcontrollers"
	JSON      = "application/json"
)

func init() {
	log = lager.NewLogger("api_v1_replicationcontroller")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateRc(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICE, token, rBody)
	if err != nil {
		log.Error("CreateRc error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateRcInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/replicationcontrollers", token, rBody)
	if err != nil {
		log.Error("CreateRcInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetRcFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchRcFromNS(c)
	}else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/replicationcontrollers/"+name, token, []byte{})
		if err != nil {
			log.Error("GetRcFromNS error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllRc(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchAllRc(c)
	}else {
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICE, token, []byte{})
		if err != nil {
			log.Error("GetAllRc error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllRcFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchAllRcFromNS(c)
	}else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/replicationcontrollers", token, []byte{})
		if err != nil {
			log.Error("GetAllRcFromNS error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetScaleRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/replicationcontrollers/"+name + "/scale", token, []byte{})
	if err != nil {
		log.Error("GetScaleRcFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetExScaleRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", "/apis/extensions/v1beta1/namespaces/" +namespace+"/replicationcontrollers/"+name + "/scale", token, []byte{})
	if err != nil {
		log.Error("GetExScaleRcFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetStatusRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/replicationcontrollers/"+name + "/status",token, []byte{})
	if err != nil {
		log.Error("GetStatusRcFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func watchRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCH+ "/" +namespace+"/replicationcontrollers/" + name,token,c.Writer,c.Request)
}

func watchAllRc(c *gin.Context) {
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCHALL,token,c.Writer,c.Request)
}

func watchAllRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCH+ "/" +namespace+"/replicationcontrollers", token, c.Writer,c.Request)
}

func UpdataRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" +namespace+"/replicationcontrollers/"+name, token, rBody)
	if err != nil {
		log.Error("UpdataRcFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataScaleRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" + namespace + "/replicationcontrollers/" + name + "/scale", token, rBody)
	if err != nil {
		log.Error("UpdataScaleRcFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataExScaleRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT",  "/apis/extensions/v1beta1/namespaces/" + namespace + "/replicationcontrollers/" + name + "/scale", token, rBody)
	if err != nil {
		log.Error("UpdataExScaleRcFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataStatusRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" + namespace + "/replicationcontrollers/" + name + "/status", token, rBody)
	if err != nil {
		log.Error("UpdataStatusRcFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" +namespace+"/replicationcontrollers/"+name, token, rBody)
	if err != nil {
		log.Error("PatchRcFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchScaleRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" + namespace + "/replicationcontrollers/" + name + "/scale", token, rBody)
	if err != nil {
		log.Error("PatchScaleRcFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchExScaleFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH",  "/apis/extensions/v1beta1/namespaces/" + namespace + "/replicationcontrollers/" + name + "/scale", token, rBody)
	if err != nil {
		log.Error("PatchExScaleFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchStatusRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" + namespace + "/replicationcontrollers/" + name + "/status", token, rBody)
	if err != nil {
		log.Error("PatchStatusRcFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest( "DELETE", SERVICENAME+ "/" +namespace+"/replicationcontrollers/"+name, token, rBody)
	if err != nil {
		log.Error("DeleteBuildFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllRcFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest( "DELETE", SERVICENAME+ "/" +namespace+"/replicationcontrollers", token, rBody)
	if err != nil {
		log.Error("DeleteAllRcFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
