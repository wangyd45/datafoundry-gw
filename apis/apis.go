package apis

import (
	apirequest "github.com/asiainfoLDP/datafoundry-gw/apirequest"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"os"
)

var logger lager.Logger

func init() {
	logger = lager.NewLogger("apis")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

//apihost=lab.new.dataos.io

func GetHPAns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req, err := apirequest.GenRequest("GET", "/apis/autoscaling/v1/namespaces/"+namespace+"/horizontalpodautoscalers/"+name, token, nil)
	if err != nil {
		logger.Error("Get a HorizontalPodAutoscaler in a namespace Fail", err)
	}
	logger.Info("Get horizontalpodautoscaler namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdateHPAns(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PUT", "/apis/autoscaling/v1/namespaces/"+namespace+"/horizontalpodautoscalers/"+name, token, rBody)
	if err != nil {
		logger.Error("Update HorizontalPodAutoscaler in a namespace error ", err)
	}
	logger.Info("Update horizontalpodautoscaler namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchHPAns(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PATCH", "/apis/autoscaling/v1/namespaces/"+namespace+"/horizontalpodautoscalers/"+name, token, rBody)
	if err != nil {
		logger.Error("Patch HorizontalPodAutoscaler in a namespace error ", err)
	}
	logger.Info("Patch horizontalpodautoscaler namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteHPAns(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("DELETE", "/apis/autoscaling/v1/namespaces/"+namespace+"/horizontalpodautoscalers/"+name, token, rBody)
	if err != nil {
		logger.Error("Delete HorizontalPodAutoscaler in a namespace error ", err)
	}
	logger.Info("Delete horizontalpodautoscaler namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetSFSns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req, err := apirequest.GenRequest("GET", "/apis/apps/v1beta1/namespaces/"+namespace+"/statefulsets/"+name, token, nil)
	if err != nil {
		logger.Error("Get a StatefulSet in a namespace Fail", err)
	}
	logger.Info("Get statefulset namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdateSFSns(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PUT", "/apis/apps/v1beta1/namespaces/"+namespace+"/statefulsets/"+name, token, rBody)
	if err != nil {
		logger.Error("Update StatefulSet in a namespace error ", err)
	}
	logger.Info("Update statefulset namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchSFSns(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PATCH", "/apis/apps/v1beta1/namespaces/"+namespace+"/statefulsets/"+name, token, rBody)
	if err != nil {
		logger.Error("Patch StatefulSet in a namespace error ", err)
	}
	logger.Info("Patch statefulset namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteSFSns(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("DELETE", "/apis/apps/v1beta1/namespaces/"+namespace+"/statefulsets/"+name, token, rBody)
	if err != nil {
		logger.Error("Delete StatefulSet in a namespace error ", err)
	}
	logger.Info("Delete statefulset namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetDeploymentns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req, err := apirequest.GenRequest("GET", "/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name, token, nil)
	if err != nil {
		logger.Error("Get a Deployment in a namespace Fail", err)
	}
	logger.Info("Get extension/deployment namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdateDeploymentns(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PUT", "/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name, token, rBody)
	if err != nil {
		logger.Error("Update Deployment in a namespace error ", err)
	}
	logger.Info("Update extension/deployment namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchDeploymentns(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PATCH", "/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name, token, rBody)
	if err != nil {
		logger.Error("Patch Deployment in a namespace error ", err)
	}
	logger.Info("Patch extension/deployment namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteDeploymentns(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("DELETE", "/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name, token, rBody)
	if err != nil {
		logger.Error("Delete Deployment in a namespace error ", err)
	}
	logger.Info("Delete extension/deployment namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetDeploymentScalens(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req, err := apirequest.GenRequest("GET", "/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name+"/scale", token, nil)
	if err != nil {
		logger.Error("Get scale of a Deployment in a namespace Fail", err)
	}
	logger.Info("Get extension/deployment/scale namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdateDeploymentScalens(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PUT", "/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name+"/scale", token, rBody)
	if err != nil {
		logger.Error("Update scale of a Deployment in a namespace error ", err)
	}
	logger.Info("Update extension/deployment/scale namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchDeploymentScalens(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PATCH", "/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name+"/scale", token, rBody)
	if err != nil {
		logger.Error("Patch scale of a Deployment in a namespace error ", err)
	}
	logger.Info("Patch extension/deployment/scale namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetRSns(c *gin.Context) {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req, err := apirequest.GenRequest("GET", "/apis/extensions/v1beta1/namespaces/"+namespace+"/replicasets/"+name, token, nil)
	if err != nil {
		logger.Error("Get a ReplicaSet in a namespace Fail", err)
	}
	logger.Info("Get replicaset namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdateRSns(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PUT", "/apis/extensions/v1beta1/namespaces/"+namespace+"/replicasets/"+name, token, rBody)
	if err != nil {
		logger.Error("Update ReplicaSet in a namespace error ", err)
	}
	logger.Info("Update replicaset namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchRSns(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PATCH", "/apis/extensions/v1beta1/namespaces/"+namespace+"/replicasets/"+name, token, rBody)
	if err != nil {
		logger.Error("Patch ReplicaSet in a namespace error ", err)
	}
	logger.Info("Patch replicaset namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteRSns(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("DELETE", "/apis/extensions/v1beta1/namespaces/"+namespace+"/replicasets/"+name, token, rBody)
	if err != nil {
		logger.Error("Delete ReplicaSet in a namespace error ", err)
	}
	logger.Info("Delete replicaset namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func WatchAllRSns(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	logger.Info("Watch collection replicaset", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	apirequest.WSRequest("/apis/extensions/v1beta1/watch/namespaces/"+namespace+"/replicasets", token, c.Writer, c.Request)
	logger.Info("Watch collection replicaset", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})
}

func WatchAllDeployns(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	logger.Info("Watch collection extension/deployment namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	apirequest.WSRequest("/apis/extensions/v1beta1/watch/namespaces/"+namespace+"/deployments", token, c.Writer, c.Request)
	logger.Info("Watch collection extension/deployment namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})
}
