package role

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
	logger = lager.NewLogger("oapi_v1_Role")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateRole(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("POST", host+"/oapi/v1/roles"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Create A Role Fail", err)
	}
	logger.Info("Create role", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func CreateRoleInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("POST", host+"/oapi/v1/namespaces/"+namespace+"/roles"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Create A Role In A Namespace Fail", err)
	}
	logger.Info("Create role namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)

}

func GetRoleInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", host+"/oapi/v1/namespaces/"+namespace+"/roles/"+name+urlParas, token, nil)

	if err != nil {
		logger.Error("Get A Role In A Namespace Fail", err)
	}
	logger.Info("Get role namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetAllRoles(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", host+"/oapi/v1/roles"+urlParas, token, nil)

	if err != nil {
		logger.Error("Get All Roles Fail", err)
	}
	logger.Info("List roles", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetRolesInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", host+"/oapi/v1/namespaces/"+namespace+"/roles"+urlParas, token, nil)

	if err != nil {
		logger.Error("Get All Roles In A Namespace Fail", err)
	}
	logger.Info("List roles namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdateRoleInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", host+"/oapi/v1/namespaces/"+namespace+"/roles/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Update A Role In A Namespace Fail", err)
	}
	logger.Info("Update role namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchRoleInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", host+"/oapi/v1/namespaces/"+namespace+"/roles/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Patch A Role In A Namespace Fail", err)
	}
	logger.Info("Patch role namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteRoleInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("DELETE", host+"/oapi/v1/namespaces/"+namespace+"/roles/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Delete A Role In A Namespace Fail", err)
	}
	logger.Info("Delete role namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
