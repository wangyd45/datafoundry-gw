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

func CreateRole(c *gin.Context)  {

}

func CreateRoleInNS(c *gin.Context)  {

}

func GetRoleInNS(c *gin.Context){

}

func GetAllRoles(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/projects",token,nil)
	if err != nil{
		logger.Error("Get All Projects Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetRolesInNS(c *gin.Context)  {

}

func UpdateRoleInNS(c *gin.Context){

}

func PatchRoleInNS(c *gin.Context){

}

func DeleteRoleInNS(c *gin.Context){

}