package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    //configapi "github.com/openshift/origin/pkg/cmd/server/api"
)

func main() {
    //设置全局环境：1.开发环境（gin.DebugMode） 2.线上环境（gin.ReleaseMode）
    gin.SetMode(gin.DebugMode)

    //获取路由实例
    router := gin.Default()

    //添加中间件
    //router.Use(Middleware)

    //注册接口
    //router.GET("/simple/server/get", GetHandler)
    //router.POST("/simple/server/post", PostHandler)
    //router.PUT("/simple/server/put", PutHandler)
    //router.DELETE("/simple/server/delete", DeleteHandler)

    router.GET("/oapi/v1/users",GetUsers)
    //router.GET("/oapi/v1/projects",GetProjects)
    router.GET("/oapi/v1/users/:name", GetaUser)



    //监听端口
    http.ListenAndServe(":10000", router)
}

func GetUsers(c *gin.Context){

    c.Data(http.StatusOK, "json", []byte(fmt.Sprintf("get success!")))
    //c.JSON(http.StatusOK, gin.H{"message":"wch" , "status": http.StatusOK})
    return
}

func GetaUser(c *gin.Context){
    name := c.Param("name")
    if "" == name{
        c.JSON(http.StatusExpectationFailed, gin.H{"message": "param is nil", "status": http.StatusExpectationFailed})
    }
    say := "hello " + name
    c.JSON(http.StatusOK, gin.H{"message":say , "status": http.StatusOK})
}
/*
func GetProjects(c *gin.Context){
    return
}
*/

/*
func Middleware(c *gin.Context) {
    fmt.Println("this is a middleware!")
}

func GetHandler(c *gin.Context) {
    value, exist := c.GetQuery("key")
    if !exist {
        value = "the key is not exist!"
    }
    c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
    return
}
func PostHandler(c *gin.Context) {
    type JsonHolder struct {
        Id   int    `json:"id"`
        Name string `json:"name"`
    }
    holder := JsonHolder{Id: 1, Name: "my name"}
    //若返回json数据，可以直接使用gin封装好的JSON方法
    c.JSON(http.StatusOK, holder)
    return
}
func PutHandler(c *gin.Context) {
    c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
    return
}
func DeleteHandler(c *gin.Context) {
    c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
    return
}
*/
