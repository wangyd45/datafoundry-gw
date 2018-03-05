package template

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
	logger = lager.NewLogger("oapi_v1_Template")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateTemplate(c *gin.Context){
	token := pkg.GetToken(c)
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("POST","/oapi/v1/templates",token,rBody)
	if err != nil{
		logger.Error("Create A Template Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func CreateTemplatenNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("POST","/oapi/v1/namespaces/"+namespace+"/templates",token,rBody)
	if err != nil{
		logger.Error("Create A Template In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GorWTemplateInNS(c *gin.Context) {

	if pkg.IsWebsocket(c){
		watchTemplateInNS(c)
	}else{
		getTemplateInNS(c)
	}
}

func GorWAllTemplates(c *gin.Context) {

	if pkg.IsWebsocket(c){
		watchAllTemplates(c)
	}else{
		getAllTemplates(c)
	}
}

func GorWAllTemplatesInNS(c *gin.Context) {

	if pkg.IsWebsocket(c){
		watchAllTemplatesInNS(c)
	}else{
		getAllTemplatesInNS(c)
	}
}

func getTemplateInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/oapi/v1/namespaces/"+namespace+"/templates/"+name,token,nil)
	if err != nil{
		logger.Error("Get A Template In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func getAllTemplates(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET","/oapi/v1/templates",token,nil)
	if err != nil{
		logger.Error("Get All Templates Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func getAllTemplatesInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req,err := oapi.GenRequest("GET","/oapi/v1/namespaces/"+namespace+"/templates",token,nil)
	if err != nil{
		logger.Error("Get All Templates In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func watchTemplateInNS(c *gin.Context){

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	oapi.WSRequest("/oapi/v1/watch/namespaces/"+namespace+"/templates/"+name,token,c.Writer,c.Request)

}

func watchAllTemplates(c *gin.Context){

	token := pkg.GetWSToken(c)
	oapi.WSRequest("/oapi/v1/watch/templates",token,c.Writer,c.Request)

}

func watchAllTemplatesInNS(c *gin.Context){

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	oapi.WSRequest("/oapi/v1/watch/namespaces/"+namespace+"/templates",token,c.Writer,c.Request)

}

func UpdateTemplateInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/oapi/v1/namespaces/"+namespace+"/templates/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A Template In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchTemplateInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/oapi/v1/namespaces/"+namespace+"/templates/"+name,token,rBody)
	if err != nil{
		logger.Error("Patch A Template In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteTemplateInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("DELETE","/oapi/v1/namespaces/"+namespace+"/templates/"+name,token,rBody)
	if err != nil{
		logger.Error("Delete A Template In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteAllTemplatesInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req,err := oapi.GenRequest("DELETE","/oapi/v1/namespaces/"+namespace+"/templates",token,nil)
	if err != nil{
		logger.Error("Delete All Templates In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}