package configmap

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
	logger = lager.NewLogger("api_v1_ConfigMap")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateConfigMap(c *gin.Context){
	token := pkg.GetToken(c)
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req,err := oapi.GenRequest("POST","/api/v1/configmaps",token,rBody)
	if err != nil{
		logger.Error("Create A ConfigMap Fail",err)
	}
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func CreateConfigMapNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req,err := oapi.GenRequest("POST","/api/v1/namespaces/"+namespace+"/configmaps",token,rBody)
	if err != nil{
		logger.Error("Create A ConfigMap In A Namespace Fail",err)
	}
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetConfigMapNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/api/v1/namespaces/"+namespace+"/configmaps/"+name,token,nil)
	if err != nil{
		logger.Error("Get A ConfigMap In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetAllConfigMap(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET","/api/v1/configmaps",token,nil)
	if err != nil{
		logger.Error("Get All ConfigMap Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetAllConfigMapNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req,err := oapi.GenRequest("GET","/api/v1/namespaces/"+namespace+"/configmaps",token,nil)
	if err != nil{
		logger.Error("Get All ConfigMap In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func WatchConfigMapNS(c *gin.Context){

	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/configmaps/"+name,token,c.Writer,c.Request)

}

func WatchAllConfigMap(c *gin.Context){

	token := pkg.GetToken(c)
	oapi.WSRequest("/api/v1/watch/configmaps",token,c.Writer,c.Request)

}

func WatchAllConfigMapNS(c *gin.Context){

	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/configmaps",token,c.Writer,c.Request)

}

func UpdateConfigMapNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/api/v1/namespaces/"+namespace+"/configmaps/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A ConfigMap In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchConfigMapNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/api/v1/namespaces/"+namespace+"/configmaps/"+name,token,rBody)
	if err != nil{
		logger.Error("Patch A ConfigMap In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteConfigMapNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("DELETE","/api/v1/namespaces/"+namespace+"/configmaps/"+name,token,rBody)
	if err != nil{
		logger.Error("Delete A ConfigMap In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteAllConfigMapNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("DELETE","/api/v1/namespaces/"+namespace+"/configmaps",token,rBody)
	if err != nil{
		logger.Error("Delete All ConfigMap In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}