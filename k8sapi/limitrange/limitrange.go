package limitrange

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
	logger = lager.NewLogger("api_v1_LimitRange")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreateLimitRange(c *gin.Context){
	token := pkg.GetToken(c)
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req,err := oapi.GenRequest("POST","/api/v1/limitranges",token,rBody)
	if err != nil{
		logger.Error("Create A LimitRange Fail",err)
	}
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func CreateLimitRangeNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req,err := oapi.GenRequest("POST","/api/v1/namespaces/"+namespace+"/limitranges",token,rBody)
	if err != nil{
		logger.Error("Create A LimitRange In A Namespace Fail",err)
	}
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GorWLimitRangeNS(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchLimitRangeNS(c)
	}else{
		getLimitRangeNS(c)
	}
}

func GorWAllLimitRanges(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchAllLimitRanges(c)
	}else{
		getAllLimitRanges(c)
	}
}

func GorWAllLimitRangesNS(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchAllLimitRangesNS(c)
	}else{
		getAllLimitRangesNS(c)
	}
}

func getLimitRangeNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/api/v1/namespaces/"+namespace+"/limitranges/"+name,token,nil)
	if err != nil{
		logger.Error("Get A LimitRange In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func getAllLimitRanges(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET","/api/v1/limitranges",token,nil)
	if err != nil{
		logger.Error("Get All LimitRanges Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func getAllLimitRangesNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req,err := oapi.GenRequest("GET","/api/v1/namespaces/"+namespace+"/limitranges",token,nil)
	if err != nil{
		logger.Error("Get All LimitRanges In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func watchLimitRangeNS(c *gin.Context){

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/limitranges/"+name,token,c.Writer,c.Request)

}

func watchAllLimitRanges(c *gin.Context){

	token := pkg.GetWSToken(c)
	oapi.WSRequest("/api/v1/watch/limitranges",token,c.Writer,c.Request)

}

func watchAllLimitRangesNS(c *gin.Context){

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/limitranges",token,c.Writer,c.Request)

}

func UpdateLimitRangeNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/api/v1/namespaces/"+namespace+"/limitranges/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A LimitRange In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchLimitRangeNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/api/v1/namespaces/"+namespace+"/limitranges/"+name,token,rBody)
	if err != nil{
		logger.Error("Patch A LimitRange In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteLimitRangeNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("DELETE","/api/v1/namespaces/"+namespace+"/limitranges/"+name,token,rBody)
	if err != nil{
		logger.Error("Delete A LimitRange In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteAllLimitRangeNS(c *gin.Context){
	token := pkg.GetToken(c)
	namespace := c.Param("namespace")
	req,err := oapi.GenRequest("DELETE","/api/v1/namespaces/"+namespace+"/limitranges",token,nil)
	if err != nil{
		logger.Error("Delete All LimitRange In A Namespace Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

//v1.LimitRange
//router.POST("/api/v1/limitranges", route.CreateLimitRange)
//router.POST("/api/v1/namespaces/:namespace/limitranges", route.CreateLimitRangeNS)
//router.GET("/api/v1/namespaces/:namespace/limitranges/:name", route.GorWLimitRangeNS)
//router.GET("/api/v1/limitranges", route.GorWAllLimitRanges)
//router.GET("/api/v1/namespaces/:namespace/limitranges", route.GorWAllLimitRangesNS)
//router.PUT("/api/v1/namespaces/:namespace/limitranges/:name", route.UpdateLimitRangeNS)
//router.PATCH("/api/v1/namespaces/:namespace/limitranges/:name", route.PatchLimitRangeNS)
//router.DELETE("/api/v1/namespaces/:namespace/limitranges/:name", route.DeleteLimitRangeNS)
//router.DELETE("/api/v1/namespaces/:namespace/limitranges", route.DeleteAllLimitRangeNS)