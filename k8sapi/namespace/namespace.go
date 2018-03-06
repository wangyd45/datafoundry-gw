package namespace

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
	logger = lager.NewLogger("api_v1_Namespace")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateNamespace(c *gin.Context){
	token := pkg.GetToken(c)
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req,err := oapi.GenRequest("POST","/api/v1/namespaces",token,rBody)
	if err != nil{
		logger.Error("Create A Namespace Fail",err)
	}
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GorWNamespace(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchNamespace(c)
	}else{
		getNamespace(c)
	}
}

func GorWAllNamespaces(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchAllNamespaces(c)
	}else{
		getAllNamespaces(c)
	}
}

func getNamespace(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/api/v1/namespaces/"+name,token,nil)
	if err != nil{
		logger.Error("Get A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func getAllNamespaces(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET","/api/v1/namespaces",token,nil)
	if err != nil{
		logger.Error("Get All Namespaces Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func watchNamespace(c *gin.Context){

	token := pkg.GetWSToken(c)
	name := c.Param("name")
	oapi.WSRequest("/api/v1/watch/namespaces/"+name,token,c.Writer,c.Request)

}

func watchAllNamespaces(c *gin.Context){

	token := pkg.GetWSToken(c)
	oapi.WSRequest("/api/v1/watch/namespaces",token,c.Writer,c.Request)

}

func UpdateNamespace(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/api/v1/namespaces/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchNamespace(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/api/v1/namespaces/"+name,token,rBody)
	if err != nil{
		logger.Error("Patch A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteNamespace(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("DELETE","/api/v1/namespaces/"+name,token,rBody)
	if err != nil{
		logger.Error("Delete A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func UpdatefinalizeofNS(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/api/v1/namespaces/"+name+"/finalize",token,rBody)
	if err != nil{
		logger.Error("Update finalize of a Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetstatusofNS(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/api/v1/namespaces/"+name+"/status",token,nil)
	if err != nil{
		logger.Error("Get status of a Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func UpdatestatusofNS(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/api/v1/namespaces/"+name+"/status",token,rBody)
	if err != nil{
		logger.Error("Update status of a Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchstatusofNS(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/api/v1/namespaces/"+name+"/status",token,rBody)
	if err != nil{
		logger.Error("Patch status of a Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

//v1.Namespace
//router.POST("/api/v1/namespaces", route.CreateNamespace)
//router.GET("/api/v1/namespaces/:name", route.GorWNamespace)
//router.GET("/api/v1/namespaces", route.GorWAllNamespaces)
//router.PUT("/api/v1/namespaces/:name", route.UpdateNamespace)
//router.PATCH("/api/v1/namespaces/:name", route.PatchNamespace)
//router.DELETE("/api/v1/namespaces/:name", route.DeleteNamespace)
//router.PUT("/api/v1/namespaces/:name/finalize", route.UpdatefinalizeofNS)
//router.GET("/api/v1/namespaces/:name/status", route.GetstatusofNS)
//router.PUT("/api/v1/namespaces/:name/status", route.UpdatestatusofNS)
//router.PATCH("/api/v1/namespaces/:name/status", route.PatchstatusofNS)