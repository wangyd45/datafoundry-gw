package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/pivotal-golang/lager"
    "os"
    ocapi "github.com/asiainfoLDP/datafoundry-gw/handler"
)

//定义日志以及其他变量
var logger lager.Logger
var ocAPIName string = "ocapi"


func main() {
    //设置全局环境：1.开发环境（gin.DebugMode） 2.线上环境（gin.ReleaseMode）
    gin.SetMode(gin.DebugMode)

    //获取路由实例
    router := gin.Default()

    //添加中间件
    //router.Use(Middleware)

    //初始化日志对象，日志输出到stdout
    logger = lager.NewLogger(ocAPIName)
    logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.INFO)) //默认日志级别

    //注册接口实例
    //router.GET("/simple/server/get", GetHandler)
    //router.POST("/simple/server/post", PostHandler)
    //router.PUT("/simple/server/put", PutHandler)
    //router.DELETE("/simple/server/delete", DeleteHandler)

    //测试接口
    router.GET("/projects",ocapi.GetProjects)
    router.GET("/projects/:project", GetaProject)
    //接口列表及注册接口


    //监听端口
    http.ListenAndServe(":10000", router)
}

//获取环境变量
func getenv(env string) string {
    env_value := os.Getenv(env)
    if env_value == "" {
        fmt.Println("FATAL: NEED ENV", env)
        fmt.Println("Exit...........")
        os.Exit(2)
    }
    fmt.Println("ENV:", env, env_value)
    return env_value
}


func GetaProject(c *gin.Context){
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
