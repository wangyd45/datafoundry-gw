package node

import (
	"encoding/json"
	oapi "github.com/asiainfoldp/datafoundry-gw/apirequest"
	"github.com/asiainfoldp/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"os"
)

var logger lager.Logger

func init() {
	logger = lager.NewLogger("api_v1_Node")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateNode(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req, err := oapi.GenRequest("POST", "/api/v1/nodes"+urlParas, token, rBody)
	if err != nil {
		logger.Error("Create A Node Fail", err)
	}
	logger.Info("Create node", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GorWNode(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchNode(c)
	} else {
		getNode(c)
	}
}

func GorWAllNodes(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllNodes(c)
	} else {
		getAllNodes(c)
	}
}

func GetAllNodesLabels(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/api/v1/nodes"+urlParas, token, nil)
	if err != nil {
		logger.Error("Get All Nodes Labels Fail", err)
	}
	logger.Info("List node Labels ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	var nodeInfo nodes
	nodesMap := make(map[string]map[string]string)
	err = json.Unmarshal(result, &nodeInfo)
	if err != nil {
		logger.Error("GetAllNodesLabels json unmarshal error ", err)
	}
	for _, v := range nodeInfo.Items {
		nodesMap[v.Name] = v.Labels
	}
	result, err = json.Marshal(nodesMap)
	if err != nil {
		logger.Error("GetAllNodesLabels json Marshal error ", err)
	}
	c.Data(req.StatusCode, "", result)
}

func getNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/api/v1/nodes/"+name+urlParas, token, nil)
	if err != nil {
		logger.Error("Get A Node Fail", err)
	}
	logger.Info("Get node", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func getAllNodes(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/api/v1/nodes"+urlParas, token, nil)
	if err != nil {
		logger.Error("Get All Nodes Fail", err)
	}
	logger.Info("List node", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func watchNode(c *gin.Context) {

	token := pkg.GetWSToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/nodes/"+name+urlParas, token, c.Writer, c.Request)
	logger.Info("Watch node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func watchAllNodes(c *gin.Context) {

	token := pkg.GetWSToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	logger.Info("Watch collection node", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "begin"})
	oapi.WSRequest("/api/v1/watch/nodes"+urlParas, token, c.Writer, c.Request)
	logger.Info("Watch collection node", map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(), "result": "end"})

}

func UpdateNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", "/api/v1/nodes/"+name+urlParas, token, rBody)
	if err != nil {
		logger.Error("Update A Node Fail", err)
	}
	logger.Info("Update node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", "/api/v1/nodes/"+name+urlParas, token, rBody)
	if err != nil {
		logger.Error("Patch A Node Fail", err)
	}
	logger.Info("Patch node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("DELETE", "/api/v1/nodes/"+name+urlParas, token, rBody)
	if err != nil {
		logger.Error("Delete A Node Fail", err)
	}
	logger.Info("Delete node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func DeleteAllNodes(c *gin.Context) {
	token := pkg.GetToken(c)
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("DELETE", "/api/v1/nodes"+urlParas, token, nil)
	if err != nil {
		logger.Error("Delete All Nodes Fail", err)
	}
	logger.Info("Delete collection node", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func GetStatusOfNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/api/v1/nodes/"+name+"/status"+urlParas, token, nil)
	if err != nil {
		logger.Error("Get Status Of A Node Fail", err)
	}
	logger.Info("Get status of node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func UpdateStatusOfNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PUT", "/api/v1/nodes/"+name+"/status"+urlParas, token, rBody)
	if err != nil {
		logger.Error("Update Status Of A Node Fail", err)
	}
	logger.Info("Update status of node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func PatchStatusOfNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)
	req, err := oapi.GenRequest("PATCH", "/api/v1/nodes/"+name+"/status"+urlParas, token, rBody)
	if err != nil {
		logger.Error("Patch Status Of A Node Fail", err)
	}
	logger.Info("Patch status of node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyOpnReqToNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("OPTIONS", "/api/v1/nodes/"+name+"/proxy"+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy OPTIONS Request To A Node Fail", err)
	}
	logger.Info("Proxy OPTIONS node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyPostReqToNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("POST", "/api/v1/nodes/"+name+"/proxy"+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy Post Request To A Node Fail", err)
	}
	logger.Info("Proxy Post node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyHeadReqToNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("HEAD", "/api/v1/nodes/"+name+"/proxy"+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy Head Request To A Node Fail", err)
	}
	logger.Info("Proxy Head node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyGetReqToNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/api/v1/nodes/"+name+"/proxy"+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy Get Request To A Node Fail", err)
	}
	logger.Info("Proxy Get node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyPutReqToNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("PUT", "/api/v1/nodes/"+name+"/proxy"+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy Put Request To A Node Fail", err)
	}
	logger.Info("Proxy Put node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyPatchReqToNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("PATCH", "/api/v1/nodes/"+name+"/proxy"+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy Patch Request To A Node Fail", err)
	}
	logger.Info("Proxy Patch node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyDelReqToNode(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("DELETE", "/api/v1/nodes/"+name+"/proxy"+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy Delete Request To A Node Fail", err)
	}
	logger.Info("Proxy Delete node names/"+name, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyOpnReqToNodeP(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("OPTIONS", "/api/v1/nodes/"+name+"/proxy/"+path+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy OPTIONS Request To A Node(with path) Fail", err)
	}
	logger.Info("Proxy OPTIONS node names/"+name+"/path/"+path, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyPostReqToNodeP(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("POST", "/api/v1/nodes/"+name+"/proxy/"+path+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy Post Request To A Node(with path) Fail", err)
	}
	logger.Info("Proxy Post node names/"+name+"/path/"+path, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyHeadReqToNodeP(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("HEAD", "/api/v1/nodes/"+name+"/proxy/"+path+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy Head Request To A Node(with path) Fail", err)
	}
	logger.Info("Proxy Head node names/"+name+"/path/"+path, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyGetReqToNodeP(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("GET", "/api/v1/nodes/"+name+"/proxy/"+path+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy Get Request To A Node(with path) Fail", err)
	}
	logger.Info("Proxy Get node names/"+name+"/path/"+path, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyPutReqToNodeP(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("PUT", "/api/v1/nodes/"+name+"/proxy/"+path+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy Put Request To A Node(with path) Fail", err)
	}
	logger.Info("Proxy Put node names/"+name+"/path/"+path, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyPatchReqToNodeP(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("PATCH", "/api/v1/nodes/"+name+"/proxy/"+path+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy Patch Request To A Node(with path) Fail", err)
	}
	logger.Info("Proxy Patch node names/"+name+"/path/"+path, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ProxyDelReqToNodeP(c *gin.Context) {
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	req, err := oapi.GenRequest("DELETE", "/api/v1/nodes/"+name+"/proxy/"+path+urlParas, token, nil)
	if err != nil {
		logger.Error("Proxy Delete Request To A Node(with path) Fail", err)
	}
	logger.Info("Proxy Delete node names/"+name+"/path/"+path, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}
