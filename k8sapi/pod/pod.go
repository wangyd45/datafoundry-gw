package pod


import (
	"os"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	api "github.com/asiainfoLDP/datafoundry-gw/apirequest"
	"strconv"
)

var log lager.Logger

const (
	SERVICE     = "/api/v1/pods"
	SERVICENAME = "/api/v1/namespaces"
	WATCH       = "/api/v1/watch/namespaces"
	WATCHALL    = "/api/v1/watch/pods"
	JSON      = "application/json"
)

func init() {
	log = lager.NewLogger("api_v1_Pod")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreatePod(c *gin.Context){
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICE, token, rBody)
	if err != nil {
		log.Error("CreatePod error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreatePodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/pods", token, rBody)
	if err != nil {
		log.Error("CreatePodInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func AttachPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/pods/"+name+"/attach", token, rBody)
	if err != nil {
		log.Error("AttachPodInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateBindPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/pods/"+name+"/binding",token, rBody)
	if err != nil {
		log.Error("CreateBindPodInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateEvtPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/pods/"+name+"/eviction",token, rBody)
	if err != nil {
		log.Error("CreateEvtPodInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateExecPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/pods/"+name+"/exec",token, rBody)
	if err != nil {
		log.Error("CreateExecPodInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PortPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/pods/"+name+"/portforward",token, rBody)
	if err != nil {
		log.Error("PortPodInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func ProxyPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/pods/"+name+"/proxy",token, rBody)
	if err != nil {
		log.Error("PortPodInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func ProxysPathInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST", SERVICENAME + "/" +namespace+"/pods/"+name+"/proxy/" + path ,token, rBody)
	if err != nil {
		log.Error("ProxysPathInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func HeadPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("HEAD", SERVICENAME + "/" +namespace+"/pods/"+name+"/proxy", token, rBody)
	if err != nil {
		log.Error("HeadPodInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func HeadProxysPathInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("HEAD", SERVICENAME + "/" +namespace+"/pods/"+name+"/proxy/" + path ,token, rBody)
	if err != nil {
		log.Error("HeadProxysPathInNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetPodFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchPodFromNS(c)
	}else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/pods/"+name, token, []byte{})
		if err != nil {
			log.Error("GetPodFromNS error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllPod(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchAllPod(c)
	}else {
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICE, token, []byte{})
		if err != nil {
			log.Error("GetAllPod error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllPodFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c){
		watchAllPodFromNS(c)
	}else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
		req, err := api.GenRequest("GET", SERVICENAME+"/"+namespace+"/pods", token, []byte{})
		if err != nil {
			log.Error("GetAllBuildFromNS error ", err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAtaPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/pods/"+name + "/attach", token, []byte{})
	if err != nil {
		log.Error("GetAtaPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetExecPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/pods/"+name + "/exec", token, []byte{})
	if err != nil {
		log.Error("GetExecPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetLogPodFromNS(c *gin.Context) {
	tailLines:=c.Query("tailLines")
	limitBytes:=c.Query("limitBytes")
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	lenth,e := stringToInt(limitBytes)
	if e != nil{
		lenth = 0
		log.Error("stringToInt error ",e )
	}
	api.WSRequestRL(lenth,SERVICENAME+ "/" +namespace+"/pods/" + name + "/log?follow=true&tailLines="+tailLines+"&limitBytes="+limitBytes,token,c.Writer,c.Request)

}

func GetPortPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/pods/"+name + "/portforward", token, []byte{})
	if err != nil {
		log.Error("GetPortPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetStatusPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/pods/"+name + "/status", token, []byte{})
	if err != nil {
		log.Error("GetPortPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetProxyPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/pods/"+name + "/proxy",token, []byte{})
	if err != nil {
		log.Error("GetProxyPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetProxyPathPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	req, err := api.GenRequest("GET", SERVICENAME + "/" +namespace+"/pods/"+name + "/proxy/" + path,token, []byte{})
	if err != nil {
		log.Error("GetProxyPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func watchPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCH+ "/" +namespace+"/pods/" + name,token,c.Writer,c.Request)
}

func watchAllPod(c *gin.Context) {
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCHALL,token,c.Writer,c.Request)
}

func watchAllPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	api.WSRequest(WATCH+ "/" +namespace+"/pods", token, c.Writer,c.Request)
}

func UpdataPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" +namespace+"/pods/"+name, token, rBody)
	if err != nil {
		log.Error("UpdataPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataStuPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" + namespace + "/pods/" + name + "/status", token, rBody)
	if err != nil {
		log.Error("UpdataStuPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataProxyPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" + namespace + "/pods/" + name + "/proxy", token, rBody)
	if err != nil {
		log.Error("UpdataProxyPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataProPathPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT", SERVICENAME+ "/" + namespace + "/pods/" + name + "/proxy/" + path, token, rBody)
	if err != nil {
		log.Error("UpdataProPathPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" +namespace+"/pods/"+name, token, rBody)
	if err != nil {
		log.Error("PatchPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchStuPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" + namespace + "/pods/" + name + "/status", token, rBody)
	if err != nil {
		log.Error("PatchStuPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchProxyPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" + namespace + "/pods/" + name + "/proxy", token, rBody)
	if err != nil {
		log.Error("PatchProxyPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchProPathPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH", SERVICENAME + "/" + namespace + "/pods/" + name + "/proxy/" + path, token, rBody)
	if err != nil {
		log.Error("PatchProPathPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func OptionsPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("OPTIONS", SERVICENAME + "/" + namespace + "/pods/" + name + "/proxy", token, rBody)
	if err != nil {
		log.Error("PatchPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func OptionsPathPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest("OPTIONS", SERVICENAME + "/" + namespace + "/pods/" + name + "/proxy/" + path, token, rBody)
	if err != nil {
		log.Error("OptionsPathPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeletePodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest( "DELETE", SERVICENAME+ "/" +namespace+"/pods/"+name, token, rBody)
	if err != nil {
		log.Error("DeletePodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteProxyPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest( "DELETE", SERVICENAME+ "/" +namespace+"/pods/"+name+ "/proxy", token, rBody)
	if err != nil {
		log.Error("DeleteProxyPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteProxyPathPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest( "DELETE", SERVICENAME+ "/" +namespace+"/pods/"+name+ "/proxy/" + path, token, rBody)
	if err != nil {
		log.Error("DeleteProxyPathPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	req, err := api.GenRequest( "DELETE", SERVICENAME+ "/" +namespace+"/pods", token, rBody)
	if err != nil {
		log.Error("DeleteAllPodFromNS error ", err)
	}
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func stringToInt(v string) (d int, err error) {
	tmp, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return
	}
	return int(tmp), err
}