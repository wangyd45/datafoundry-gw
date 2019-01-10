package endpoints

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
	logger = lager.NewLogger("api_v1_EndPoints")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateEndpoints(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", "/api/v1/endpoints"+urlParas, token, rBody)
	if err != nil {
		logger.Error("Create A Endpoints Fail", err)
	}
	logger.Info("Create endpoints", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func CreateEndpointsNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", "/api/v1/namespaces/"+namespace+"/endpoints"+urlParas, token, rBody)
	if err != nil {
		logger.Error("Create A Endpoints In A Namespace Fail", err)
	}
	logger.Info("Create endpoints namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	//返回结果JSON格式
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GorWEndpointsNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchEndpointsNS(c)
	} else {
		getEndpointsNS(c)
	}
}

func GorWAllEndpoints(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllEndpoints(c)
	} else {
		getAllEndpoints(c)
	}
}

func GorWAllEndpointsNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllEndpointsNS(c)
	} else {
		getAllEndpointsNS(c)
	}
}

func getEndpointsNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/api/v1/namespaces/"+namespace+"/endpoints/"+name+urlParas, token, nil)
	if err != nil {
		logger.Error("Get A Endpoints In A Namespace Fail", err)
	}
	logger.Info("Get endpoints namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func getAllEndpoints(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/api/v1/endpoints"+urlParas, token, nil)
	if err != nil {
		logger.Error("Get All Endpoints Fail", err)
	}
	logger.Info("List endpoints", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func getAllEndpointsNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/api/v1/namespaces/"+namespace+"/endpoints"+urlParas, token, nil)
	if err != nil {
		logger.Error("Get All Endpoints In A Namespace Fail", err)
	}
	logger.Info("List endpoints namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func watchEndpointsNS(c *gin.Context) {

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch endpoints namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/endpoints/"+name+urlParas, token, c.Writer, c.Request)
	logger.Info("Watch endpoints namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllEndpoints(c *gin.Context) {

	token := pkg.GetWSToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch collection endpoints", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/endpoints"+urlParas, token, c.Writer, c.Request)
	logger.Info("Watch collection endpoints", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllEndpointsNS(c *gin.Context) {

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch collection endpoints namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/endpoints"+urlParas, token, c.Writer, c.Request)
	logger.Info("Watch collection endpoints namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func UpdateEndpointsNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", "/api/v1/namespaces/"+namespace+"/endpoints/"+name+urlParas, token, rBody)
	if err != nil {
		logger.Error("Update A Endpoints In A Namespace Fail", err)
	}
	logger.Info("Update endpoints namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchEndpointsNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", "/api/v1/namespaces/"+namespace+"/endpoints/"+name+urlParas, token, rBody)
	if err != nil {
		logger.Error("Patch A Endpoints In A Namespace Fail", err)
	}
	logger.Info("Patch endpoints namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteEndpointsNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("DELETE", "/api/v1/namespaces/"+namespace+"/endpoints/"+name+urlParas, token, rBody)
	if err != nil {
		logger.Error("Delete A Endpoints In A Namespace Fail", err)
	}
	logger.Info("Delete endpoints namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteAllEndpointsNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("DELETE", "/api/v1/namespaces/"+namespace+"/endpoints"+urlParas, token, nil)
	if err != nil {
		logger.Error("Delete All Endpoints In A Namespace Fail", err)
	}
	logger.Info("Delete collection endpoints namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
