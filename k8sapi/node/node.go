package node

import (
	"os"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"
)

var logger lager.Logger

func init() {
	logger = lager.NewLogger("api_v1_Node")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateNode(c *gin.Context){
	token := pkg.GetToken(c)
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req,err := oapi.GenRequest("POST","/api/v1/nodes",token,rBody)
	if err != nil{
		logger.Error("Create A Node Fail",err)
	}
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GorWNode(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchNode(c)
	}else{
		getNode(c)
	}
}

func GorWAllNodes(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchAllNodes(c)
	}else{
		getAllNodes(c)
	}
}

func getNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/api/v1/nodes/"+name,token,nil)
	if err != nil{
		logger.Error("Get A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func getAllNodes(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET","/api/v1/nodes",token,nil)
	if err != nil{
		logger.Error("Get All Nodes Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func watchNode(c *gin.Context){

	token := pkg.GetWSToken(c)
	name := c.Param("name")
	oapi.WSRequest("/api/v1/watch/nodes/"+name,token,c.Writer,c.Request)

}

func watchAllNodes(c *gin.Context){

	token := pkg.GetWSToken(c)
	oapi.WSRequest("/api/v1/watch/nodes",token,c.Writer,c.Request)

}

func UpdateNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/api/v1/nodes/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/api/v1/nodes/"+name,token,rBody)
	if err != nil{
		logger.Error("Patch A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("DELETE","/api/v1/nodes/"+name,token,rBody)
	if err != nil{
		logger.Error("Delete A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteAllNodes(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("DELETE","/api/v1/nodes",token,nil)
	if err != nil{
		logger.Error("Delete All Nodes Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetStatusOfNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/api/v1/nodes/"+name+"/status",token,nil)
	if err != nil{
		logger.Error("Get Status Of A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func UpdateStatusOfNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/api/v1/nodes/"+name+"/status",token,rBody)
	if err != nil{
		logger.Error("Update Status Of A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchStatusOfNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/api/v1/nodes/"+name+"/status",token,rBody)
	if err != nil{
		logger.Error("Patch Status Of A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyOpnReqToNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("OPTIONS","/api/v1/nodes/"+name+"/proxy",token,nil)
	if err != nil{
		logger.Error("Proxy OPTIONS Request To A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyPostReqToNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("POST","/api/v1/nodes/"+name+"/proxy",token,nil)
	if err != nil{
		logger.Error("Proxy Post Request To A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyHeadReqToNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("HEAD","/api/v1/nodes/"+name+"/proxy",token,nil)
	if err != nil{
		logger.Error("Proxy Head Request To A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyGetReqToNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/api/v1/nodes/"+name+"/proxy",token,nil)
	if err != nil{
		logger.Error("Proxy Get Request To A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyPutReqToNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("PUT","/api/v1/nodes/"+name+"/proxy",token,nil)
	if err != nil{
		logger.Error("Proxy Put Request To A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyPatchReqToNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("PATCH","/api/v1/nodes/"+name+"/proxy",token,nil)
	if err != nil{
		logger.Error("Proxy Patch Request To A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyDelReqToNode(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("DELETE","/api/v1/nodes/"+name+"/proxy",token,nil)
	if err != nil{
		logger.Error("Proxy Delete Request To A Node Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyOpnReqToNodeP(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	req,err := oapi.GenRequest("OPTIONS","/api/v1/nodes/"+name+"/proxy/"+path,token,nil)
	if err != nil{
		logger.Error("Proxy OPTIONS Request To A Node(with path) Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyPostReqToNodeP(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	req,err := oapi.GenRequest("POST","/api/v1/nodes/"+name+"/proxy/"+path,token,nil)
	if err != nil{
		logger.Error("Proxy Post Request To A Node(with path) Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyHeadReqToNodeP(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	req,err := oapi.GenRequest("HEAD","/api/v1/nodes/"+name+"/proxy/"+path,token,nil)
	if err != nil{
		logger.Error("Proxy Head Request To A Node(with path) Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyGetReqToNodeP(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	req,err := oapi.GenRequest("GET","/api/v1/nodes/"+name+"/proxy/"+path,token,nil)
	if err != nil{
		logger.Error("Proxy Get Request To A Node(with path) Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyPutReqToNodeP(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	req,err := oapi.GenRequest("PUT","/api/v1/nodes/"+name+"/proxy/"+path,token,nil)
	if err != nil{
		logger.Error("Proxy Put Request To A Node(with path) Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyPatchReqToNodeP(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	req,err := oapi.GenRequest("PATCH","/api/v1/nodes/"+name+"/proxy/"+path,token,nil)
	if err != nil{
		logger.Error("Proxy Patch Request To A Node(with path) Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func ProxyDelReqToNodeP(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	path := c.Param("path")
	req,err := oapi.GenRequest("DELETE","/api/v1/nodes/"+name+"/proxy/"+path,token,nil)
	if err != nil{
		logger.Error("Proxy Delete Request To A Node(with path) Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}
