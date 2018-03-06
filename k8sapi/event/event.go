package event

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
	logger = lager.NewLogger("api_v1_Event")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateEvent(c *gin.Context){
	token := pkg.GetToken(c)
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req,err := oapi.GenRequest("POST","/api/v1/events",token,rBody)
	if err != nil{
		logger.Error("Create A Event Fail",err)
	}
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func CreateEventNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req,err := oapi.GenRequest("POST","/api/v1/namespaces/"+namespace+"/events",token,rBody)
	if err != nil{
		logger.Error("Create A Event In A Namespace Fail",err)
	}
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GorWEventNS(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchEventNS(c)
	}else{
		getEventNS(c)
	}
}

func GorWAllEvents(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchAllEvents(c)
	}else{
		getAllEvents(c)
	}
}

func GorWAllEventsNS(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchAllEventsNS(c)
	}else{
		getAllEventsNS(c)
	}
}

func getEventNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/api/v1/namespaces/"+namespace+"/events/"+name,token,nil)
	if err != nil{
		logger.Error("Get A Event In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func getAllEvents(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET","/api/v1/events",token,nil)
	if err != nil{
		logger.Error("Get All Events Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func getAllEventsNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req,err := oapi.GenRequest("GET","/api/v1/namespaces/"+namespace+"/events",token,nil)
	if err != nil{
		logger.Error("Get All Events In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func watchEventNS(c *gin.Context){

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/events/"+name,token,c.Writer,c.Request)

}

func watchAllEvents(c *gin.Context){

	token := pkg.GetWSToken(c)
	oapi.WSRequest("/api/v1/watch/events",token,c.Writer,c.Request)

}

func watchAllEventsNS(c *gin.Context){

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/events",token,c.Writer,c.Request)

}

func UpdateEventNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/api/v1/namespaces/"+namespace+"/events/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A Event In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchEventNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/api/v1/namespaces/"+namespace+"/events/"+name,token,rBody)
	if err != nil{
		logger.Error("Patch A Event In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteEventNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("DELETE","/api/v1/namespaces/"+namespace+"/events/"+name,token,rBody)
	if err != nil{
		logger.Error("Delete A Event In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteAllEventNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req,err := oapi.GenRequest("DELETE","/api/v1/namespaces/"+namespace+"/events",token,nil)
	if err != nil{
		logger.Error("Delete All Event In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

//v1.Event
//router.POST("/api/v1/events", route.CreateEvent)
//router.POST("/api/v1/namespaces/:namespace/events", route.CreateEventNS)
//router.GET("/api/v1/namespaces/:namespace/events/:name", route.GorWEventNS)
//router.GET("/api/v1/events", route.GorWAllEvents)
//router.GET("/api/v1/namespaces/:namespace/events", route.GorWAllEventsNS)
//router.PUT("/api/v1/namespaces/:namespace/events/:name", route.UpdateEventNS)
//router.PATCH("/api/v1/namespaces/:namespace/events/:name", route.PatchEventNS)
//router.DELETE("/api/v1/namespaces/:namespace/events/:name", route.DeleteEventNS)
//router.DELETE("/api/v1/namespaces/:namespace/events", route.DeleteAllEventNS)