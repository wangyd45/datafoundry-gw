package resourcequota


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
	SERVICE     = "/api/v1/resourcequotas"
	SERVICENAME = "/api/v1/namespaces"
	WATCH       = "/api/v1/watch/namespaces"
	WATCHALL    = "/api/v1/watch/resourcequota"
	JSON      = "application/json"
)

func init() {
	log = lager.NewLogger("api_v1_ResourceQuota")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateRq(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICE, token, rBody)
	if err != nil {
		log.Error("CreateRq error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateRqInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/resourcequotas", token, rBody)
	if err != nil {
		log.Error("CreateRqInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetRqFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchRqFromNS(c)
	}else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/resourcequotas/"+name, token, []byte{})
		if err != nil {
			log.Error("GetRqFromNS error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllRq(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchAllRq(c)
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

func GetAllRqFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchAllRqFromNS(c)
	}else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/resourcequotas", token, []byte{})
		if err != nil {
			log.Error("GetAllSecretFromNS error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetStuRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/resourcequotas/"+name + "/status", token, []byte{})
	if err != nil {
		log.Error("GetStuRqFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func watchRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCH+ "/" +namespace+"/resourcequotas/" + name,token,c.Writer,c.Request)
}

func watchAllRq(c *gin.Context) {
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCHALL,token,c.Writer,c.Request)
}

func watchAllRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCH+ "/" +namespace+"/resourcequotas", token, c.Writer,c.Request)
}

func UpdataRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" +namespace+"/resourcequotas/"+name, token, rBody)
	if err != nil {
		log.Error("UpdataRqFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataStuRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" + namespace + "/resourcequotas/" + name + "/status", token, rBody)
	if err != nil {
		log.Error("UpdataStuRqFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" +namespace+"/resourcequotas/"+name, token, rBody)
	if err != nil {
		log.Error("PatchRqFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchStuRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME+ "/" + namespace + "/resourcequotas/" + name + "/status", token, rBody)
	if err != nil {
		log.Error("UpdataStuRqFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
func DeleteRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest( "DELETE", SERVICENAME+ "/" +namespace+"/resourcequotas/"+name, token, rBody)
	if err != nil {
		log.Error("DeleteRqFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest( "DELETE", SERVICENAME+ "/" +namespace+"/resourcequotas", token, rBody)
	if err != nil {
		log.Error("DeleteSecretFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
