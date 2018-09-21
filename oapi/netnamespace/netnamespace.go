package netnamespace

import (
	oapi "github.com/asiainfoldp/datafoundry-gw/apirequest"
	"github.com/asiainfoldp/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"os"
)

var logger lager.Logger

func init() {
	logger = lager.NewLogger("oapi_v1_NetNamespace")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateNetNamespace(c *gin.Context) {
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("POST", "/oapi/v1/netnamespaces", token, rBody)
	if err != nil {
		logger.Error("Create A NetNamespace Fail", err)
	}
	logger.Info("Create netnamespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GorWNetNamespace(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchNetNamespace(c)
	} else {
		getNetNamespace(c)
	}
}

func GorWAllNetNamespaces(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllNetNamespaces(c)
	} else {
		getAllNetNamespaces(c)
	}
}

func getNetNamespace(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	req, err := oapi.GenRequest("GET", "/oapi/v1/netnamespaces/"+name, token, nil)
	if err != nil {
		logger.Error("Get A NetNamespaces Fail", err)
	}
	logger.Info("Get netnamespaces/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func getAllNetNamespaces(c *gin.Context) {
	token := pkg.GetToken(c)
	req, err := oapi.GenRequest("GET", "/oapi/v1/netnamespaces", token, nil)
	if err != nil {
		logger.Error("Get All NetNamespaces Fail", err)
	}
	logger.Info("List netnamespaces", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func watchNetNamespace(c *gin.Context) {

	token := pkg.GetWSToken(c)
	name := c.Param("name")
	logger.Info("Watch netnamespaces/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/oapi/v1/watch/netnamespaces/"+name, token, c.Writer, c.Request)
	logger.Info("Watch netnamespaces/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllNetNamespaces(c *gin.Context) {

	token := pkg.GetWSToken(c)
	logger.Info("Watch collection netnamespaces", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/oapi/v1/watch/netnamespaces", token, c.Writer, c.Request)
	logger.Info("Watch collection netnamespaces", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func UpdateNetNamespace(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", "/oapi/v1/netnamespaces/"+name, token, rBody)
	if err != nil {
		logger.Error("Update A NetNamespace Fail", err)
	}
	logger.Info("Update netnamespaces/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)

}

func PatchNetNamespace(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", "/oapi/v1/netnamespaces/"+name, token, rBody)
	if err != nil {
		logger.Error("Patch A NetNamespaces Fail", err)
	}
	logger.Info("Patch netnamespaces/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteNetNamespace(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("DELETE", "/oapi/v1/netnamespaces/"+name, token, rBody)
	if err != nil {
		logger.Error("Delete A NetNamespace Fail", err)
	}
	logger.Info("Delete netnamespaces/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteAllNetNamespaces(c *gin.Context) {
	token := pkg.GetToken(c)
	req, err := oapi.GenRequest("DELETE", "/oapi/v1/netnamespaces", token, nil)
	if err != nil {
		logger.Error("Delete All NetNamespaces Fail", err)
	}
	logger.Info("Delete collection netnamespaces", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
