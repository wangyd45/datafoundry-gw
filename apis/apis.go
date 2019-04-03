package apis

import (
	apirequest "github.com/asiainfoldp/datafoundry-gw/apirequest"
	"github.com/asiainfoldp/datafoundry-gw/pkg"
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
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := apirequest.GenRequest("GET", host+"/apis/autoscaling/v1/namespaces/"+namespace+"/horizontalpodautoscalers/"+name+urlParas, token, nil)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PUT", host+"/apis/autoscaling/v1/namespaces/"+namespace+"/horizontalpodautoscalers/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PATCH", host+"/apis/autoscaling/v1/namespaces/"+namespace+"/horizontalpodautoscalers/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("DELETE", host+"/apis/autoscaling/v1/namespaces/"+namespace+"/horizontalpodautoscalers/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := apirequest.GenRequest("GET", host+"/apis/apps/v1beta1/namespaces/"+namespace+"/statefulsets/"+name+urlParas, token, nil)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PUT", host+"/apis/apps/v1beta1/namespaces/"+namespace+"/statefulsets/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PATCH", host+"/apis/apps/v1beta1/namespaces/"+namespace+"/statefulsets/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("DELETE", host+"/apis/apps/v1beta1/namespaces/"+namespace+"/statefulsets/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := apirequest.GenRequest("GET", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name+urlParas, token, nil)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PUT", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PATCH", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("DELETE", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := apirequest.GenRequest("GET", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name+"/scale"+urlParas, token, nil)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PUT", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name+"/scale"+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PATCH", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments/"+name+"/scale"+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := apirequest.GenRequest("GET", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/replicasets/"+name+urlParas, token, nil)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PUT", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/replicasets/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("PATCH", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/replicasets/"+name+urlParas, token, rBody)

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
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := apirequest.GenRequest("DELETE", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/replicasets/"+name+urlParas, token, rBody)

	if err != nil {
		logger.Error("Delete ReplicaSet in a namespace error ", err)
	}
	logger.Info("Delete replicaset namespaces/"+namespace+"/names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GorWAllRSns(c *gin.Context) {

	if pkg.IsWebsocket(c) {
		watchAllRSns(c)
	} else {
		getAllRSns(c)
	}
}

func GorWAllDeployns(c *gin.Context) {

	if pkg.IsWebsocket(c) {
		watchAllDeployns(c)
	} else {
		getAllDeployns(c)
	}
}

func getAllRSns(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := apirequest.GenRequest("GET", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/replicasets"+urlParas, token, nil)

	if err != nil {
		logger.Error("Get all ReplicaSet in a namespace Fail", err)
	}
	logger.Info("Get all replicaset namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func watchAllRSns(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch collection replicaset", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	apirequest.WSRequest(host+"/apis/extensions/v1beta1/watch/namespaces/"+namespace+"/replicasets"+urlParas, token, c.Writer, c.Request)

	logger.Info("Watch collection replicaset", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})
}

func getAllDeployns(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := apirequest.GenRequest("GET", host+"/apis/extensions/v1beta1/namespaces/"+namespace+"/deployments"+urlParas, token, nil)

	if err != nil {
		logger.Error("Get all Deployment in a namespace Fail", err)
	}
	logger.Info("Get extension/deployments namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func watchAllDeployns(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch collection extension/deployment namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	apirequest.WSRequest(host+"/apis/extensions/v1beta1/watch/namespaces/"+namespace+"/deployments"+urlParas, token, c.Writer, c.Request)

	logger.Info("Watch collection extension/deployment namespaces/"+namespace, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})
}
