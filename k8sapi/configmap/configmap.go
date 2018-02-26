package configmap

import (
	"github.com/pivotal-golang/lager"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
)

/*
import (
	"os"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"
	"fmt"
)

var logger lager.Logger

func init() {
	logger = lager.NewLogger("api_v1_ConfigMap")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}


func GetAllProjects(c *gin.Context){
token := pkg.GetToken(c)
req,err := oapi.GenRequest("GET","/oapi/v1/projects",token,nil)
if err != nil{
logger.Error("Get All Projects Fail",err)
}
result, _:= ioutil.ReadAll(req.Body)
defer req.Body.Close()
c.Data(req.StatusCode, "application/json",result)
}

func WatchAProject(c *gin.Context) {

token := pkg.GetToken(c)
name := c.Param("name")
oapi.WSRequest("/oapi/v1/watch/projects/"+name,token,c.Writer,c.Request)
}

Create a ConfigMap
Create a ConfigMap in a namespace
Get a ConfigMap in a namespace
Get all ConfigMaps
Get all ConfigMaps in a namespace
Watch a ConfigMap in a namespace
Watch all ConfigMaps
Watch all ConfigMaps in a namespace
Update a ConfigMap in a namespace
Patch a ConfigMap in a namespace
Delete a ConfigMap in a namespace
Delete all ConfigMaps in a namespace