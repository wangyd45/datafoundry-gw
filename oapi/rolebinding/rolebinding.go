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

func CreateRoleBinding(c *gin.Context)  {

}

func CreateRoleBindingInNS(c *gin.Context)  {

}

func GetRoleBindingInNS(c *gin.Context){

}

func GetAllRoleBindings(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/projects",token,nil)
	if err != nil{
		logger.Error("Get All Projects Fail",err)
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
