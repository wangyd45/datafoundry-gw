package rolebinding

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
	logger = lager.NewLogger("oapi_v1_RoleBinding")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateRoleBinding(c *gin.Context){
	token := pkg.GetToken(c)
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"POST","/oapi/v1/rolebindings",token,rBody)
	if err != nil{
		logger.Error("Create A RoleBinding Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func CreateRoleBindingInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"POST","/oapi/v1/namespaces/"+namespace+"/rolebindings",token,rBody)
	if err != nil{
		logger.Error("Create A RoleBinding In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetRoleBindingInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req,err := oapi.Request(10,"GET","/oapi/v1/namespaces/"+namespace+"/rolebindings/"+name,token,nil)
	if err != nil{
		logger.Error("Get A RoleBinding In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetAllRoleBindings(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/rolebindings",token,nil)
	if err != nil{
		logger.Error("Get All RoleBindings Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetRoleBindingsInNS(c *gin.Context)  {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req,err := oapi.Request(10,"GET","/oapi/v1/namespaces/"+namespace+"/rolebindings",token,nil)
	if err != nil{
		logger.Error("Get All RoleBindings In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func UpdateRoleBindingInNS(c *gin.Context){

}

func PatchRoleBindingInNS(c *gin.Context){

}

func DeleteRoleBindingInNS(c *gin.Context){

}
