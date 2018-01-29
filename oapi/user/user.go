package user

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"

)

func CreateUser(c *gin.Context){
	parm := "{ \"kind\": \"User\",\"apiVersion\": \"v1\"}"
	req,err := oapi.Request(10,"POST","/oapi/v1/users","JlzsISKq7ZGY4W51KfL3o1GbqHBgoQ4D6dgsV3DkRho",[]byte(parm))
	if err != nil{
		fmt.Println("CreateUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func GetUser(c *gin.Context){
	str := "sb-instanceid-anaconda3-nt9z3"
	req,err := oapi.Request(10,"GET","/oapi/v1/users","JlzsISKq7ZGY4W51KfL3o1GbqHBgoQ4D6dgsV3DkRho",[]byte(str))
	if err != nil{
		fmt.Println("----- CetUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	c.JSON(http.StatusOK, gin.H{"status": string(result)})
}

func GetAllUser(c *gin.Context){
	token := pkg.GetToken(c)
	req,err := oapi.Request(10,"GET","/oapi/v1/users",token,[]byte{})
	if err != nil{
		fmt.Println("------  CetAllUser error ",err)
	}
	result, _:= ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	//c.JSON(http.StatusOK,req.Body)
	c.Data(http.StatusOK,"ok",result)
}

func WatchUser(c *gin.Context){
	//name := c.Param("name")
	//if "" == name{
	//	c.JSON(http.StatusExpectationFailed, gin.H{"message": "param is nil", "status": http.StatusExpectationFailed})
	//}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func WatchAllUser(c *gin.Context){
	//name := c.Param("name")
	//if "" == name{
	//	c.JSON(http.StatusExpectationFailed, gin.H{"message": "param is nil", "status": http.StatusExpectationFailed})
	//}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func UpdataUser(c *gin.Context){
	//name := c.Param("name")
	//if "" == name{
	//	c.JSON(http.StatusExpectationFailed, gin.H{"message": "param is nil", "status": http.StatusExpectationFailed})
	//}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func PatchUser(c *gin.Context){
	//name := c.Param("name")
	//if "" == name{
	//	c.JSON(http.StatusExpectationFailed, gin.H{"message": "param is nil", "status": http.StatusExpectationFailed})
	//}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func DeleteUser(c *gin.Context){
	//name := c.Param("name")
	//if "" == name{
	//	c.JSON(http.StatusExpectationFailed, gin.H{"message": "param is nil", "status": http.StatusExpectationFailed})
	//}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func DeleteAllUser(c *gin.Context){
	//name := c.Param("name")
	//if "" == name{
	//	c.JSON(http.StatusExpectationFailed, gin.H{"message": "param is nil", "status": http.StatusExpectationFailed})
	//}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
