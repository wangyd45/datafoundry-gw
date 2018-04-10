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
	req,err := oapi.GenRequest("POST","/oapi/v1/roles",token,rBody)
	if err != nil{
		logger.Error("Create A Role Fail",err)
	}
	logger.Info("Create role",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func CreateRoleInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("POST","/oapi/v1/namespaces/"+namespace+"/roles",token,rBody)
	if err != nil{
		logger.Error("Create A Role In A Namespace Fail",err)
	}
	logger.Info("Create role namespaces/"+namespace,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)

}

func GetRoleInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/oapi/v1/namespaces/"+namespace+"/roles/"+name,token,nil)
	if err != nil{
		logger.Error("Get A Role In A Namespace Fail",err)
	}
	logger.Info("Get role namespaces/"+namespace+"/names/"+name,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetAllRoles(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET","/oapi/v1/roles",token,nil)
	if err != nil{
		logger.Error("Get All Roles Fail",err)
	}
	logger.Info("List roles",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetRolesInNS(c *gin.Context)  {
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req,err := oapi.GenRequest("GET","/oapi/v1/namespaces/"+namespace+"/roles",token,nil)
	if err != nil{
		logger.Error("Get All Roles In A Namespace Fail",err)
	}
	logger.Info("List roles namespaces/"+namespace,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func UpdateRoleInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/oapi/v1/namespaces/"+namespace+"/roles/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A Role In A Namespace Fail",err)
	}
	logger.Info("Update role namespaces/"+namespace+"/names/"+name,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchRoleInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/oapi/v1/namespaces/"+namespace+"/roles/"+name,token,rBody)
	if err != nil{
		logger.Error("Patch A Role In A Namespace Fail",err)
	}
	logger.Info("Patch role namespaces/"+namespace+"/names/"+name,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteRoleInNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("DELETE","/oapi/v1/namespaces/"+namespace+"/roles/"+name,token,rBody)
	if err != nil{
		logger.Error("Delete A Role In A Namespace Fail",err)
	}
	logger.Info("Delete role namespaces/"+namespace+"/names/"+name,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}