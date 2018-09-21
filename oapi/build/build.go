package build

import (
	oapi "github.com/asiainfoldp/datafoundry-gw/apirequest"
	"github.com/asiainfoldp/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"os"
)

const (
	BUILD     = "/oapi/v1/builds"
	BUILDNAME = "/oapi/v1/namespaces"
	WATCH     = "/oapi/v1/watch/namespaces"
	WATCHALL  = "/oapi/v1/watch/builds"
	JSON      = "application/json"
)

var log lager.Logger

func init() {
	log = lager.NewLogger("oapi_v1_Build")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
}

func CreateBuild(c *gin.Context) {
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateBuild Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("POST", BUILD, token, rBody)
	if err != nil {
		log.Error("CreateBuild error ", err)
	}
	log.Info("Create Build", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateBuild Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateBuildInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateBuildInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("POST", BUILDNAME+"/"+namespace+"/builds", token, rBody)
	if err != nil {
		log.Error("CreateBuildInNS error ", err)
	}
	log.Info("Create Build In Namespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateBuildInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateCloneInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateCloneInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("POST", BUILDNAME+"/"+namespace+"/builds/"+name+"/clone", token, rBody)
	if err != nil {
		log.Error("CreateCloneInNS error ", err)
	}
	log.Info("Create Clone In Namespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateCloneInNS Read req.Body error ", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetBuildFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchBuildFromNS(c)
	} else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
		req, err := oapi.GenRequest("GET", BUILDNAME+"/"+namespace+"/builds/"+name, token, []byte{})
		if err != nil {
			log.Error("GetBuildFromNS error ", err)
		}
		log.Info("Cet Build From Namespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetBuildFromNS Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllBuilds(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllBuilds(c)
	} else {
		token := pkg.GetToken(c)
		req, err := oapi.GenRequest("GET", BUILD, token, []byte{})
		if err != nil {
			log.Error("GetAllBuilds error ", err)
		}
		log.Info("List Builds ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetAllBuilds Read req.Body error ", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllBuildFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllBuildFromNS(c)
	} else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
		req, err := oapi.GenRequest("GET", BUILDNAME+"/"+namespace+"/builds", token, []byte{})
		if err != nil {
			log.Error("GetAllBuildFromNS error ", err)
		}
		log.Info("List Builds From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetAllBuildFromNS Read req.Body error ", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetLogBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := oapi.GenRequest("GET", BUILDNAME+"/"+namespace+"/builds/"+name+"/log", token, []byte{})
	if err != nil {
		log.Error("GetLogBuildFromNS error ", err)
	}
	log.Info("Get Build Log From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetLogBuildFromNS Read req.Body error ", err)
	}
	defer req.Body.Close()
	jstring := "{ \"message\": \"" + string(result) + "\"}"
	c.Data(req.StatusCode, JSON, []byte(jstring))
}

func watchBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	log.Info("Watch Build From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	oapi.WSRequest(WATCH+"/"+namespace+"/builds/"+name, token, c.Writer, c.Request)
	log.Info("Watch Build From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func watchAllBuilds(c *gin.Context) {
	token := pkg.GetWSToken(c)
	log.Info("Watch collection Builds", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	oapi.WSRequest(WATCHALL, token, c.Writer, c.Request)
	log.Info("Watch collection Builds", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func watchAllBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	log.Info("Watch collection Builds From Namespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	oapi.WSRequest(WATCH+"/"+namespace+"/builds", token, c.Writer, c.Request)
	log.Info("Watch collection Builds From Namespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func UpdataBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataBuildFromNS read Request.Body error ", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("PUT", BUILDNAME+"/"+namespace+"/builds/"+name, token, rBody)
	if err != nil {
		log.Error("UpdataBuildFromNS error ", err)
	}
	log.Info("Updata A Build From Namespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataBuildFromNS Read req.Body error ", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataDetailsInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataDetailsInNS read Request.Body error ", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("PUT", BUILDNAME+"/"+namespace+"/builds/"+name+"/details", token, rBody)
	if err != nil {
		log.Error("UpdataDetailsInNS request error ", err)
	}
	log.Info("Updata Deatils In Namespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataDetailsInNS Read req.Body error ", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchBuildFromNS read Request.Body error ", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("PATCH", BUILDNAME+"/"+namespace+"/builds/"+name, token, rBody)
	if err != nil {
		log.Error("PatchBuildFromNS error ", err)
	}
	log.Info("Patch Build From Namespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataDetailsInNS Read req.Body error ", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteBuildFromNS read Request.Body error ", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("DELETE", BUILDNAME+"/"+namespace+"/builds/"+name, token, rBody)
	if err != nil {
		log.Error("DeleteBuildFromNS error ", err)
	}
	log.Info("Delete Build From Namespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteBuildFromNS Read req.Body error ", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteAllBuildFromNS read Request.Body error ", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("DELETE", BUILDNAME+"/"+namespace+"/builds", token, rBody)
	if err != nil {
		log.Error("DeleteAllBuildFromNS error ", err)
	}
	log.Info("Delete Collection Build From Namespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteAllBuildFromNS Read req.Body error ", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
