package persistentvolume

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
	logger = lager.NewLogger("api_v1_PersistentVolume")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreatePV(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST",host + "/api/v1/persistentvolumes"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Create A PersistentVolume Fail", err)
	}
	logger.Info("Create persistentvolume", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GorWPV(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchPV(c)
	} else {
		getPV(c)
	}
}

func GorWAllPVs(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllPVs(c)
	} else {
		getAllPVs(c)
	}
}

func getPV(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET",host + "/api/v1/persistentvolumes/"+name+urlParas, token, nil)

	if err != nil {
		logger.Error("Get A PersistentVolume Fail", err)
	}
	logger.Info("Get persistentvolume names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func getAllPVs(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET",host + "/api/v1/persistentvolumes"+urlParas, token, nil)

	if err != nil {
		logger.Error("Get All PersistentVolumes Fail", err)
	}
	logger.Info("List persistentvolume", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func watchPV(c *gin.Context) {

	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch persistentvolume names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest(host +"/api/v1/watch/persistentvolumes/"+name+urlParas, token, c.Writer, c.Request)

	logger.Info("Watch persistentvolume names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllPVs(c *gin.Context) {

	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch collection persistentvolume", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest(host +"/api/v1/watch/persistentvolumes"+urlParas, token, c.Writer, c.Request)

	logger.Info("Watch collection persistentvolume", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func UpdatePV(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT",host + "/api/v1/persistentvolumes/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Update A PersistentVolume Fail", err)
	}
	logger.Info("Update persistentvolume names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchPV(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH",host + "/api/v1/persistentvolumes/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Patch A PersistentVolume Fail", err)
	}
	logger.Info("Patch persistentvolume names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeletePV(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("DELETE",host + "/api/v1/persistentvolumes/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Delete A PersistentVolume Fail", err)
	}
	logger.Info("Delete persistentvolume names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteAllPVs(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("DELETE",host + "/api/v1/persistentvolumes"+urlParas, token, nil)

	if err != nil {
		logger.Error("Delete All PersistentVolumes Fail", err)
	}
	logger.Info("Delete collection persistentvolume", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetstatusofPV(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET",host + "/api/v1/persistentvolumes/"+name+"/status"+urlParas, token, nil)

	if err != nil {
		logger.Error("Get status of a PersistentVolume Fail", err)
	}
	logger.Info("Get status of persistentvolume names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdatestatusofPV(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT",host + "/api/v1/persistentvolumes/"+name+"/status"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Update status of a PersistentVolume Fail", err)
	}
	logger.Info("Update status of persistentvolume names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchstatusofPV(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH",host + "/api/v1/persistentvolumes/"+name+"/status"+urlParas, token, rBody)

	if err != nil {
		logger.Error("Patch status of a PersistentVolume Fail", err)
	}
	logger.Info("Patch status of persistentvolume names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
