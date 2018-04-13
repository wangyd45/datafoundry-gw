package service

import (
	api "github.com/asiainfoLDP/datafoundry-gw/apirequest"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"os"
)

var log lager.Logger

const (
	SERVICE     = "/api/v1/services"
	SERVICENAME = "/api/v1/namespaces"
	WATCH       = "/api/v1/watch/namespaces"
	WATCHALL    = "/api/v1/watch/services"
	JSON        = "application/json"
)

func init() {
	log = lager.NewLogger("api_v1_Services")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateService(c *gin.Context) {
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateService Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICE, token, rBody)
	if err != nil {
		log.Error("CreateService error ", err)
	}
	log.Info("Create Service", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateService Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateServiceInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateServiceInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME+"/"+namespace+"/services", token, rBody)
	if err != nil {
		log.Error("CreateServiceInNS error ", err)
	}
	log.Info("Create Service In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateServiceInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateProxysInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateProxysInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy", token, rBody)
	if err != nil {
		log.Error("CreateProxysInNS error ", err)
	}
	log.Info("Create Proxys In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateProxysInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateProxysPathInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateProxysPathInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy/"+path, token, rBody)
	if err != nil {
		log.Error("CreateProxysPathInNS error ", err)
	}
	log.Info("Create Proxys Path In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateProxysPathInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func HeadProxysInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("HeadProxysInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("HEAD", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy", token, rBody)
	if err != nil {
		log.Error("HeadProxysInNS error ", err)
	}
	log.Info("Head Proxys In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("HeadProxysInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func HeadProxysPathInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("HeadProxysPathInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("HEAD", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy/"+path, token, rBody)
	if err != nil {
		log.Error("HeadProxysPathInNS error ", err)
	}
	log.Info("Head Proxys Path In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("HeadProxysPathInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetServiceFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchServicesFromNS(c)
	} else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/services/"+name, token, []byte{})
		if err != nil {
			log.Error("GetServiceFromNS error ", err)
		}
		log.Info("Get Service From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetServiceFromNS Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllServices(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllServices(c)
	} else {
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICE, token, []byte{})
		if err != nil {
			log.Error("GetAllServices error ", err)
		}
		log.Info("List Services ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetAllServices Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllServicesFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllServicesFromNS(c)
	} else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/services", token, []byte{})
		if err != nil {
			log.Error("GetAllServicesFromNS error ", err)
		}
		log.Info("List Services From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetAllServices Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetStuServiceFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/services/"+name+"/status", token, []byte{})
	if err != nil {
		log.Error("GetStuServiceFromNS error ", err)
	}
	log.Info("Get Status Services From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetStuServiceFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetProServiceFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy", token, []byte{})
	if err != nil {
		log.Error("GetProServiceFromNS error ", err)
	}
	log.Info("Get Pro Services From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetProServiceFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetProPathServiceFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy/"+path, token, []byte{})
	if err != nil {
		log.Error("GetProPathServiceFromNS error ", err)
	}
	log.Info("Get Pro Path Services From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetProPathServiceFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func watchServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	log.Info("Watch Service From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	api.WSRequest(WATCH+"/"+namespace+"/services/"+name, token, c.Writer, c.Request)
	log.Info("Watch Service From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func watchAllServices(c *gin.Context) {
	token := pkg.GetWSToken(c)
	log.Info("Watch Collection Service", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	api.WSRequest(WATCHALL, token, c.Writer, c.Request)
	log.Info("Watch Collection Service", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func watchAllServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	log.Info("Watch Collection Service From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	api.WSRequest(WATCH+"/"+namespace+"/services", token, c.Writer, c.Request)
	log.Info("Watch Collection Service From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func UpdataServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataServicesFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+"/"+namespace+"/services/"+name, token, rBody)
	if err != nil {
		log.Error("UpdataServicesFromNS error ", err)
	}
	log.Info("Upadata Services From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataServicesFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataStuServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataStuServicesFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+"/"+namespace+"/services/"+name+"/status", token, rBody)
	if err != nil {
		log.Error("UpdataStuServicesFromNS error ", err)
	}
	log.Info("Upadata Status Services From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataStuServicesFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataProServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataProServicesFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy", token, rBody)
	if err != nil {
		log.Error("UpdataProServicesFromNS error ", err)
	}
	log.Info("Upadata Pro Services From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataProServicesFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataProPathServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataProPathServicesFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy/"+path, token, rBody)
	if err != nil {
		log.Error("UpdataProPathServicesFromNS error ", err)
	}
	log.Info("Upadata Pro Path Services From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataProPathServicesFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchServicesFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME+"/"+namespace+"/services/"+name, token, rBody)
	if err != nil {
		log.Error("PatchServicesFromNS error ", err)
	}
	log.Info("Patch Services From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchServicesFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchStuServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchStuServicesFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME+"/"+namespace+"/services/"+name+"/status", token, rBody)
	if err != nil {
		log.Error("PatchStuServicesFromNS error ", err)
	}
	log.Info("Patch Status Services From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchStuServicesFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchProServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchProServicesFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy", token, rBody)
	if err != nil {
		log.Error("PatchProServicesFromNS error ", err)
	}
	log.Info("Patch Pro Services From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchProServicesFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchProPathServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchProPathServicesFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy/"+path, token, rBody)
	if err != nil {
		log.Error("PatchProPathServicesFromNS error ", err)
	}
	log.Info("Patch Pro Path Services From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchProPathServicesFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func OptionsServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("OptionsServicesFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("OPTIONS", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy", token, rBody)
	if err != nil {
		log.Error("OptionsServicesFromNS error ", err)
	}
	log.Info("Options Services From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("OptionsServicesFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func OptionsPathServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("OptionsPathServicesFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("OPTIONS", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy/"+path, token, rBody)
	if err != nil {
		log.Error("OptionsPathServicesFromNS error ", err)
	}
	log.Info("Options Path Services From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("OptionsPathServicesFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteProServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteProServicesFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("DELETE", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy", token, rBody)
	if err != nil {
		log.Error("DeleteProServicesFromNS error ", err)
	}
	log.Info("Delete Pro Services From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteProServicesFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteProPathServicesFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteProPathServicesFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("DELETE", SERVICENAME+"/"+namespace+"/services/"+name+"/proxy/"+path, token, rBody)
	if err != nil {
		log.Error("DeleteProPathServicesFromNS error ", err)
	}
	log.Info("Delete Pro Path Service From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteProPathServicesFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
