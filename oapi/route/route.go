package route

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
	logger = lager.NewLogger("oapi_v1_Route")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateRoute(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("POST",host + "/oapi/v1/routes"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Create A Route Fail", err)
	}
	logger.Info("Create route", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func CreateRouteInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("POST",host + "/oapi/v1/namespaces/"+namespace+"/routes"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Create A Route In A Namespace Fail", err)
	}
	logger.Info("Create route namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GorWRouteInNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchRouteInNS(c)
	} else {
		getRouteInNS(c)
	}
}

func GorWAllRoutes(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllRoutes(c)
	} else {
		getAllRoutes(c)
	}
}

func GorWAllRoutesInNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllRoutesInNS(c)
	} else {
		getAllRoutesInNS(c)
	}
}

func getRouteInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET",host + "/oapi/v1/namespaces/"+namespace+"/routes/"+name+urlParas, token, nil)

	if err != nil {
		logger.Error("Get A Route In A Namespace Fail", err)
	}
	logger.Info("Get route namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func getAllRoutes(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET",host + "/oapi/v1/routes"+urlParas, token, nil)

	if err != nil {
		logger.Error("Get All Routes Fail", err)
	}
	logger.Info("List route", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func getAllRoutesInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET",host + "/oapi/v1/namespaces/"+namespace+"/routes"+urlParas, token, nil)

	if err != nil {
		logger.Error("Get All Routes In A Namespace Fail", err)
	}
	logger.Info("List route namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func watchRouteInNS(c *gin.Context) {

	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch route namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest(host +"/oapi/v1/watch/namespaces/"+namespace+"/routes/"+name+urlParas, token, c.Writer, c.Request)

	logger.Info("Watch route namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllRoutes(c *gin.Context) {

	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch collection route", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest(host +"/oapi/v1/watch/routes"+urlParas, token, c.Writer, c.Request)

	logger.Info("Watch collection route", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllRoutesInNS(c *gin.Context) {

	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch collection route namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest(host +"/oapi/v1/watch/namespaces/"+namespace+"/routes"+urlParas, token, c.Writer, c.Request)

	logger.Info("Watch collection route namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func UpdateRouteInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT",host + "/oapi/v1/namespaces/"+namespace+"/routes/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Update A Route In A Namespace Fail", err)
	}
	logger.Info("Update route", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchRouteInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH",host + "/oapi/v1/namespaces/"+namespace+"/routes/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Patch A Route In A Namespace Fail", err)
	}
	logger.Info("Patch route namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteRouteInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("DELETE",host + "/oapi/v1/namespaces/"+namespace+"/routes/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Delete A Route In A Namespace Fail", err)
	}
	logger.Info("Delete route namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteAllRoutesInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("DELETE",host + "/oapi/v1/namespaces/"+namespace+"/routes"+urlParas, token, nil)

	if err != nil {
		logger.Error("Delete All Routes In A Namespace Fail", err)
	}
	logger.Info("Delete collection route namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetRouteStatusInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET",host + "/oapi/v1/namespaces/"+namespace+"/routes/"+name+"/status"+urlParas, token, nil)

	if err != nil {
		logger.Error("Get Status Of A Route In A Namespace Fail", err)
	}
	logger.Info("Get route status namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdateRouteStatusInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT",host + "/oapi/v1/namespaces/"+namespace+"/routes/"+name+"/status"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Update Status Of A Route In A Namespace Fail", err)
	}
	logger.Info("Update route status namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchRouteStatusInNS(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH",host + "/oapi/v1/namespaces/"+namespace+"/routes/"+name+"/status"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Patch Status Of A Route In A Namespace Fail", err)
	}
	logger.Info("Patch route status namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
