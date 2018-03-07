package secret

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
	SERVICE     = "/api/v1/secrets"
	SERVICENAME = "/api/v1/namespaces"
	WATCH       = "/api/v1/watch/namespaces"
	WATCHALL    = "/api/v1/watch/secrets"
	JSON      = "application/json"
)

func init() {
	log = lager.NewLogger("api_v1_Secret")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateSecret(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICE, token, rBody)
	if err != nil {
		log.Error("CreateSecret error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateSecretInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/secrets", token, rBody)
	if err != nil {
		log.Error("CreateSecretInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetSecretFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchSecretFromNS(c)
	}else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/secrets/"+name, token, []byte{})
		if err != nil {
			log.Error("GetSecretFromNS error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllSecret(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchAllSecret(c)
	}else {
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICE, token, []byte{})
		if err != nil {
			log.Error("GetAllSecret error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllSecretFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchAllSecretFromNS(c)
	}else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/secrets", token, []byte{})
		if err != nil {
			log.Error("GetAllSecretFromNS error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func watchSecretFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCH+ "/" +namespace+"/secrets/" + name,token,c.Writer,c.Request)
}

func watchAllSecret(c *gin.Context) {
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCHALL,token,c.Writer,c.Request)
}

func watchAllSecretFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCH+ "/" +namespace+"/secrets", token, c.Writer,c.Request)
}

func UpdataSecretFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" +namespace+"/secrets/"+name, token, rBody)
	if err != nil {
		log.Error("UpdataSecretFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchSecretFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" +namespace+"/secrets/"+name, token, rBody)
	if err != nil {
		log.Error("PatchSecretFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteSecretFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest( "DELETE", SERVICENAME+ "/" +namespace+"/secrets/"+name, token, rBody)
	if err != nil {
		log.Error("DeleteSecretFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllSecretFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest( "DELETE", SERVICENAME+ "/" +namespace+"/secrets", token, rBody)
	if err != nil {
		log.Error("DeleteAllSecretFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}