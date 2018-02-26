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
	router.GET("/users/:name", user.GetUser)
	router.GET("/users", user.GetAllUser)
	//router.GET("/watch/users/:name",user.WatchUser)
	//router.GET("watch/users",user.WatchAllUser)
	//router.PUT("/users/:name",user.UpdataUser)
	//router.PATCH("/users/:name",user.PatchUser)
	router.DELETE("/users/:name", user.DeleteUser)
	//router.DELETE("/users/",user.DeleteAllUser)

	//v1.project
	router.POST("/projects", project.CreateProject)
	router.GET("/projects/:name", project.GetProject)
	router.GET("/projects", project.GetAllProjects)
	router.GET("/watch/projects/:name",project.WatchAProject)
	router.GET("/watch/projects",project.WatchAllProjects)
	router.PUT("/projects/:name", project.UpdateProject)
	router.PATCH("/projects/:name",project.PatchAProject)
	router.DELETE("/projects/:name", project.DeleteProject)

	//v1.Build NS -> NameSpace
	router.POST("/build", build.CreateBuild)
	router.POST("/build/:namespace", build.CreateBuildInNS)
	router.POST("/clone/build/:namespace/:name", build.CreateCloneInNS)
	router.GET("/build/:namespace/:name", build.GetBuildFromNS)
	router.GET("/build", build.GetAllBuilds)
	router.GET("/build/:namespace", build.GetAllBuildFromNS)
	router.GET("/log/build/:namespace/:name", build.GetLogBuildFromNS)
	router.GET("/watch/build/:namespace/:name", build.WatchBuildFromNS)
	router.GET("/watch/build", build.WatchAllBuilds)
	router.GET("/watch/build/:namespace", build.WatchAllBuildFromNS)
	router.PUT("/build/:namespace/:name", build.UpdataBuildFromNS)
	router.PUT("/detail/build/:namespace/:name", build.UpdataDetailsInNS)
	router.PATCH("/build/:namespace/:name", build.PatchBuildFromNS)
	router.DELETE("/build/:namespace/:name", build.DeleteBuildFromNS)
	router.DELETE("/build/:namespace", build.DeleteAllBuildFromNS)

	//v1.BuildConfig BC -> BuildConfig NS -> NameSpace NSP -> NameSpacePath
	router.POST("/buildconfig", buildconfig.CreateBC)
	router.POST("/buildconfig/:namespace", buildconfig.CreateBCInNS)
	router.POST("/ins/buildconfig/:namespace/:name", buildconfig.CreateInsInNS)
	router.POST("/inst/buildconfig/:namespace/:name", buildconfig.CreateInstInNS)
	router.POST("/web/buildconfig/:namespace/:name", buildconfig.CreateWebInNS)
	router.POST("/web/buildconfig/:namespace/:name/:path", buildconfig.CreateWebInNSP)
	router.GET("/buildconfig/:namespace/:name", buildconfig.GetBCFromNS)
	router.GET("/buildconfig", buildconfig.GetAllBC)
	router.GET("/buildconfig/:namespace", buildconfig.GetAllBCFromNS)
	router.GET("/watch/buildconfig/:namespace/:name", buildconfig.WatchBCFromNS)
	router.GET("/watch/buildconfig", buildconfig.WatchAllBC)
	router.GET("/watch/buildconfig/:namespace", buildconfig.WatchAllBCFromNS)
	router.PUT("/buildconfig/:namesapce/:name", buildconfig.UpdataBCFromNS)
	router.PATCH("/buildconfig/:namesapce/:name", buildconfig.PatchBCFromNS)
	router.DELETE("/buildconfig/:namesapce/:name", buildconfig.DeleteBCFromNS)
	router.DELETE("/buildconfig/:namesapce", buildconfig.DeleteAllBuildFromNS)

	//v1.DeploymentConfig DC -> DeploymentConfig Dep -> Deployment
	router.POST("/deploymentconfig", dep.CreateDC)
	router.POST("/deploymentconfig/:namespace", dep.CreateDCInNS)
	router.POST("/ins/deploymentconfig/:namespace/:name", dep.CreateInsInNS)
	router.POST("/roolback/deploymentconfig/:namespace/:name", dep.CreateRollBackInNS)
	router.POST("/web/deploymentconfig/:namespace/:name", buildconfig.CreateWebInNS)
	router.POST("/web/deploymentconfig/:namespace/:name/:path", buildconfig.CreateWebInNSP)
	router.GET("/deploymentconfig/:namespace/:name", dep.GetDCFromNS)
	router.GET("/deploymentconfig", dep.GetAllDC)
	router.GET("/deploymentconfig/:namespace", dep.GetAllDCFromNS)
	router.GET("/watch/deploymentconfig/:namespace/:name", dep.WatchDCFromNS)
	router.GET("/watch/deploymentconfig", dep.WatchAllDC)
	router.GET("/watch/deploymentconfig/:namespace", dep.WatchAllDCFromNS)
	router.GET("/log/deploymentconfig/:namespace/:name", dep.GetLogDepFromNS)
	router.GET("/scale/deploymentconfig/:namespace/:name", dep.GetScaleDepFromNS)
	router.GET("/status/deploymentconfig/:namespace/:name", dep.GetStatusDepFromNS)
	router.PUT("/deploymentconfig/:namesapce/:name", dep.UpdataDCFromNS)
	router.PUT("/scale/deploymentconfig/:namesapce/:name", dep.UpdataScaleDCFromNS)
	router.PUT("/status/deploymentconfig/:namesapce/:name", dep.UpdataStatusDCFromNS)
	router.PATCH("/deploymentconfig/:namesapce/:name", dep.PatchDCFromNS)
	router.PATCH("/scale/deploymentconfig/:namesapce/:name", dep.PatchScaleDCFromNS)
	router.PATCH("/status/deploymentconfig/:namesapce/:name", dep.PatchStatusDCFromNS)
	router.DELETE("/deploymentconfig/:namesapce/:name", dep.DeleteDCFromNS)
	router.DELETE("/deploymentconfig/:namesapce", dep.DeleteAllDepFromNS)

	//v1.ImageStream IS -> ImageStream NS -> NameSpace SecImage ->SecretsImage
	router.POST("/imagestream", image.CreateIS)
	router.POST("/imagestream/:namespace", image.CreateImageInNS)
	router.GET("/imagestream/:namespace/:name", image.GetImageFromNS)
	router.GET("/imagestream", image.GetAllImage)
	router.GET("/imagestream/:namespace", image.GetAllImageFromNS)
	router.GET("/secrets/imagestream/:namespace/:name", image.GetSecImageFromNS)
	router.GET("/status/imagestream/:namespace/:name", image.GetStaImageFromNS)
	router.GET("/watch/imagestream/:namespace/:name", image.WatchImageFromNS)
	router.GET("/watch/imagestream", image.WatchAllImage)
	router.GET("/watch/imagestream/:namespace", image.WatchAllImageFromNS)
	router.PUT("/imagestream/:namespace/:name", image.UpdataImageFromNS)
	router.PUT("/status/imagestream/:namespace/:name", image.UpdataStaImageFromNS)
	router.PATCH("/imagestream/:namespace/:name", image.PatchImageFromNS)
	router.PATCH("/status/imagestream/:namespace/:name", image.PatchStaImageFromNS)
	router.DELETE("/imagestream/:namespace/:name", image.DeleteImageFromNS)
	router.DELETE("/imagestream/:namespace", image.DeleteAllImageFromNS)

	//v1.ImageStreamImport IS -> ImageStream NS -> NameSpace
	router.POST("/imagestreamimport", imagestreamimport.CreateISImport)
	router.POST("/imagestreamimport/:namespace", imagestreamimport.CreateISImportInNS)

	//v1.ImageStreamTag IST -> ImageStreamTag NS -> NameSpace
	router.POST("/imagestreamtag", tag.CreateIST)
	router.POST("/imagestreamtag/:namespace", tag.CreateImageTagInNS)
	router.GET("/imagestreamtag/:namespace/:name", tag.GetImageTagFromNS)
	router.GET("/imagestreamtag", tag.GetAllImageTag)
	router.GET("/imagestreamtag/:namespace", tag.GetAllImageTagFromNS)
	router.PUT("/imagestreamtag/:namespace/:name", tag.UpdataImageTagFromNS)
	router.PATCH("/imagestreamtag/:namespace/:name", tag.PatchImageTagFromNS)
	router.DELETE("/imagestreamtag/:namespace/:name", tag.DeleteImageTagFromNS)

	//v1.NetNamespace
	router.POST("/netnamespaces", netnamespace.CreateNetNamespace)
	router.GET("/netnamespaces/:name", netnamespace.GetNetNamespace)
	router.GET("/netnamespaces", netnamespace.GetAllNetNamespaces)
	router.GET("/watch/netnamespaces/:name", netnamespace.WatchNetNamespace)
	router.GET("/watch/netnamespaces", netnamespace.WatchAllNetNamespaces)
	router.PUT("/netnamespaces/:name", netnamespace.UpdateNetNamespace)
	router.PATCH("/netnamespaces/:name", netnamespace.PatchNetNamespace)
	router.DELETE("/netnamespaces/:name", netnamespace.DeleteNetNamespace)
	router.DELETE("/netnamespaces", netnamespace.DeleteAllNetNamespaces)

	//v1.Role
	router.POST("/roles", role.CreateRole)
	router.POST("/roles/:namespace", role.CreateRoleInNS)
	router.GET("/roles/:namespace/:name", role.GetRoleInNS)
	router.GET("/roles", role.GetAllRoles)
	router.GET("/roles/:namespace", role.GetRolesInNS)
	router.PUT("/roles/:namespace/:name", role.UpdateRoleInNS)
	router.PATCH("/roles/:namespace/:name", role.PatchRoleInNS)
	router.DELETE("/roles/:namespace/:name", role.DeleteRoleInNS)

	//v1.RoleBinding
	router.POST("/rolebindings", rolebinding.CreateRoleBinding)
	router.POST("/rolebindings/:namespace", rolebinding.CreateRoleBindingInNS)
	router.GET("/rolebindings/:namespace/:name", rolebinding.GetRoleBindingInNS)
	router.GET("/rolebindings", rolebinding.GetAllRoleBindings)
	router.GET("/rolebindings/:namespace", rolebinding.GetRoleBindingsInNS)
	router.PUT("/rolebindings/:namespace/:name", rolebinding.UpdateRoleBindingInNS)
	router.PATCH("/rolebindings/:namespace/:name", rolebinding.PatchRoleBindingInNS)
	router.DELETE("/rolebindings/:namespace/:name", rolebinding.DeleteRoleBindingInNS)

	//v1.Route
	router.POST("/routes", route.CreateRoute)
	router.POST("/routes/:namespace", route.CreateRouteInNS)
	router.GET("/routes/:namespace/:name", route.GetRouteInNS)
	router.GET("/routes", route.GetAllRoutes)
	router.GET("/routes/:namespace", route.GetAllRoutesInNS)
	router.GET("/watch/routes/:namespace/:name", route.WatchRouteInNS)
	router.GET("/watch/routes", route.WatchAllRoutes)
	router.GET("/watch/routes/:namespace", route.WatchAllRoutesInNS)
	router.GET("/status/routes/:namespace/:name", route.GetRouteStatusInNS)
	router.PUT("/routes/:namespace/:name", route.UpdateRouteInNS)
	router.PUT("/status/routes/:namespace/:name", route.UpdateRouteStatusInNS)
	router.PATCH("/routes/:namespace/:name", route.PatchRouteInNS)
	router.PATCH("/status/routes/:namespace/:name", route.PatchRouteStatusInNS)
	router.DELETE("/routes/:namespace/:name", route.DeleteRouteInNS)
	router.DELETE("/routes/:namespace", route.DeleteAllRoutesInNS)

	//v1.Template
	router.POST("/templates", template.CreateTemplate)
	router.POST("/templates/:namespace", template.CreateTemplatenNS)
	router.GET("/templates/:namespace/:name", template.GetTemplateInNS)
	router.GET("/templates", template.GetAllTemplates)
	router.GET("/templates/:namespace", template.GetAllTemplatesInNS)
	router.GET("/watch/templates/:namespace/:name", template.WatchTemplateInNS)
	router.GET("/watch/templates", template.WatchAllTemplates)
	router.GET("/watch/templates/:namespace", template.WatchAllTemplatesInNS)
	router.PUT("/templates/:namespace/:name", template.UpdateTemplateInNS)
	router.PATCH("/templates/:namespace/:name", template.PatchTemplateInNS)
	router.DELETE("/templates/:namespace/:name", template.DeleteTemplateInNS)
	router.DELETE("/templates/:namespace", template.DeleteAllTemplatesInNS)

	return
}

/*
var origin = "http://new.dataos.io:8443/"
var url = "ws://new.dataos.io:8443?token=zTUpzpPPTaKtY0ZCTB80FP0djjwQ1e36TGT17a3OR1M/oapi/v1/watch/projects/wutest001"

func main() {

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println("-----------------------")
		log.Fatal(err)
	}

	var msg = make([]byte, 512)
	m, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg[:m])

	ws.Close()//关闭连接
}
*/