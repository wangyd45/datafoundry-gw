package endpoints

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
	logger = lager.NewLogger("api_v1_EndPoints")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateEndpoints(c *gin.Context){
	token := pkg.GetToken(c)
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req,err := oapi.GenRequest("POST","/api/v1/endpoints",token,rBody)
	if err != nil{
		logger.Error("Create A Endpoints Fail",err)
	}
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func CreateEndpointsNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req,err := oapi.GenRequest("POST","/api/v1/namespaces/"+namespace+"/endpoints",token,rBody)
	if err != nil{
		logger.Error("Create A Endpoints In A Namespace Fail",err)
	}
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetEndpointsNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/api/v1/namespaces/"+namespace+"/endpoints/"+name,token,nil)
	if err != nil{
		logger.Error("Get A Endpoints In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetAllEndpoints(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET","/api/v1/endpoints",token,nil)
	if err != nil{
		logger.Error("Get All Endpoints Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetAllEndpointsNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req,err := oapi.GenRequest("GET","/api/v1/namespaces/"+namespace+"/endpoints",token,nil)
	if err != nil{
		logger.Error("Get All Endpoints In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func WatchEndpointsNS(c *gin.Context){

	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/endpoints/"+name,token,c.Writer,c.Request)

}

func WatchAllEndpoints(c *gin.Context){

	token := pkg.GetToken(c)
	oapi.WSRequest("/api/v1/watch/endpoints",token,c.Writer,c.Request)

}

func WatchAllEndpointsNS(c *gin.Context){

	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/endpoints",token,c.Writer,c.Request)

}

func UpdateEndpointsNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/api/v1/namespaces/"+namespace+"/endpoints/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A Endpoints In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchEndpointsNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/api/v1/namespaces/"+namespace+"/endpoints/"+name,token,rBody)
	if err != nil{
		logger.Error("Patch A Endpoints In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteEndpointsNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("DELETE","/api/v1/namespaces/"+namespace+"/endpoints/"+name,token,rBody)
	if err != nil{
		logger.Error("Delete A Endpoints In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteAllEndpointsNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("DELETE","/api/v1/namespaces/"+namespace+"/endpoints",token,rBody)
	if err != nil{
		logger.Error("Delete All Endpoints In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}
