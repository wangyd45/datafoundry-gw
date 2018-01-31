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

func CreateNetNamespace(c *gin.Context)  {

}

func GetNetNamespace(c *gin.Context){

}

func GetAllNetNamespaces(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/projects",token,nil)
	if err != nil{
		logger.Error("Get All Projects Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func WatchNetNamespace(c *gin.Context){

}

func WatchAllNetNamespaces(c *gin.Context){

}

func UpdateNetNamespace(c *gin.Context){

}

func PatchNetNamespace(c *gin.Context){

}

func DeleteNetNamespace(c *gin.Context){

}

func DeleteAllNetNamespaces(c *gin.Context){

}