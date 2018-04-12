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
	logger.Info("Create limitrange",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
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
	logger.Info("Create limitrange namespaces/"+namespace,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
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
	logger.Info("Get limitrange namespaces/"+namespace+"/names/"+name,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
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
	logger.Info("List limitrange",map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
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
	logger.Info("List limitrange namespaces/"+namespace,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func watchLimitRangeNS(c *gin.Context){

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	logger.Info("Watch limitrange namespaces/"+namespace+"/names/"+name,map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(),"result":"begin"})
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/limitranges/"+name,token,c.Writer,c.Request)
	logger.Info("Watch limitrange namespaces/"+namespace+"/names/"+name,map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(),"result":"end"})

}

func watchAllLimitRanges(c *gin.Context){

	token := pkg.GetWSToken(c)
	logger.Info("Watch collection limitrange",map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(),"result":"begin"})
	oapi.WSRequest("/api/v1/watch/limitranges",token,c.Writer,c.Request)
	logger.Info("Watch collection limitrange",map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(),"result":"end"})

}

func watchAllLimitRangesNS(c *gin.Context){

	token := pkg.GetWSToken(c)
	namespace := c.Param("namespace")
	logger.Info("Watch collection limitrange namespaces/"+namespace,map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(),"result":"begin"})
	oapi.WSRequest("/api/v1/watch/namespaces/"+namespace+"/limitranges",token,c.Writer,c.Request)
	logger.Info("Watch collection limitrange namespaces/"+namespace,map[string]interface{}{"user": pkg.GetUserFromToken(token), "time": pkg.GetTimeNow(),"result":"end"})

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
	logger.Info("Update limitrange namespaces/"+namespace+"/names/"+name,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
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
	logger.Info("Patch limitrange namespaces/"+namespace+"/names/"+name,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
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
	logger.Info("Delete limitrange namespaces/"+namespace+"/names/"+name,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
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
	logger.Info("Delete collection limitrange namespaces/"+namespace,map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(),"result":req.StatusCode})
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}
