package pod

import (
	api "github.com/asiainfoldp/datafoundry-gw/apirequest"
	"github.com/asiainfoldp/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"os"
	"strconv"
)

var log lager.Logger

const (
	SERVICE     = "/api/v1/pods"
	SERVICENAME = "/api/v1/namespaces"
	WATCH       = "/api/v1/watch/namespaces"
	WATCHALL    = "/api/v1/watch/pods"
	JSON        = "application/json"
)

func init() {
	log = lager.NewLogger("api_v1_Pod")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreatePod(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreatePod Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST",host + SERVICE+urlParas, token, rBody)

	if err != nil {
		log.Error("CreatePod error ", err)
	}
	log.Info("Create Pod", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreatePod Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreatePodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreatePodInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST",host + SERVICENAME+"/"+namespace+"/pods"+urlParas, token, rBody)

	if err != nil {
		log.Error("CreatePodInNS error ", err)
	}
	log.Info("Create Pod In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreatePodInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func AttachPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("AttachPodInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/attach"+urlParas, token, rBody)

	if err != nil {
		log.Error("AttachPodInNS error ", err)
	}
	log.Info("Attach Pod In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("AttachPodInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateBindPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateBindPodInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/binding"+urlParas, token, rBody)

	if err != nil {
		log.Error("CreateBindPodInNS error ", err)
	}
	log.Info("Create Bind Pod In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateBindPodInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateEvtPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateEvtPodInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/eviction"+urlParas, token, rBody)

	if err != nil {
		log.Error("CreateEvtPodInNS error ", err)
	}
	log.Info("Create Evt Pod In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateEvtPodInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateExecPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateExecPodInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/exec"+urlParas, token, rBody)

	if err != nil {
		log.Error("CreateExecPodInNS error ", err)
	}
	log.Info("Create Exec Pod In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateExecPodInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PortPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PortPodInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/portforward"+urlParas, token, rBody)

	if err != nil {
		log.Error("PortPodInNS error ", err)
	}
	log.Info("Port Pod In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PortPodInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func ProxyPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("ProxyPodInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy"+urlParas, token, rBody)

	if err != nil {
		log.Error("ProxyPodInNS error ", err)
	}
	log.Info("Proxy Pod In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("ProxyPodInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func ProxysPathInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("ProxysPathInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("POST",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy/"+path+urlParas, token, rBody)

	if err != nil {
		log.Error("ProxysPathInNS error ", err)
	}
	log.Info("Proxys Path In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("ProxysPathInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func HeadPodInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("HeadPodInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("HEAD",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy"+urlParas, token, rBody)

	if err != nil {
		log.Error("HeadPodInNS error ", err)
	}
	log.Info("Head Pod In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("HeadPodInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func HeadProxysPathInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("HeadProxysPathInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("HEAD",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy/"+path+urlParas, token, rBody)

	if err != nil {
		log.Error("HeadProxysPathInNS error ", err)
	}
	log.Info("Head Proxys Path In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("HeadProxysPathInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetPodFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchPodFromNS(c)
	} else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
	host := pkg.GetHost(c)
		urlParas := pkg.SliceURL(c.Request.URL.String())
		req, err := api.GenRequest("GET",host + SERVICENAME+"/"+namespace+"/pods/"+name+urlParas, token, []byte{})

		if err != nil {
			log.Error("GetPodFromNS error ", err)
		}
		log.Info("Get Pod From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("HeadProxysPathInNS Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllPod(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllPod(c)
	} else {
		token := pkg.GetToken(c)
	host := pkg.GetHost(c)
		urlParas := pkg.SliceURL(c.Request.URL.String())
		req, err := api.GenRequest("GET",host + SERVICE+urlParas, token, []byte{})

		if err != nil {
			log.Error("GetAllPod error ", err)
		}
		log.Info("List Pod ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetAllPod Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllPodFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllPodFromNS(c)
	} else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
	host := pkg.GetHost(c)
		urlParas := pkg.SliceURL(c.Request.URL.String())
		req, err := api.GenRequest("GET",host + SERVICENAME+"/"+namespace+"/pods"+urlParas, token, []byte{})

		if err != nil {
			log.Error("GetAllPodFromNS error ", err)
		}
		log.Info("List Pod From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetAllPodFromNS Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAtaPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := api.GenRequest("GET",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/attach"+urlParas, token, []byte{})

	if err != nil {
		log.Error("GetAtaPodFromNS error ", err)
	}
	log.Info("Get Ata Pod From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetAtaPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetExecPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := api.GenRequest("GET",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/exec"+urlParas, token, []byte{})

	if err != nil {
		log.Error("GetExecPodFromNS error ", err)
	}
	log.Info("Get Exec Pod From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetExecPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetLogPodFromNS(c *gin.Context) {
	//tailLines := c.Query("tailLines")
	limitBytes := c.Query("limitBytes")
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	lenth, e := stringToInt(limitBytes)
	if e != nil {
		lenth = 0
		log.Error("GetLogPodFromNS stringToInt error ", e)
	}
	log.Info("Get Pod Log From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	//api.WSRequestRL(lenth, SERVICENAME+"/"+namespace+"/pods/"+name+"/log?follow=true&tailLines="+tailLines+"&limitBytes="+limitBytes, token, c.Writer, c.Request)

	api.WSRequestRL(lenth,host + SERVICENAME+"/"+namespace+"/pods/"+name+"/log"+urlParas, token, c.Writer, c.Request)

	log.Info("Get Pod Log From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})

}

func GetPortPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := api.GenRequest("GET",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/portforward"+urlParas, token, []byte{})

	if err != nil {
		log.Error("GetPortPodFromNS error ", err)
	}
	log.Info("Get Port Pod From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetPortPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetStatusPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := api.GenRequest("GET",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/status"+urlParas, token, []byte{})

	if err != nil {
		log.Error("GetStatusPodFromNS error ", err)
	}
	log.Info("Get Status Pod From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetStatusPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetProxyPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := api.GenRequest("GET",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy"+urlParas, token, []byte{})

	if err != nil {
		log.Error("GetProxyPodFromNS error ", err)
	}
	log.Info("Get Proxy Pod From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetProxyPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetProxyPathPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := api.GenRequest("GET",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy/"+path+urlParas, token, []byte{})

	if err != nil {
		log.Error("GetProxyPathPodFromNS error ", err)
	}
	log.Info("Get Proxy Path Pod From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetProxyPathPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func watchPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	log.Info("Watch Pod From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	api.WSRequest(host +WATCH+"/"+namespace+"/pods/"+name+urlParas, token, c.Writer, c.Request)

	log.Info("Watch Pod From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func watchAllPod(c *gin.Context) {
	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	log.Info("Watch collection ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	api.WSRequest(host +WATCHALL+urlParas, token, c.Writer, c.Request)

	log.Info("Watch collection ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func watchAllPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	host := pkg.GetWsHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	log.Info("Watch collection Pod From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	api.WSRequest(host +WATCH+"/"+namespace+"/pods"+urlParas, token, c.Writer, c.Request)

	log.Info("Watch collection Pod From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func UpdataPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT",host + SERVICENAME+"/"+namespace+"/pods/"+name+urlParas, token, rBody)

	if err != nil {
		log.Error("UpdataPodFromNS error ", err)
	}
	log.Info("Upadata Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataStuPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataStuPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/status"+urlParas, token, rBody)

	if err != nil {
		log.Error("UpdataStuPodFromNS error ", err)
	}
	log.Info("Upadata Status Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataStuPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataProxyPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataProxyPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy"+urlParas, token, rBody)

	if err != nil {
		log.Error("UpdataProxyPodFromNS error ", err)
	}
	log.Info("Upadata Proxy Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataProxyPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataProPathPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataProPathPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PUT",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy/"+path+urlParas, token, rBody)

	if err != nil {
		log.Error("UpdataProPathPodFromNS error ", err)
	}
	log.Info("Upadata Pro Path Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataProPathPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH",host + SERVICENAME+"/"+namespace+"/pods/"+name+urlParas, token, rBody)

	if err != nil {
		log.Error("PatchPodFromNS error ", err)
	}
	log.Info("Patch Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchStuPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchStuPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/status"+urlParas, token, rBody)

	if err != nil {
		log.Error("PatchStuPodFromNS error ", err)
	}
	log.Info("Patch Status Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchStuPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchProxyPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchProxyPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy"+urlParas, token, rBody)

	if err != nil {
		log.Error("PatchProxyPodFromNS error ", err)
	}
	log.Info("Patch Proxy Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchProxyPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchProPathPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchProPathPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("PATCH",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy/"+path+urlParas, token, rBody)

	if err != nil {
		log.Error("PatchProPathPodFromNS error ", err)
	}
	log.Info("Patch Pro Path Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchProPathPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func OptionsPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("OptionsPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("OPTIONS",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy"+urlParas, token, rBody)

	if err != nil {
		log.Error("OptionsPodFromNS error ", err)
	}
	log.Info("Options Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("OptionsPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func OptionsPathPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("OptionsPathPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("OPTIONS",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy/"+path+urlParas, token, rBody)

	if err != nil {
		log.Error("OptionsPathPodFromNS error ", err)
	}
	log.Info("Options Path Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("OptionsPathPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeletePodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeletePodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("DELETE",host + SERVICENAME+"/"+namespace+"/pods/"+name+urlParas, token, rBody)

	if err != nil {
		log.Error("DeletePodFromNS error ", err)
	}
	log.Info("Delete Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeletePodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteProxyPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteProxyPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("DELETE",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy"+urlParas, token, rBody)

	if err != nil {
		log.Error("DeleteProxyPodFromNS error ", err)
	}
	log.Info("Delete Proxy Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteProxyPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteProxyPathPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	path := c.Param("path")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteProxyPathPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("DELETE",host + SERVICENAME+"/"+namespace+"/pods/"+name+"/proxy/"+path+urlParas, token, rBody)

	if err != nil {
		log.Error("DeleteProxyPathPodFromNS error ", err)
	}
	log.Info("Delete Proxy Path Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteProxyPathPodFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllPodFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteAllPodFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := api.GenRequest("DELETE",host + SERVICENAME+"/"+namespace+"/pods"+urlParas, token, rBody)

	if err != nil {
		log.Error("DeleteAllPodFromNS error ", err)
	}
	log.Info("Delete Collection Pod From NameSpace  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteAllPodFromNS Read req.Body error", err)
	}
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
