package role

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
	logger = lager.NewLogger("oapi_v1_Role")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateRole(c *gin.Context){
	token := pkg.GetToken(c)
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"POST","/oapi/v1/roles",token,rBody)
	if err != nil{
		logger.Error("Create A Role Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func CreateRoleInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"POST","/oapi/v1/namespaces/"+namespace+"/roles",token,rBody)
	if err != nil{
		logger.Error("Create A Role In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)

}

func GetRoleInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req,err := oapi.Request(10,"GET","/oapi/v1/namespaces/"+namespace+"/roles/"+name,token,nil)
	if err != nil{
		logger.Error("Get A Role In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetAllRoles(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/roles",token,nil)
	if err != nil{
		logger.Error("Get All Roles Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetRolesInNS(c *gin.Context)  {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req,err := oapi.Request(10,"GET","/oapi/v1/namespaces/"+namespace+"/roles",token,nil)
	if err != nil{
		logger.Error("Get All Roles In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func UpdateRoleInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"PUT","/oapi/v1/namespaces/"+namespace+"/roles/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A Role In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchRoleInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"PATCH","/oapi/v1/namespaces/"+namespace+"/roles/"+name,token,rBody)
	if err != nil{
		logger.Error("Patch A Role In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteRoleInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.Request(10,"DELETE","/oapi/v1/namespaces/"+namespace+"/roles/"+name,token,rBody)
	if err != nil{
		logger.Error("Delete A Role In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}