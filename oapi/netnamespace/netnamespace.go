package netnamespace

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
	logger = lager.NewLogger("oapi_v1_NetNamespace")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateNetNamespace(c *gin.Context){
	token := pkg.GetToken(c)
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"POST","/oapi/v1/netnamespaces",token,rBody)
	if err != nil{
		logger.Error("Create A NetNamespaces Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetNetNamespace(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.Request(10,"GET","/oapi/v1/netnamespaces/"+name,token,nil)
	if err != nil{
		logger.Error("Get A NetNamespaces Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetAllNetNamespaces(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/netnamespaces",token,nil)
	if err != nil{
		logger.Error("Get All NetNamespaces Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func WatchNetNamespace(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.Request(10,"GET","/oapi/v1/watch/netnamespaces/"+name,token,nil)
	if err != nil{
		logger.Error("Watch A NetNamespaces Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func WatchAllNetNamespaces(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/watch/netnamespaces",token,nil)
	if err != nil{
		logger.Error("Watch All NetNamespaces Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)

}

func UpdateNetNamespace(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"PUT","/oapi/v1/netnamespaces/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A NetNamespaces Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)

}

func PatchNetNamespace(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"PATCH","/oapi/v1/netnamespaces/"+name,token,rBody)
	if err != nil{
		logger.Error("Patch A NetNamespaces Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteNetNamespace(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"DELETE","/oapi/v1/netnamespaces/"+name,token,rBody)
	if err != nil{
		logger.Error("Delete A NetNamespaces Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteAllNetNamespaces(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"DELETE","/oapi/v1/netnamespaces",token,nil)
	if err != nil{
		logger.Error("Delete All NetNamespaces Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}