package configmap

import (
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"os"
)

var logger lager.Logger

func init() {
	logger = lager.NewLogger("api_v1_ConfigMap")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateConfigMap(c *gin.Context) {
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", "/api/v1/configmaps", token, rBody)
	if err != nil {
		logger.Error("Create A ConfigMap Fail", err)
	}
	logger.Info("Create configmap", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	//返回结果JSON格式
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func CreateConfigMapNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", "/api/v1/namespaces/"+namespace+"/configmaps", token, rBody)
	if err != nil {
		logger.Error("Create A ConfigMap In A Namespace Fail", err)
	}
	logger.Info("Create configmap namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	//返回结果JSON格式
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GorWConfigMapNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchConfigMapNS(c)
	} else {
		getConfigMapNS(c)
	}
}

func GorWAllConfigMap(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllConfigMap(c)
	} else {
		getAllConfigMap(c)
	}
}

func GorWAllConfigMapNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllConfigMapNS(c)
	} else {
		getAllConfigMapNS(c)
	}
}

func getConfigMapNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req, err := oapi.GenRequest("GET", "/api/v1/namespaces/"+namespace+"/configmaps/"+name, token, nil)
	if err != nil {
		logger.Error("Get A ConfigMap In A Namespace Fail", err)
	}
	logger.Info("Get configmap namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func getAllConfigMap(c *gin.Context) {
	token := pkg.GetToken(c)
	req, err := oapi.GenRequest("GET", "/api/v1/configmaps", token, nil)
	if err != nil {
		logger.Error("Get All ConfigMap Fail", err)
	}
	logger.Info("List configmap", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func getAllConfigMapNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req, err := oapi.GenRequest("GET", "/api/v1/namespaces/"+namespace+"/configmaps", token, nil)
	if err != nil {
		logger.Error("Get All ConfigMap In A Namespace Fail", err)
	}
	logger.Info("List configmap namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func watchConfigMapNS(c *gin.Context) {

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	logger.Info("Watch configmap namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/configmaps/"+name, token, c.Writer, c.Request)
	logger.Info("Watch configmap namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllConfigMap(c *gin.Context) {

	token := pkg.GetWSToken(c)
	logger.Info("Watch collection configmap", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/configmaps", token, c.Writer, c.Request)
	logger.Info("Watch collection configmap", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllConfigMapNS(c *gin.Context) {

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	logger.Info("Watch collection configmap namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/configmaps", token, c.Writer, c.Request)
	logger.Info("Watch collection configmap namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func UpdateConfigMapNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", "/api/v1/namespaces/"+namespace+"/configmaps/"+name, token, rBody)
	if err != nil {
		logger.Error("Update A ConfigMap In A Namespace Fail", err)
	}
	logger.Info("Update configmap namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchConfigMapNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", "/api/v1/namespaces/"+namespace+"/configmaps/"+name, token, rBody)
	if err != nil {
		logger.Error("Patch A ConfigMap In A Namespace Fail", err)
	}
	logger.Info("Patch configmap namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteConfigMapNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("DELETE", "/api/v1/namespaces/"+namespace+"/configmaps/"+name, token, rBody)
	if err != nil {
		logger.Error("Delete A ConfigMap In A Namespace Fail", err)
	}
	logger.Info("Delete configmap namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteAllConfigMapNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req, err := oapi.GenRequest("DELETE", "/api/v1/namespaces/"+namespace+"/configmaps", token, nil)
	if err != nil {
		logger.Error("Delete All ConfigMap In A Namespace Fail", err)
	}
	logger.Info("Delete  collection configmap namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
