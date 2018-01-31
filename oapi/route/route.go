package route

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
	logger = lager.NewLogger("oapi_v1_Route")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateRoute(c *gin.Context)  {

}

func CreateRouteInNS(c *gin.Context)  {

}

func GetRouteInNS(c *gin.Context){

}

func GetAllRoutes(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/projects",token,nil)
	if err != nil{
		logger.Error("Get All Projects Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetAllRoutesInNS(c *gin.Context){

}

func WatchRouteInNS(c *gin.Context){

}

func WatchAllRoutes(c *gin.Context){

}

func WatchAllRoutesInNS(c *gin.Context){

}

func UpdateRouteInNS(c *gin.Context){

}

func PatchRouteInNS(c *gin.Context){

}

func DeleteRouteInNS(c *gin.Context){

}

func DeleteAllRoutesInNS(c *gin.Context){

}

func GetRouteStatusInNS(c *gin.Context){

}

func UpdateRouteStatusInNS(c *gin.Context){

}

func PatchRouteStatusInNS(c *gin.Context){

}