package main

import (
    "os"
    "fmt"
    "time"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/pivotal-golang/lager"
    "github.com/asiainfoLDP/datafoundry-gw/oapi/user"
    "github.com/asiainfoLDP/datafoundry-gw/oapi/build"
    "github.com/asiainfoLDP/datafoundry-gw/oapi/project"
    "github.com/asiainfoLDP/datafoundry-gw/oapi/buildconfig"
    dep "github.com/asiainfoLDP/datafoundry-gw/oapi/deploymentconfig"
    "github.com/asiainfoLDP/datafoundry-gw/oapi/netnamespace"
    "github.com/asiainfoLDP/datafoundry-gw/oapi/role"
    "github.com/asiainfoLDP/datafoundry-gw/oapi/rolebinding"
    "github.com/asiainfoLDP/datafoundry-gw/oapi/route"
    "github.com/asiainfoLDP/datafoundry-gw/oapi/template"
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
    //router.POST("/users",user.CreateUser)
    router.GET("/users/:name",user.GetUser)
    router.GET("/users",user.GetAllUser)
    //router.GET("/watch/users/:name",user.WatchUser)
    //router.GET("watch/users",user.WatchAllUser)
    //router.PUT("/users/:name",user.UpdataUser)
    //router.PATCH("/users/:name",user.PatchUser)
    router.DELETE("/users/:name",user.DeleteUser)
    //router.DELETE("/users/",user.DeleteAllUser)

    //v1.project
    router.POST("/projects",project.CreateProject)
    router.GET("/projects/:name",project.GetProject)
    router.GET("/projects",project.GetAllProjects)
    //router.GET("/watch/projects/:name",project.WatchAProject)
    //router.GET("/watch/projects",project.WatchAllProjects)
    router.PUT("/projects/:name",project.UpdateProject)
    //router.PATCH("/projects/:name",project.PatchAProject)
    router.DELETE("/projects/:name",project.DeleteProject)

    //v1.Build
    router.POST("/build",build.CreateBuild)
    router.POST("/build/:namespace",build.CreateBuildInNameSpace)
    router.POST("/clone/build/:namespace/:name",build.CreateCloneInNameSpace)
    router.GET("/build/:namespace/:name",build.GetBuildFromNameSpace)
    router.GET("/build",build.GetAllBuilds)
    router.GET("/build/:namespace",build.GetAllBuildFromNameSpace)
    router.GET("/log/build/:namespace/:name",build.GetLogBuildFromNameSpace)
    router.GET("/watch/build/:namespace/:name",build.WatchBuildFromNameSpace)
    router.GET("/watch/build",build.WatchAllBuilds)
    router.GET("/watch/build/:namespace",build.WatchAllBuildFromNameSpace)
    router.PUT("/build/:namespace/:name",build.UpdataBuildFromNameSpace)
    router.PUT("/detail/build/:namespace/:name",build.UpdataDetailsInNameSpace)
    router.PATCH("/build/:namespace/:name",build.PatchBuildFromNameSpace)
    router.DELETE("/build/:namespace/:name",build.DeleteBuildFromNameSpace)
    router.DELETE("/build/:namespace",build.DeleteAllBuildFromNameSpace)

    //v1.BuildConfig
    router.POST("/buildconfig",buildconfig.CreateBuildConfig)
    router.POST("/buildconfig/:namespace",buildconfig.CreateBuildConfigInNameSpace)
    router.POST("/ins/buildconfig/:namespace/:name",buildconfig.CreateInsInNameSpace)
    router.POST("/inst/buildconfig/:namespace/:name",buildconfig.CreateInstInNameSpace)
    router.POST("/web/buildconfig/:namespace/:name",buildconfig.CreateWebInNameSpace)
    router.POST("/web/buildconfig/:namespace/:name/:path",buildconfig.CreateWebInNameSpacePath)
    router.GET("/buildconfig/:namespace/:name",buildconfig.GetBuildConfigFromNameSpace)
    router.GET("/buildconfig",buildconfig.GetAllBuildConfig)
    router.GET("/buildconfig/:namespace",buildconfig.GetAllBuildConfigFromNameSpace)
    router.GET("/watch/buildconfig/:namespace/:name",buildconfig.WatchBuildConfigFromNameSpace)
    router.GET("/watch/buildconfig",buildconfig.WatchAllBuildConfig)
    router.GET("/watch/buildconfig/:namespace",buildconfig.WatchAllBuildConfigFromNameSpace)
    router.PUT("/buildconfig/:namesapce/:name",buildconfig.UpdataBuildConfigFromNameSpace)
    router.PATCH("/buildconfig/:namesapce/:name",buildconfig.PatchBuildConfigFromNameSpace)
    router.DELETE("/buildconfig/:namesapce/:name",buildconfig.DeleteBuildConfigFromNameSpace)
    router.DELETE("/buildconfig/:namesapce",buildconfig.DeleteAllBuildFromNameSpace)

    //v1.DeploymentConfig
    router.POST("/deploymentconfig",dep.CreateDeploymentConfig)
    router.POST("/deploymentconfig/:namespace",dep.CreateDeploymentConfigInNameSpace)
    router.POST("/ins/deploymentconfig/:namespace/:name",dep.CreateInsInNameSpace)
    router.POST("/roolback/deploymentconfig/:namespace/:name",dep.CreateRollBackInNameSpace)
    router.POST("/web/deploymentconfig/:namespace/:name",buildconfig.CreateWebInNameSpace)
    router.POST("/web/deploymentconfig/:namespace/:name/:path",buildconfig.CreateWebInNameSpacePath)
    router.GET("/deploymentconfig/:namespace/:name",dep.GetDeploymentConfigFromNameSpace)
    router.GET("/deploymentconfig",dep.GetAllDeploymentConfig)
    router.GET("/deploymentconfig/:namespace",dep.GetAllDeploymentConfigFromNameSpace)
    router.GET("/watch/deploymentconfig/:namespace/:name",dep.WatchDeploymentConfigFromNameSpace)
    router.GET("/watch/deploymentconfig",dep.WatchAllDeploymentConfig)
    router.GET("/watch/deploymentconfig/:namespace",dep.WatchAllDeploymentConfigFromNameSpace)
    router.GET("/log/deploymentconfig/:namespace/:name",dep.GetLogDeploymentFromNameSpace)
    router.GET("/scale/deploymentconfig/:namespace/:name",dep.GetScaleDeploymentFromNameSpace)
    router.GET("/status/deploymentconfig/:namespace/:name",dep.GetStatusDeploymentFromNameSpace)
    router.PUT("/deploymentconfig/:namesapce/:name",dep.UpdataDeploymentConfigFromNameSpace)
    router.PUT("/scale/deploymentconfig/:namesapce/:name",dep.UpdataScaleDeploymentConfigFromNameSpace)
    router.PUT("/status/deploymentconfig/:namesapce/:name",dep.UpdataStatusDeploymentConfigFromNameSpace)
    router.PATCH("/deploymentconfig/:namesapce/:name",dep.PatchDeploymentConfigFromNameSpace)
    router.PATCH("/scale/deploymentconfig/:namesapce/:name",dep.PatchScaleDeploymentConfigFromNameSpace)
    router.PATCH("/status/deploymentconfig/:namesapce/:name",dep.PatchStatusDeploymentConfigFromNameSpace)
    router.DELETE("/deploymentconfig/:namesapce/:name",dep.DeleteDeploymentConfigFromNameSpace)
    router.DELETE("/deploymentconfig/:namesapce",dep.DeleteAllDeploymentFromNameSpace)

    //v1.NetNamespace
    router.POST("/netnamespaces",netnamespace.CreateNetNamespace)
    router.GET("/netnamespaces/:name",netnamespace.GetNetNamespace)
    router.GET("/netnamespaces",netnamespace.GetAllNetNamespaces)
    router.GET("/watch/netnamespaces/:name",netnamespace.WatchNetNamespace)
    router.GET("/watch/netnamespaces",netnamespace.WatchAllNetNamespaces)
    router.PUT("/netnamespaces/:name",netnamespace.UpdateNetNamespace)
    router.PATCH("/netnamespaces/:name",netnamespace.PatchNetNamespace)
    router.DELETE("/netnamespaces/:name",netnamespace.DeleteNetNamespace)
    router.DELETE("/netnamespaces",netnamespace.DeleteAllNetNamespaces)

    //v1.Role
    router.POST("/roles",role.CreateRole)
    router.POST("/namespaces/:namespace/roles",role.CreateRoleInNS)
    router.GET("/namespaces/:namespace/roles/:name",role.GetRoleInNS)
    router.GET("/roles",role.GetAllRoles)
    router.GET("/namespaces/:namespace/roles",role.GetRolesInNS)
    router.PUT("/namespaces/:namespace/roles/:name",role.UpdateRoleInNS)
    router.PATCH("/namespaces/:namespace/roles/:name",role.PatchRoleInNS)
    router.DELETE("/namespaces/:namespace/roles/:name",role.DeleteRoleInNS)

    //v1.RoleBinding
    router.POST("/rolebindings",rolebinding.CreateRoleBinding)
    router.POST("/namespaces/:namespace/rolebindings",rolebinding.CreateRoleBindingInNS)
    router.GET("/namespaces/:namespace/rolebindings/:name",rolebinding.GetRoleBindingInNS)
    router.GET("/rolebindings",rolebinding.GetAllRoleBindings)
    router.GET("/namespaces/:namespace/rolebindings",rolebinding.GetRoleBindingInNS)
    router.PUT("/namespaces/:namespace/rolebindings/:name",rolebinding.UpdateRoleBindingInNS)
    router.PATCH("/namespaces/:namespace/rolebindings/:name",rolebinding.PatchRoleBindingInNS)
    router.DELETE("/namespaces/:namespace/rolebindings/:name",rolebinding.DeleteRoleBindingInNS)

    //v1.Route
    router.POST("/routes",route.CreateRoute)
    router.POST("/namespaces/:namespace/routes",route.CreateRouteInNS)
    router.GET("/namespaces/:namespace/routes/:name",route.GetRouteInNS)
    router.GET("/routes",route.GetAllRoutes)
    router.GET("/namespaces/:namespace/routes",route.GetAllRoutesInNS)
    router.GET("/watch/namespaces/:namespace/routes/:name",route.WatchRouteInNS)
    router.GET("/watch/routes",route.WatchAllRoutes)
    router.GET("/watch/namespaces/:namespace/routes",route.WatchAllRoutesInNS)
    router.PUT("/namespaces/:namespace/routes/:name",route.UpdateRouteInNS)
    router.PATCH("/namespaces/:namespace/routes/:name",route.PatchRouteInNS)
    router.DELETE("/namespaces/:namespace/routes/:name",route.DeleteRouteInNS)
    router.DELETE("/namespaces/:namespace/routes",route.DeleteAllRoutesInNS)
    router.GET("namespaces/:namespace/routes/:name/status",route.GetRouteStatusInNS)
    router.PUT("/namespaces/:namespace/routes/:name/status",route.UpdateRouteStatusInNS)
    router.PATCH("/namespaces/:namespace/routes/:name/status",route.PatchRouteStatusInNS)

    //v1.Template
    router.POST("/templates",template.CreateTemplate)
    router.POST("/namespaces/:namespace/templates",template.CreateTemplatenNS)
    router.GET("/namespaces/:namespace/templates/:name",template.GetTemplateInNS)
    router.GET("/templates",template.GetAllTemplates)
    router.GET("/namespaces/:namespace/templates",template.GetAllTemplatesInNS)
    router.GET("/watch/namespaces/:namespace/templates/:name",template.WatchTemplateInNS)
    router.GET("/watch/templates",template.WatchAllTemplates)
    router.GET("/watch/namespaces/:namespace/templates",template.WatchAllTemplatesInNS)
    router.PUT("/namespaces/:namespace/templates/:name",template.UpdateTemplateInNS)
    router.PATCH("/namespaces/:namespace/templates/:name",template.PatchTemplateInNS)
    router.DELETE("/namespaces/:namespace/templates/:name",template.DeleteTemplateInNS)
    router.DELETE("/namespaces/:namespace/templates",template.DeleteAllTemplatesInNS)


    return
}


