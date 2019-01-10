package secret

import (
	api "github.com/asiainfoldp/datafoundry-gw/apirequest"
	"github.com/asiainfoldp/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"os"
)

var log lager.Logger

const (
	SERVICE     = "/api/v1/secrets"
	SERVICENAME = "/api/v1/namespaces"
	WATCH       = "/api/v1/watch/namespaces"
	WATCHALL    = "/api/v1/watch/secrets"
	JSON        = "application/json"
)

func init() {
	log = lager.NewLogger("api_v1_Secret")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateSecret(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateSecret Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICE+urlParas, token, rBody)
	if err != nil {
		log.Error("CreateSecret error ", err)
	}
	log.Info("Create Secret", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateSecret Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateSecretInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateSecretInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME+"/"+namespace+"/secrets"+urlParas, token, rBody)
	if err != nil {
		log.Error("CreateSecretInNS error ", err)
	}
	log.Info("Create Secret In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateSecretInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetSecretFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchSecretFromNS(c)
	} else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
		urlParas := pkg.SliceURL(c.Request.URL.String())
		req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/secrets/"+name+urlParas, token, []byte{})
		if err != nil {
			log.Error("GetSecretFromNS error ", err)
		}
		log.Info("Get Secret From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetSecretFromNS Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllSecret(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllSecret(c)
	} else {
		token := pkg.GetToken(c)
		urlParas := pkg.SliceURL(c.Request.URL.String())
		req, err := api.GenRequest("GET", SERVICE+urlParas, token, []byte{})
		if err != nil {
			log.Error("GetAllSecret error ", err)
		}
		log.Info("List Secret ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetAllSecret Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllSecretFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllSecretFromNS(c)
	} else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
		urlParas := pkg.SliceURL(c.Request.URL.String())
		req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/secrets"+urlParas, token, []byte{})
		if err != nil {
			log.Error("GetAllSecretFromNS error ", err)
		}
		log.Info("List Secret From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetAllSecretFromNS Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func watchSecretFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	log.Info("Watch Secret From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	api.WSRequest(WATCH+"/"+namespace+"/secrets/"+name+urlParas, token, c.Writer, c.Request)
	log.Info("Watch Secret From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func watchAllSecret(c *gin.Context) {
	token := pkg.GetWSToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	log.Info("Watch Collection Secret", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	api.WSRequest(WATCHALL+urlParas, token, c.Writer, c.Request)
	log.Info("Watch Collection Secret", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func watchAllSecretFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	log.Info("Watch Collection Secret From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	api.WSRequest(WATCH+"/"+namespace+"/secrets"+urlParas, token, c.Writer, c.Request)
	log.Info("Watch Collection Secret From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func UpdataSecretFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataSecretFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+"/"+namespace+"/secrets/"+name+urlParas, token, rBody)
	if err != nil {
		log.Error("UpdataSecretFromNS error ", err)
	}
	log.Info("Upadata Secret From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataSecretFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchSecretFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchSecretFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME+"/"+namespace+"/secrets/"+name+urlParas, token, rBody)
	if err != nil {
		log.Error("PatchSecretFromNS error ", err)
	}
	log.Info("Patch Secret From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchSecretFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteSecretFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteSecretFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("DELETE", SERVICENAME+"/"+namespace+"/secrets/"+name+urlParas, token, rBody)
	if err != nil {
		log.Error("DeleteSecretFromNS error ", err)
	}
	log.Info("Delete Secret From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteSecretFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllSecretFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteAllSecretFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("DELETE", SERVICENAME+"/"+namespace+"/secrets"+urlParas, token, rBody)
	if err != nil {
		log.Error("DeleteAllSecretFromNS error ", err)
	}
	log.Info("Delete Collection Secret From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteAllSecretFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
