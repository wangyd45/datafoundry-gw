package configmap

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
	logger = lager.NewLogger("api_v1_ConfigMap")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateConfigMap(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", host+"/api/v1/configmaps"+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", host+"/api/v1/namespaces/"+namespace+"/configmaps"+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", host+"/api/v1/namespaces/"+namespace+"/configmaps/"+name+urlParas, token, nil)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", host+"/api/v1/configmaps"+urlParas, token, nil)

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
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", host+"/api/v1/namespaces/"+namespace+"/configmaps"+urlParas, token, nil)

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
	host := pkg.GetWsHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch configmap namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest(host+"/api/v1/watch/namespaces/"+namespace+"/configmaps/"+name+urlParas, token, c.Writer, c.Request)

	logger.Info("Watch configmap namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllConfigMap(c *gin.Context) {

	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch collection configmap", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest(host+"/api/v1/watch/configmaps"+urlParas, token, c.Writer, c.Request)

	logger.Info("Watch collection configmap", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllConfigMapNS(c *gin.Context) {

	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch collection configmap namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest(host+"/api/v1/watch/namespaces/"+namespace+"/configmaps"+urlParas, token, c.Writer, c.Request)

	logger.Info("Watch collection configmap namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func UpdateConfigMapNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", host+"/api/v1/namespaces/"+namespace+"/configmaps/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", host+"/api/v1/namespaces/"+namespace+"/configmaps/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("DELETE", host+"/api/v1/namespaces/"+namespace+"/configmaps/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("DELETE", host+"/api/v1/namespaces/"+namespace+"/configmaps"+urlParas, token, nil)

	if err != nil {
		logger.Error("Delete All ConfigMap In A Namespace Fail", err)
	}
	logger.Info("Delete  collection configmap namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
