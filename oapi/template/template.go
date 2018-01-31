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

func CreateTemplate(c *gin.Context)  {

}

func CreateTemplatenNS(c *gin.Context)  {

}

func GetTemplateInNS(c *gin.Context){

}

func GetAllTemplates(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/projects",token,nil)
	if err != nil{
		logger.Error("Get All Projects Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetAllTemplatesInNS(c *gin.Context){

}

func WatchTemplateInNS(c *gin.Context){

}

func WatchAllTemplates(c *gin.Context){

}

func WatchAllTemplatesInNS(c *gin.Context){

}

func UpdateTemplateInNS(c *gin.Context){

}

func PatchTemplateInNS(c *gin.Context){

}

func DeleteTemplateInNS(c *gin.Context){

}

func DeleteAllTemplatesInNS(c *gin.Context){

}