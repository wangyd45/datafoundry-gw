package main

import (
    "os"
    "fmt"
    "time"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/pivotal-golang/lager"
    "github.com/asiainfoLDP/datafoundry-gw/oapi/user"
    "github.com/asiainfoLDP/datafoundry-gw/oapi/project"
)

//定义日志以及其他变量
var logger lager.Logger
var ocAPIName string = "ocapi"

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

func main() {
    //初始化日志对象，日志输出到stdout
    logger = lager.NewLogger(ocAPIName)
    logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.INFO)) //默认日志级别
    router := handle()
    s :=  &http.Server{
        Addr:           ":10012",
        Handler:        router,
        ReadTimeout:    30*time.Second,
        WriteTimeout:   30*time.Second,
        MaxHeaderBytes: 0,
    }
    //监听端口
    s.ListenAndServe()
}

func handle()(router *gin.Engine){
    //设置全局环境：1.开发环境（gin.DebugMode） 2.线上环境（gin.ReleaseMode）
    gin.SetMode(gin.DebugMode)
    //获取路由实例
    router = gin.Default()

    //v1.user
    router.POST("/users",user.CreateUser)
    router.GET("/users",user.GetUser)
    router.GET("/users/getAllUser",user.GetAllUser)
    router.GET("/users/watchUser",user.WatchUser)
    router.GET("/users/watchAllUser",user.WatchAllUser)
    router.PUT("/users/updataUser",user.UpdataUser)
    router.PATCH("/users/patchUser",user.PatchUser)
    router.DELETE("/users/deleteUser",user.DeleteUser)
    router.DELETE("/users/deleteAllUser",user.DeleteAllUser)
    //router.GET("/api/v1/namespace/:name/users/:user")

    //v1.project
    router.POST("/projects",project.CreateAProject)
    router.GET("/projects/:name",project.GetAProject)
    router.GET("/projects",project.GetAllProjects)
    router.GET("/watch/projects/:name",project.WatchAProject)
    router.GET("/watch/projects",project.WatchAllProjects)
    router.PUT("/projects/:name",project.UpdateAProject)
    router.PATCH("/projects/:name",project.PatchAProject)
    router.DELETE("/projects/:name",project.DeleteAProject)

    return
}


