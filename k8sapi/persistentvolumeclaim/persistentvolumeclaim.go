package persistentvolumeclaim

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
	logger = lager.NewLogger("api_v1_PersistentVolumeClaim")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreatePVC(c *gin.Context) {
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", "/api/v1/persistentvolumeclaims", token, rBody)
	if err != nil {
		logger.Error("Create A PersistentVolumeClaim Fail", err)
	}
	logger.Info("Create persistentvolumeclaim", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func CreatePVCns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", "/api/v1/namespaces/"+namespace+"/persistentvolumeclaims", token, rBody)
	if err != nil {
		logger.Error("Create A PersistentVolumeClaim In A Namespace Fail", err)
	}
	logger.Info("Create persistentvolumeclaim namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GorWPVCns(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchPVCns(c)
	} else {
		getPVCns(c)
	}
}

func GorWAllPVC(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllPVC(c)
	} else {
		getAllPVC(c)
	}
}

func GorWAllPVCns(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllPVCns(c)
	} else {
		getAllPVCns(c)
	}
}

func getPVCns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req, err := oapi.GenRequest("GET", "/api/v1/namespaces/"+namespace+"/persistentvolumeclaims/"+name, token, nil)
	if err != nil {
		logger.Error("Get A PersistentVolumeClaim In A Namespace Fail", err)
	}
	logger.Info("Get persistentvolumeclaim namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func getAllPVC(c *gin.Context) {
	token := pkg.GetToken(c)
	req, err := oapi.GenRequest("GET", "/api/v1/persistentvolumeclaims", token, nil)
	if err != nil {
		logger.Error("Get All PersistentVolumeClaims Fail", err)
	}
	logger.Info("List persistentvolumeclaim", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func getAllPVCns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req, err := oapi.GenRequest("GET", "/api/v1/namespaces/"+namespace+"/persistentvolumeclaims", token, nil)
	if err != nil {
		logger.Error("Get All PersistentVolumeClaim In A Namespace Fail", err)
	}
	logger.Info("List persistentvolumeclaim namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func watchPVCns(c *gin.Context) {

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	logger.Info("Watch persistentvolumeclaim namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/persistentvolumeclaims/"+name, token, c.Writer, c.Request)
	logger.Info("Watch persistentvolumeclaim namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllPVC(c *gin.Context) {

	token := pkg.GetWSToken(c)
	logger.Info("Watch collection persistentvolumeclaim", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/persistentvolumeclaims", token, c.Writer, c.Request)
	logger.Info("Watch collection persistentvolumeclaim", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllPVCns(c *gin.Context) {

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	logger.Info("Watch collection persistentvolumeclaim namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/persistentvolumeclaims", token, c.Writer, c.Request)
	logger.Info("Watch collection persistentvolumeclaim namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func UpdatePVCns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", "/api/v1/namespaces/"+namespace+"/persistentvolumeclaims/"+name, token, rBody)
	if err != nil {
		logger.Error("Update A PersistentVolumeClaim In A Namespace Fail", err)
	}
	logger.Info("Update persistentvolumeclaim namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchPVCns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", "/api/v1/namespaces/"+namespace+"/persistentvolumeclaims/"+name, token, rBody)
	if err != nil {
		logger.Error("Patch A PersistentVolumeClaim In A Namespace Fail", err)
	}
	logger.Info("Patch persistentvolumeclaim namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeletePVCns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("DELETE", "/api/v1/namespaces/"+namespace+"/persistentvolumeclaims/"+name, token, rBody)
	if err != nil {
		logger.Error("Delete A PersistentVolumeClaim In A Namespace Fail", err)
	}
	logger.Info("Delete persistentvolumeclaim namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteAllPVCns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req, err := oapi.GenRequest("DELETE", "/api/v1/namespaces/"+namespace+"/persistentvolumeclaims", token, nil)
	if err != nil {
		logger.Error("Delete All PersistentVolumeClaim In A Namespace Fail", err)
	}
	logger.Info("Delete collection persistentvolumeclaim", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetstatusofPVCns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req, err := oapi.GenRequest("GET", "/api/v1/namespaces/"+namespace+"/persistentvolumeclaims/"+name+"/status", token, nil)
	if err != nil {
		logger.Error("Get status of a PersistentVolumeClaim In A Namespace Fail", err)
	}
	logger.Info("Get status of persistentvolumeclaim namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdatestatusofPVCns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", "/api/v1/namespaces/"+namespace+"/persistentvolumeclaims/"+name+"/status", token, rBody)
	if err != nil {
		logger.Error("Update status of a PersistentVolumeClaim In A Namespace Fail", err)
	}
	logger.Info("Update status of persistentvolumeclaim namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchstatusofPVCns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", "/api/v1/namespaces/"+namespace+"/persistentvolumeclaims/"+name+"/status", token, rBody)
	if err != nil {
		logger.Error("Patch status of a PersistentVolumeClaim In A Namespace Fail", err)
	}
	logger.Info("Patch status of persistentvolumeclaim namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
