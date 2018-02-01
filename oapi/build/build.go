package build

import (
	"os"
	//"fmt"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
)

const (
	BUILD     = "/oapi/v1/builds"
	BUILDNAME = "/oapi/v1/namespaces/"
	WATCH     = "/oapi/v1/watch/namespaces/"
	WATCHALL  = "/oapi/v1/watch/builds"
	JSON      = "application/json"
)

var log lager.Logger

func init() {
	log = lager.NewLogger("oapi_v1_Build.log")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
}

func CreateBuild(c *gin.Context) {
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := oapi.Request(10, "POST", BUILD, token, rBody)
	if err != nil {
		log.Error("CreateBuild error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateBuildInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := oapi.Request(10, "POST", BUILDNAME+namespace+"/builds", token, rBody)
	if err != nil {
		log.Error("CreateBuildInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateCloneInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := oapi.Request(10, "POST", BUILDNAME+namespace+"/builds/"+name+"/clone", token, rBody)
	if err != nil {
		log.Error("CreateCloneInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := oapi.Request(10, "GET", BUILDNAME+namespace+"/builds/"+name, token, []byte{})
	if err != nil {
		log.Error("GetBuildFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllBuilds(c *gin.Context) {
	token := pkg.GetToken(c)
	req, err := oapi.Request(10, "GET", BUILD, token, []byte{})
	if err != nil {
		log.Error("GetAllBuilds error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetAllBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	req, err := oapi.Request(10, "GET", BUILDNAME+namespace+"/builds", token, []byte{})
	if err != nil {
		log.Error("GetAllBuildFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetLogBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := oapi.Request(10, "GET", BUILDNAME+namespace+"/builds/"+name+"/log", token, []byte{})
	if err != nil {
		log.Error("GetLogBuildFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := oapi.Request(10, "GET", WATCH+namespace+"/builds"+name, token, []byte{})
	if err != nil {
		log.Error("WatchBuildFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchAllBuilds(c *gin.Context) {
	token := pkg.GetToken(c)
	req, err := oapi.Request(10, "GET", WATCHALL, token, []byte{})
	if err != nil {
		log.Error("WatchAllBuilds error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func WatchAllBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	req, err := oapi.Request(10, "GET", WATCH+namespace+"/builds", token, []byte{})
	if err != nil {
		log.Error("WatchAllBuildFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := oapi.Request(10, "PUT", BUILDNAME+namespace+"/builds"+name, token, rBody)
	if err != nil {
		log.Error("UpdataBuildFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataDetailsInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := oapi.Request(10, "PUT", BUILDNAME+namespace+"/builds/"+name+"/details", token, rBody)
	if err != nil {
		log.Error("UpdataDetailsInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := oapi.Request(10, "PATCH", BUILDNAME+namespace+"/builds"+name, token, rBody)
	if err != nil {
		log.Error("PatchBuildFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := oapi.Request(10, "DELETE", BUILDNAME+namespace+"/builds"+name, token, rBody)
	if err != nil {
		log.Error("DeleteBuildFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllBuildFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := oapi.Request(10, "DELETE", BUILDNAME+namespace+"/builds", token, rBody)
	if err != nil {
		log.Error("DeleteAllBuildFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
