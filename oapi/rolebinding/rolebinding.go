package rolebinding

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
	logger = lager.NewLogger("oapi_v1_RoleBinding")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateRoleBinding(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("POST", "/oapi/v1/rolebindings"+urlParas, token, rBody)
	if err != nil {
		logger.Error("Create A RoleBinding Fail", err)
	}
	logger.Info("Create rolebinding", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func CreateRoleBindingInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("POST", "/oapi/v1/namespaces/"+namespace+"/rolebindings"+urlParas, token, rBody)
	if err != nil {
		logger.Error("Create A RoleBinding In A Namespace Fail", err)
	}
	logger.Info("Create rolebinding namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetRoleBindingInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/oapi/v1/namespaces/"+namespace+"/rolebindings/"+name+urlParas, token, nil)
	if err != nil {
		logger.Error("Get A RoleBinding In A Namespace Fail", err)
	}
	logger.Info("Get rolebinding namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetAllRoleBindings(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/oapi/v1/rolebindings"+urlParas, token, nil)
	if err != nil {
		logger.Error("Get All RoleBindings Fail", err)
	}
	logger.Info("List rolebindings", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetRoleBindingsInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/oapi/v1/namespaces/"+namespace+"/rolebindings"+urlParas, token, nil)
	if err != nil {
		logger.Error("Get All RoleBindings In A Namespace Fail", err)
	}
	logger.Info("List rolebindings namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdateRoleBindingInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", "/oapi/v1/namespaces/"+namespace+"/rolebindings/"+name+urlParas, token, rBody)
	if err != nil {
		logger.Error("Update A RoleBinding In A Namespace Fail", err)
	}
	logger.Info("Update rolebinding namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchRoleBindingInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", "/oapi/v1/namespaces/"+namespace+"/rolebindings/"+name+urlParas, token, rBody)
	if err != nil {
		logger.Error("Patch A RoleBinding In A Namespace Fail", err)
	}
	logger.Info("Patch rolebinding namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteRoleBindingInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("DELETE", "/oapi/v1/namespaces/"+namespace+"/rolebindings/"+name+urlParas, token, rBody)
	if err != nil {
		logger.Error("Delete A RoleBinding In A Namespace Fail", err)
	}
	logger.Info("Delete rolebinding namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
