package persistentvolume

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
	logger = lager.NewLogger("api_v1_PersistentVolume")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CreatePV(c *gin.Context){
	token := pkg.GetToken(c)
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	//调用原生接口
	req,err := oapi.GenRequest("POST","/api/v1/persistentvolumes",token,rBody)
	if err != nil{
		logger.Error("Create A PersistentVolume Fail",err)
	}
	//返回结果JSON格式
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GorWPV(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchPV(c)
	}else{
		getPV(c)
	}
}

func GorWAllPVs(c *gin.Context){
	if pkg.IsWebsocket(c){
		watchAllPVs(c)
	}else{
		getAllPVs(c)
	}
}

func getPV(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/api/v1/persistentvolumes/"+name,token,nil)
	if err != nil{
		logger.Error("Get A PersistentVolume Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func getAllPVs(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("GET","/api/v1/persistentvolumes",token,nil)
	if err != nil{
		logger.Error("Get All PersistentVolumes Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func watchPV(c *gin.Context){

	token := pkg.GetWSToken(c)
	name := c.Param("name")
	oapi.WSRequest("/api/v1/watch/persistentvolumes/"+name,token,c.Writer,c.Request)

}

func watchAllPVs(c *gin.Context){

	token := pkg.GetWSToken(c)
	oapi.WSRequest("/api/v1/watch/persistentvolumes",token,c.Writer,c.Request)

}

func UpdatePV(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/api/v1/persistentvolumes/"+name,token,rBody)
	if err != nil{
		logger.Error("Update A PersistentVolume Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchPV(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/api/v1/persistentvolumes/"+name,token,rBody)
	if err != nil{
		logger.Error("Patch A PersistentVolume Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeletePV(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("DELETE","/api/v1/persistentvolumes/"+name,token,rBody)
	if err != nil{
		logger.Error("Delete A PersistentVolume Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func DeleteAllPVs(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.GenRequest("DELETE","/api/v1/persistentvolumes",token,nil)
	if err != nil{
		logger.Error("Delete All PersistentVolumes Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func GetstatusofPV(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	req,err := oapi.GenRequest("GET","/api/v1/persistentvolumes/"+name+"/status",token,nil)
	if err != nil{
		logger.Error("Get status of a PersistentVolume Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func UpdatestatusofPV(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PUT","/api/v1/persistentvolumes/"+name+"/status",token,rBody)
	if err != nil{
		logger.Error("Update status of a PersistentVolume Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

func PatchstatusofPV(c *gin.Context){
	token := pkg.GetToken(c)
	name := c.Param("name")
	rBody,_ := ioutil.ReadAll(c.Request.Body)
	req,err := oapi.GenRequest("PATCH","/api/v1/persistentvolumes/"+name+"/status",token,rBody)
	if err != nil{
		logger.Error("Patch status of a PersistentVolume Fail",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json",result)
}

//v1.PersistentVolume
//router.POST("/api/v1/persistentvolumes", route.CreatePV)
//router.GET("/api/v1/persistentvolumes/:name", route.GorWPV)
//router.GET("/api/v1/persistentvolumes", route.GorWAllPVs)
//router.PUT("/api/v1/persistentvolumes/:name", route.UpdatePV)
//router.PATCH("/api/v1/persistentvolumes/:name", route.PatchPV)
//router.DELETE("/api/v1/persistentvolumes/:name", route.DeletePV)
//router.DELETE("/api/v1/persistentvolumes", route.DeleteAllPVs)
//router.GET("/api/v1/persistentvolumes/:name/status", route.GetstatusofPV)
//router.PUT("/api/v1/persistentvolumes/:name/status", route.UpdatestatusofPV)
//router.PATCH("/api/v1/persistentvolumes/:name/status", route.PatchstatusofPV)