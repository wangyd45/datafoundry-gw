package persistentvolume

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
	logger = lager.NewLogger("api_v1_PersistentVolume")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreatePV(c *gin.Context) {
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", "/api/v1/persistentvolumes", token, rBody)
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
	name := c.Param("name")
	req, err := oapi.GenRequest("GET", "/api/v1/persistentvolumes/"+name, token, nil)
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
	req, err := oapi.GenRequest("GET", "/api/v1/persistentvolumes", token, nil)
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
	name := c.Param("name")
	logger.Info("Watch persistentvolume names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/persistentvolumes/"+name, token, c.Writer, c.Request)
	logger.Info("Watch persistentvolume names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllPVs(c *gin.Context) {

	token := pkg.GetWSToken(c)
	logger.Info("Watch collection persistentvolume", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/persistentvolumes", token, c.Writer, c.Request)
	logger.Info("Watch collection persistentvolume", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func UpdatePV(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", "/api/v1/persistentvolumes/"+name, token, rBody)
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
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", "/api/v1/persistentvolumes/"+name, token, rBody)
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
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("DELETE", "/api/v1/persistentvolumes/"+name, token, rBody)
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
	req, err := oapi.GenRequest("DELETE", "/api/v1/persistentvolumes", token, nil)
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
	name := c.Param("name")
	req, err := oapi.GenRequest("GET", "/api/v1/persistentvolumes/"+name+"/status", token, nil)
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
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", "/api/v1/persistentvolumes/"+name+"/status", token, rBody)
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
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", "/api/v1/persistentvolumes/"+name+"/status", token, rBody)
	if err != nil {
		logger.Error("Patch status of a PersistentVolume Fail", err)
	}
	logger.Info("Patch status of persistentvolume names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
