package resourcequota

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
	SERVICE     = "/api/v1/resourcequotas"
	SERVICENAME = "/api/v1/namespaces"
	WATCH       = "/api/v1/watch/namespaces"
	WATCHALL    = "/api/v1/watch/resourcequota"
	JSON        = "application/json"
)

func init() {
	log = lager.NewLogger("api_v1_ResourceQuota")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateRq(c *gin.Context) {
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateRq Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICE, token, rBody)
	if err != nil {
		log.Error("CreateRq error ", err)
	}
	log.Info("Create Rq", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateRq Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateRqInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateRqInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME+"/"+namespace+"/resourcequotas", token, rBody)
	if err != nil {
		log.Error("CreateRqInNS error ", err)
	}
	log.Info("Create Rq In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateRqInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetRqFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchRqFromNS(c)
	} else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/resourcequotas/"+name, token, []byte{})
		if err != nil {
			log.Error("GetRqFromNS error ", err)
		}
		log.Info("Get Rq From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetRqFromNS Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllRq(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllRq(c)
	} else {
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICE, token, []byte{})
		if err != nil {
			log.Error("GetAllRq error ", err)
		}
		log.Info("List Rq ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetAllRq Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllRqFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllRqFromNS(c)
	} else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/resourcequotas", token, []byte{})
		if err != nil {
			log.Error("GetAllRqFromNS error ", err)
		}
		log.Info("List Rq From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetAllRq Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetStuRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/resourcequotas/"+name+"/status", token, []byte{})
	if err != nil {
		log.Error("GetStuRqFromNS error ", err)
	}
	log.Info("Get Status Rq From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetStuRqFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func watchRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	log.Info("Watch Rq From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	api.WSRequest(WATCH+"/"+namespace+"/resourcequotas/"+name, token, c.Writer, c.Request)
	log.Info("Watch Rq From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func watchAllRq(c *gin.Context) {
	token := pkg.GetWSToken(c)
	log.Info("Watch Collection Rq", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	api.WSRequest(WATCHALL, token, c.Writer, c.Request)
	log.Info("Watch Collection Rq", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func watchAllRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	log.Info("Watch Collection Rq From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	api.WSRequest(WATCH+"/"+namespace+"/resourcequotas", token, c.Writer, c.Request)
	log.Info("Watch Collection Rq From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func UpdataRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataRqFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+"/"+namespace+"/resourcequotas/"+name, token, rBody)
	if err != nil {
		log.Error("UpdataRqFromNS error ", err)
	}
	log.Info("Upadata Rq From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataRqFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataStuRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataStuRqFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+"/"+namespace+"/resourcequotas/"+name+"/status", token, rBody)
	if err != nil {
		log.Error("UpdataStuRqFromNS error ", err)
	}
	log.Info("Upadata Status Rq From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataStuRqFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchRqFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME+"/"+namespace+"/resourcequotas/"+name, token, rBody)
	if err != nil {
		log.Error("PatchRqFromNS error ", err)
	}
	log.Info("Patch Rq From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchRqFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchStuRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchStuRqFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME+"/"+namespace+"/resourcequotas/"+name+"/status", token, rBody)
	if err != nil {
		log.Error("PatchStuRqFromNS error ", err)
	}
	log.Info("Patch Status Rq From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchStuRqFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteRqFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("DELETE", SERVICENAME+"/"+namespace+"/resourcequotas/"+name, token, rBody)
	if err != nil {
		log.Error("DeleteRqFromNS error ", err)
	}
	log.Info("Delete Rq From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteRqFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllRqFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteRqFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("DELETE", SERVICENAME+"/"+namespace+"/resourcequotas", token, rBody)
	if err != nil {
		log.Error("DeleteAllRqFromNS error ", err)
	}
	log.Info("Delete Collection Rq From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteRqFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
