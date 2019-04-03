package namespace

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
	logger = lager.NewLogger("api_v1_Namespace")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateNamespace(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST",host + "/api/v1/namespaces"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Create A Namespace Fail", err)
	}
	logger.Info("Create namespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GorWNamespace(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchNamespace(c)
	} else {
		getNamespace(c)
	}
}

func GorWAllNamespaces(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllNamespaces(c)
	} else {
		getAllNamespaces(c)
	}
}

func getNamespace(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET",host + "/api/v1/namespaces/"+name+urlParas, token, nil)

	if err != nil {
		logger.Error("Get A Namespace Fail", err)
	}
	logger.Info("Get namespace names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func getAllNamespaces(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET",host + "/api/v1/namespaces"+urlParas, token, nil)

	if err != nil {
		logger.Error("Get All Namespaces Fail", err)
	}
	logger.Info("List namespace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func watchNamespace(c *gin.Context) {

	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	name := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch namespace names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest(host +"/api/v1/watch/namespaces/"+name+urlParas, token, c.Writer, c.Request)

	logger.Info("Watch namespace names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllNamespaces(c *gin.Context) {

	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch collection namespace", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest(host +"/api/v1/watch/namespaces"+urlParas, token, c.Writer, c.Request)

	logger.Info("Watch collection namespace", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func UpdateNamespace(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT",host + "/api/v1/namespaces/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Update A Namespace Fail", err)
	}
	logger.Info("Update namespace names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchNamespace(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH",host + "/api/v1/namespaces/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Patch A Namespace Fail", err)
	}
	logger.Info("Patch namespace names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteNamespace(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("DELETE",host + "/api/v1/namespaces/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Delete A Namespace Fail", err)
	}
	logger.Info("Delete namespace names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdatefinalizeofNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT",host + "/api/v1/namespaces/"+name+"/finalize"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Update finalize of a Namespace Fail", err)
	}
	logger.Info("Update finalize of namespace names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetstatusofNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET",host + "/api/v1/namespaces/"+name+"/status"+urlParas, token, nil)

	if err != nil {
		logger.Error("Get status of a Namespace Fail", err)
	}
	logger.Info("Get status of namespace names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdatestatusofNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT",host + "/api/v1/namespaces/"+name+"/status"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Update status of a Namespace Fail", err)
	}
	logger.Info("Update status of namespace names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchstatusofNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH",host + "/api/v1/namespaces/"+name+"/status"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Patch status of a Namespace Fail", err)
	}
	logger.Info("Patch status of namespace names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
