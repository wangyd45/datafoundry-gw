package service

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
	SERVICE     = "/api/v1/services"
	SERVICENAME = "/api/v1/namespaces"
	WATCH       = "/api/v1/watch/namespaces"
	WATCHALL    = "/api/v1/watch/services"
	JSON      = "application/json"
)

func init() {
	log = lager.NewLogger("api_v1_Services")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateService(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICE, token, rBody)
	if err != nil {
		log.Error("CreateService error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateServiceInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/services", token, rBody)
	if err != nil {
		log.Error("CreateServiceInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateProxysInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/services/"+name+"/proxy", token, rBody)
	if err != nil {
		log.Error("CreateProxysInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateProxysPathInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/services/"+name+"/proxy/" + path,token, rBody)
	if err != nil {
		log.Error("CreateProxysPathInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func HeadProxysInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("HEAD", SERVICENAME + "/" +namespace+"/services/"+name+"/proxy", token, rBody)
	if err != nil {
		log.Error("HeadProxysInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func HeadProxysPathInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("HEAD", SERVICENAME + "/" +namespace+"/services/"+name+"/proxy/" + path ,token, rBody)
	if err != nil {
		log.Error("HeadProxysPathInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetServiceFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchServicesFromNS(c)
	}else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/services/"+name, token, []byte{})
		if err != nil {
			log.Error("GetServiceFromNS error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllServices(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchAllServices(c)
	}else {
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICE, token, []byte{})
		if err != nil {
			log.Error("GetAllServices error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllServicesFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchAllServicesFromNS(c)
	}else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/services", token, []byte{})
		if err != nil {
			log.Error("GetAllServicesFromNS error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetStuServiceFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/services/"+name + "/status", token, []byte{})
	if err != nil {
		log.Error("GetStuServiceFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetProServiceFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/services/"+name + "/proxy", token, []byte{})
	if err != nil {
		log.Error("GetProServiceFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetProPathServiceFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/services/"+name + "/proxy/" + path,token, []byte{})
	if err != nil {
		log.Error("GetProPathServiceFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func watchServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCH+ "/" +namespace+"/services/" + name,token,c.Writer,c.Request)
}

func watchAllServices(c *gin.Context) {
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCHALL,token,c.Writer,c.Request)
}

func watchAllServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCH+ "/" +namespace+"/services", token, c.Writer,c.Request)
}

func UpdataServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" +namespace+"/services/"+name, token, rBody)
	if err != nil {
		log.Error("UpdataServicesFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataStuServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" + namespace + "/services/" + name + "/status", token, rBody)
	if err != nil {
		log.Error("UpdataStuServicesFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataProServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" + namespace + "/services/" + name + "/proxy", token, rBody)
	if err != nil {
		log.Error("UpdataProServicesFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataProPathServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" + namespace + "/services/" + name + "/proxy/" + path, token, rBody)
	if err != nil {
		log.Error("UpdataProPathServicesFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" +namespace+"/services/"+name, token, rBody)
	if err != nil {
		log.Error("PatchServicesFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchStuServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" + namespace + "/services/" + name + "/status", token, rBody)
	if err != nil {
		log.Error("PatchStuServicesFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchProServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" + namespace + "/services/" + name + "/proxy", token, rBody)
	if err != nil {
		log.Error("PatchProServicesFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchProPathServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" + namespace + "/services/" + name + "/proxy/" + path, token, rBody)
	if err != nil {
		log.Error("PatchProPathServicesFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func OptionsServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("OPTIONS", SERVICENAME + "/" + namespace + "/services/" + name + "/proxy", token, rBody)
	if err != nil {
		log.Error("OptionsServicesFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func OptionsPathServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("OPTIONS", SERVICENAME + "/" + namespace + "/services/" + name + "/proxy/" + path, token, rBody)
	if err != nil {
		log.Error("OptionsPathServicesFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteProServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest( "DELETE", SERVICENAME+ "/" +namespace+"/services/"+name+ "/proxy", token, rBody)
	if err != nil {
		log.Error("DeleteProServicesFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteProPathServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest( "DELETE", SERVICENAME+ "/" +namespace+"/services/"+name+ "/proxy/" + path, token, rBody)
	if err != nil {
		log.Error("DeleteProPathServicesFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}