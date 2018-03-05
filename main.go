package main

import (

    "os"
    "fmt"
    "time"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/pivotal-golang/lager"
	"github.com/asiainfoLDP/datafoundry-gw/oapi/build"
	"github.com/asiainfoLDP/datafoundry-gw/oapi/buildconfig"
	dep "github.com/asiainfoLDP/datafoundry-gw/oapi/deploymentconfig"
	image "github.com/asiainfoLDP/datafoundry-gw/oapi/imagestream"
	"github.com/asiainfoLDP/datafoundry-gw/oapi/imagestreamimport"
	tag "github.com/asiainfoLDP/datafoundry-gw/oapi/imagestreamtag"
	"github.com/asiainfoLDP/datafoundry-gw/oapi/netnamespace"
	"github.com/asiainfoLDP/datafoundry-gw/oapi/project"
	"github.com/asiainfoLDP/datafoundry-gw/oapi/role"
	"github.com/asiainfoLDP/datafoundry-gw/oapi/rolebinding"
	"github.com/asiainfoLDP/datafoundry-gw/oapi/route"
	"github.com/asiainfoLDP/datafoundry-gw/oapi/template"
	"github.com/asiainfoLDP/datafoundry-gw/oapi/user"
/*

	"log"
	"golang.org/x/net/websocket"
	"fmt"
*/
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
	s := &http.Server{
		Addr:           ":10012",
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 0,
	}
	//监听端口
	s.ListenAndServe()
}

func handle() (router *gin.Engine) {
	//设置全局环境：1.开发环境（gin.DebugMode） 2.线上环境（gin.ReleaseMode）
	gin.SetMode(gin.DebugMode)
	//获取路由实例
	router = gin.Default()

	//v1.user
	//router.POST("/users",user.CreateUser)
	router.GET("/oapi/v1/users/:name", user.GetUser)
	router.GET("/oapi/v1/users", user.GetAllUser)
	//router.GET("/watch/users/:name",user.WatchUser)
	//router.GET("watch/users",user.WatchAllUser)
	//router.PUT("/users/:name",user.UpdataUser)
	//router.PATCH("/users/:name",user.PatchUser)
	router.DELETE("/oapi/v1/users/:name", user.DeleteUser)
	//router.DELETE("/users/",user.DeleteAllUser)

	//v1.project
	router.POST("/oapi/v1/projectrequests", project.CreateProject)
	//router.GET("/oapi/v1/projects/:name", project.GetProject)
	router.GET("/oapi/v1/projects/:name", project.GorWProject)
	router.GET("/oapi/v1/projects", project.GorWAllProjects)
	//router.GET("/oapi/v1/projects", project.GetAllProjects)
	//router.GET("/oapi/v1/watch/projects/:name",project.WatchAProject)
	//router.GET("/oapi/v1/watch/projects",project.WatchAllProjects)
	router.PUT("/oapi/v1/projects/:name", project.UpdateProject)
	router.PATCH("/oapi/v1/projects/:name",project.PatchAProject)
	router.DELETE("/oapi/v1/projects/:name", project.DeleteProject)

	//v1.Build NS -> NameSpace
	router.POST("/oapi/v1/builds", build.CreateBuild)
	router.POST("/oapi/v1/namespaces/:namespace/builds", build.CreateBuildInNS)
	router.POST("/oapi/v1/namespaces/:namespace/builds/:name/clone", build.CreateCloneInNS)
	router.GET("/oapi/v1/namespaces/:namespace/builds/:name", build.GetBuildFromNS)
	router.GET("/oapi/v1/builds", build.GetAllBuilds)
	router.GET("/oapi/v1/namespaces/:namespace/builds", build.GetAllBuildFromNS)
	router.GET("/oapi/v1/namespaces/:namespace/builds/:name/log", build.GetLogBuildFromNS)
	router.GET("/oapi/v1/watch/namespaces/:namespace/builds/:name", build.WatchBuildFromNS)
	router.GET("/oapi/v1/watch/builds", build.WatchAllBuilds)
	router.GET("/oapi/v1/watch/namespaces/:namespace/builds", build.WatchAllBuildFromNS)
	router.PUT("/oapi/v1/namespaces/:namespace/builds/:name", build.UpdataBuildFromNS)
	router.PUT("/oapi/v1/namespaces/:namespace/builds/:name/details", build.UpdataDetailsInNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/builds/:name", build.PatchBuildFromNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/builds/:name", build.DeleteBuildFromNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/builds", build.DeleteAllBuildFromNS)

	//v1.BuildConfig BC -> BuildConfig NS -> NameSpace NSP -> NameSpacePath
	router.POST("/oapi/v1/buildconfigs", buildconfig.CreateBC)
	router.POST("/oapi/v1/namespaces/:namespace/buildconfigs", buildconfig.CreateBCInNS)
	router.POST("/oapi/v1/namespaces/:namespace/buildconfigs/:name/instantiate", buildconfig.CreateInsInNS)
	router.POST("/oapi/v1/namespaces/:namespace/buildconfigs/:name/instantiatebinary", buildconfig.CreateInstInNS)
	router.POST("/oapi/v1/namespaces/:namespace/buildconfigs/:name/webhooks", buildconfig.CreateWebInNS)
	router.POST("/oapi/v1/namespaces/:namespace/buildconfigs/:name/webhooks/:path", buildconfig.CreateWebInNSP)
	router.GET("/oapi/v1/namespaces/:namespace/buildconfigs/:name", buildconfig.GetBCFromNS)
	router.GET("/oapi/v1/buildconfigs", buildconfig.GetAllBC)
	router.GET("/oapi/v1/namespaces/:namespace/buildconfigs", buildconfig.GetAllBCFromNS)
	router.GET("/oapi/v1/watch/namespaces/:namespace/buildconfigs/:name", buildconfig.WatchBCFromNS)
	router.GET("/oapi/v1/watch/buildconfigs", buildconfig.WatchAllBC)
	router.GET("/oapi/v1/watch/namespaces/:namespace/buildconfigs", buildconfig.WatchAllBCFromNS)
	router.PUT("/oapi/v1/namespaces/:namespace/buildconfigs/:name", buildconfig.UpdataBCFromNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/buildconfigs/:name", buildconfig.PatchBCFromNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/buildconfigs/:name", buildconfig.DeleteBCFromNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/buildconfigs", buildconfig.DeleteAllBuildFromNS)

	//v1.DeploymentConfig DC -> DeploymentConfig Dep -> Deployment
	router.POST("/oapi/v1/deploymentconfigs", dep.CreateDC)
	router.POST("/oapi/v1/namespaces/:namespace/deploymentconfigs", dep.CreateDCInNS)
	router.POST("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name/instantiate", dep.CreateInsInNS)
	router.POST("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name/rollback", dep.CreateRollBackInNS)
	//router.POST("/web/deploymentconfig/:namespace/:name", buildconfig.CreateWebInNS)
	//router.POST("/web/deploymentconfig/:namespace/:name/:path", buildconfig.CreateWebInNSP)
	router.GET("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name", dep.GetDCFromNS)
	router.GET("/oapi/v1/deploymentconfigs", dep.GetAllDC)
	router.GET("/oapi/v1/namespaces/:namespace/deploymentconfigs", dep.GetAllDCFromNS)
	router.GET("/oapi/v1/watch/namespaces/:namespace/deploymentconfigs/:name", dep.WatchDCFromNS)
	router.GET("/oapi/v1/watch/deploymentconfigs", dep.WatchAllDC)
	router.GET("/oapi/v1/watch/namespaces/:namespace/deploymentconfigs", dep.WatchAllDCFromNS)
	router.GET("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name/log", dep.GetLogDepFromNS)
	router.GET("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name/scale", dep.GetScaleDepFromNS)
	router.GET("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name/status", dep.GetStatusDepFromNS)
	router.PUT("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name", dep.UpdataDCFromNS)
	router.PUT("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name/scale", dep.UpdataScaleDCFromNS)
	router.PUT("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name/status", dep.UpdataStatusDCFromNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name", dep.PatchDCFromNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name/scale", dep.PatchScaleDCFromNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name/status", dep.PatchStatusDCFromNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name", dep.DeleteDCFromNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/deploymentconfigs", dep.DeleteAllDepFromNS)

	//v1.ImageStream IS -> ImageStream NS -> NameSpace SecImage ->SecretsImage
	router.POST("/oapi/v1/imagestreams", image.CreateIS)
	router.POST("/oapi/v1/namespaces/:namespace/imagestreams", image.CreateImageInNS)
	router.GET("/oapi/v1/namespaces/:namespace/imagestreams/:name", image.GetImageFromNS)
	router.GET("/oapi/v1/imagestreams", image.GetAllImage)
	router.GET("/oapi/v1/namespaces/:namespace/imagestreams", image.GetAllImageFromNS)
	router.GET("/oapi/v1/namespaces/:namespace/imagestreams/:name/secrets", image.GetSecImageFromNS)
	router.GET("/oapi/v1/namespaces/:namespace/imagestreams/:name/status", image.GetStaImageFromNS)
	router.GET("/oapi/v1/watch/namespaces/:namespace/imagestreams/:name", image.WatchImageFromNS)
	router.GET("/oapi/v1/watch/imagestreams", image.WatchAllImage)
	router.GET("/oapi/v1/watch/namespaces/:namespace/imagestreams", image.WatchAllImageFromNS)
	router.PUT("/oapi/v1/namespaces/:namespace/imagestreams/:name", image.UpdataImageFromNS)
	router.PUT("/oapi/v1/namespaces/:namespace/imagestreams/:name/status", image.UpdataStaImageFromNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/imagestreams/:name", image.PatchImageFromNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/imagestreams/:name/status", image.PatchStaImageFromNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/imagestreams/:name", image.DeleteImageFromNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/imagestreams", image.DeleteAllImageFromNS)

	//v1.ImageStreamImport IS -> ImageStream NS -> NameSpace
	router.POST("/oapi/v1/imagestreamimports", imagestreamimport.CreateISImport)
	router.POST("/oapi/v1/namespaces/:namespace/imagestreamimports", imagestreamimport.CreateISImportInNS)

	//v1.ImageStreamTag IST -> ImageStreamTag NS -> NameSpace
	router.POST("/oapi/v1/imagestreamtags", tag.CreateIST)
	router.POST("/oapi/v1/namespaces/:namespace/imagestreamtags", tag.CreateImageTagInNS)
	router.GET("/oapi/v1/namespaces/:namespace/imagestreamtags/:name", tag.GetImageTagFromNS)
	router.GET("/oapi/v1/imagestreamtags", tag.GetAllImageTag)
	router.GET("/oapi/v1/namespaces/:namespace/imagestreamtags", tag.GetAllImageTagFromNS)
	router.PUT("/oapi/v1/namespaces/:namespace/imagestreamtags/:name", tag.UpdataImageTagFromNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/imagestreamtags/:name", tag.PatchImageTagFromNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/imagestreamtags/:name", tag.DeleteImageTagFromNS)

	//v1.NetNamespace
	router.POST("/oapi/v1/netnamespaces", netnamespace.CreateNetNamespace)
	router.GET("/oapi/v1/netnamespaces/:name", netnamespace.GorWNetNamespace)
	router.GET("/oapi/v1/netnamespaces", netnamespace.GorWAllNetNamespaces)
	//router.GET("/oapi/v1/netnamespaces/:name", netnamespace.GetNetNamespace)
	//router.GET("/oapi/v1/netnamespaces", netnamespace.GetAllNetNamespaces)
	//router.GET("/oapi/v1/watch/netnamespaces/:name", netnamespace.WatchNetNamespace)
	//router.GET("/oapi/v1/watch/netnamespaces", netnamespace.WatchAllNetNamespaces)

	router.PUT("/oapi/v1/netnamespaces/:name", netnamespace.UpdateNetNamespace)
	router.PATCH("/oapi/v1/netnamespaces/:name", netnamespace.PatchNetNamespace)
	router.DELETE("/oapi/v1/netnamespaces/:name", netnamespace.DeleteNetNamespace)
	router.DELETE("/oapi/v1/netnamespaces", netnamespace.DeleteAllNetNamespaces)

	//v1.Role
	router.POST("/oapi/v1/roles", role.CreateRole)
	router.POST("/oapi/v1/namespaces/:namespace/roles", role.CreateRoleInNS)
	router.GET("/oapi/v1/namespaces/:namespace/roles/:name", role.GetRoleInNS)
	router.GET("/oapi/v1/roles", role.GetAllRoles)
	router.GET("/oapi/v1/namespaces/:namespace/roles", role.GetRolesInNS)
	router.PUT("/oapi/v1/namespaces/:namespace/roles/:name", role.UpdateRoleInNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/roles/:name", role.PatchRoleInNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/roles/:name", role.DeleteRoleInNS)

	//v1.RoleBinding
	router.POST("/oapi/v1/rolebindings", rolebinding.CreateRoleBinding)
	router.POST("/oapi/v1/namespaces/:namespace/rolebindings", rolebinding.CreateRoleBindingInNS)
	router.GET("/oapi/v1/namespaces/:namespace/rolebindings/:name", rolebinding.GetRoleBindingInNS)
	router.GET("/oapi/v1/rolebindings", rolebinding.GetAllRoleBindings)
	router.GET("/oapi/v1/namespaces/:namespace/rolebindings", rolebinding.GetRoleBindingsInNS)
	router.PUT("/oapi/v1/namespaces/:namespace/rolebindings/:name", rolebinding.UpdateRoleBindingInNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/rolebindings/:name", rolebinding.PatchRoleBindingInNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/rolebindings/:name", rolebinding.DeleteRoleBindingInNS)

	//v1.Route
	router.POST("/oapi/v1/routes", route.CreateRoute)
	router.POST("/oapi/v1/namespaces/:namespace/routes", route.CreateRouteInNS)
	router.GET("/oapi/v1/namespaces/:namespace/routes/:name", route.GorWRouteInNS)
	router.GET("/oapi/v1/routes", route.GorWAllRoutes)
	router.GET("/oapi/v1/namespaces/:namespace/routes", route.GorWAllRoutesInNS)
	//router.GET("/oapi/v1/namespaces/:namespace/routes/:name", route.GetRouteInNS)
	//router.GET("/oapi/v1/routes", route.GetAllRoutes)
	//router.GET("/oapi/v1/namespaces/:namespace/routes", route.GetAllRoutesInNS)
	//router.GET("/oapi/v1/watch/namespaces/:namespace/routes/:name", route.WatchRouteInNS)
	//router.GET("/oapi/v1/watch/routes", route.WatchAllRoutes)
	//router.GET("/oapi/v1/watch/namespaces/:namespace/routes", route.WatchAllRoutesInNS)
	router.GET("/oapi/v1/namespaces/:namespace/routes/:name/status", route.GetRouteStatusInNS)
	router.PUT("/oapi/v1/namespaces/:namespace/routes/:name", route.UpdateRouteInNS)
	router.PUT("/oapi/v1/namespaces/:namespace/routes/:name/status", route.UpdateRouteStatusInNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/routes/:name", route.PatchRouteInNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/routes/:name/status", route.PatchRouteStatusInNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/routes/:name", route.DeleteRouteInNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/routes", route.DeleteAllRoutesInNS)

	//v1.Template
	router.POST("/oapi/v1/templates", template.CreateTemplate)
	router.POST("/oapi/v1/namespaces/:namespace/templates", template.CreateTemplatenNS)
	router.GET("/oapi/v1/namespaces/:namespace/templates/:name", template.GorWTemplateInNS)
	router.GET("/oapi/v1/templates", template.GorWAllTemplates)
	router.GET("/oapi/v1/namespaces/:namespace/templates", template.GorWAllTemplatesInNS)
	//router.GET("/oapi/v1/namespaces/:namespace/templates/:name", template.GetTemplateInNS)
	//router.GET("/oapi/v1/templates", template.GetAllTemplates)
	//router.GET("/oapi/v1/namespaces/:namespace/templates", template.GetAllTemplatesInNS)
	//router.GET("/oapi/v1/watch/namespaces/:namespace/templates/:name", template.WatchTemplateInNS)
	//router.GET("/oapi/v1/watch/templates", template.WatchAllTemplates)
	//router.GET("/oapi/v1/watch/namespaces/:namespace/templates", template.WatchAllTemplatesInNS)
	router.PUT("/oapi/v1/namespaces/:namespace/templates/:name", template.UpdateTemplateInNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/templates/:name", template.PatchTemplateInNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/templates/:name", template.DeleteTemplateInNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/templates", template.DeleteAllTemplatesInNS)

	return
}