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
	"github.com/asiainfoLDP/datafoundry-gw/k8sapi/service"
	"github.com/asiainfoLDP/datafoundry-gw/k8sapi/secret"
	rq "github.com/asiainfoLDP/datafoundry-gw/k8sapi/resourcequota"
	rc "github.com/asiainfoLDP/datafoundry-gw/k8sapi/replicationcontroller"
	"github.com/asiainfoLDP/datafoundry-gw/k8sapi/pod"
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
	//router.GET("/oapi/v1/watch/namespaces/:namespace/builds/:name", build.watchBuildFromNS)
	//router.GET("/oapi/v1/watch/builds", build.watchAllBuilds)
	//router.GET("/oapi/v1/watch/namespaces/:namespace/builds", build.watchAllBuildFromNS)
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
	//router.GET("/oapi/v1/watch/namespaces/:namespace/buildconfigs/:name", buildconfig.watchBCFromNS)
	//router.GET("/oapi/v1/watch/buildconfigs", buildconfig.watchAllBC)
	//router.GET("/oapi/v1/watch/namespaces/:namespace/buildconfigs", buildconfig.watchAllBCFromNS)
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
	//router.GET("/oapi/v1/watch/namespaces/:namespace/deploymentconfigs/:name", dep.watchDCFromNS)
	//router.GET("/oapi/v1/watch/deploymentconfigs", dep.watchAllDC)
	//router.GET("/oapi/v1/watch/namespaces/:namespace/deploymentconfigs", dep.watchAllDCFromNS)
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
	//router.GET("/oapi/v1/watch/namespaces/:namespace/imagestreams/:name", image.watchImageFromNS)
	//router.GET("/oapi/v1/watch/imagestreams", image.watchAllImage)
	//router.GET("/oapi/v1/watch/namespaces/:namespace/imagestreams", image.watchAllImageFromNS)
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

	//k8s api
	//v1.Pod
	router.POST("/api/v1/pods",pod.CreatePod)
	router.POST("/api/v1/namespaces/:namespace/pods", pod.CreatePodInNS)
	router.POST("/api/v1/namespaces/:namespace/pods/:name/attach", pod.AttachPodInNS)
	router.POST("/api/v1/namespaces/:namespace/pods/:name/binding", pod.CreateBindPodInNS)
	router.POST("/api/v1/namespaces/:namespace/pods/:name/eviction", pod.CreateEvtPodInNS)
	router.POST("/api/v1/namespaces/:namespace/pods/:name/exec", pod.CreateExecPodInNS)
	router.POST("/api/v1/namespaces/:namespace/pods/:name/portforward", pod.PortPodInNS)
	router.POST("/api/v1/namespaces/:namespace/pods/:name/proxy", pod.ProxyPodInNS)
	router.POST("/api/v1/namespaces/:namespace/pods/:name/proxy/:path", pod.ProxysPathInNS)
	router.HEAD("/api/v1/namespaces/:namespace/pods/:name/proxy", pod.HeadPodInNS)
	router.HEAD("/api/v1/namespaces/:namespace/pods/:name/proxy/:path", pod.HeadProxysPathInNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name",pod.GetPodFromNS)
	router.GET("/api/v1/services", pod.GetAllPod)
	router.GET("/api/v1/namespaces/:namespace/pods", pod.GetAllPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/attach",pod.GetAtaPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/exec",pod.GetExecPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/log",pod.GetLogPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/portforward",pod.GetPortPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/status",pod.GetStatusPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/proxy",pod.GetProxyPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/proxy/:path",pod.GetProxyPathPodFromNS)
	router.PUT("/api/v1/namespaces/:namespace/pods/:name", pod.UpdataPodFromNS)
	router.PUT("/api/v1/namespaces/:namespace/pods/:name/status", pod.UpdataStuPodFromNS)
	router.PUT("/api/v1/namespaces/:namespace/pods/:name/proxy", pod.UpdataProxyPodFromNS)
	router.PUT("/api/v1/namespaces/:namespace/pods/:name/proxy/:path", pod.UpdataProPathPodFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/pods/:name", pod.PatchPodFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/pods/:name/status", pod.PatchStuPodFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/pods/:name/proxy", pod.PatchProxyPodFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/pods/:name/proxy/:path", pod.PatchProPathPodFromNS)
	router.OPTIONS("/api/v1/namespaces/:namespace/pods/:name/proxy", pod.OptionsPodFromNS)
	router.OPTIONS("/api/v1/namespaces/:namespace/pods/:name/proxy/:path", pod.OptionsPathPodFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/pods/:name", pod.DeletePodFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/pods/:name/proxy", pod.DeleteProxyPodFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/pods/:name/proxy/:path", pod.DeleteProxyPathPodFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/pods", pod.DeleteAllPodFromNS)

	//v1.ReplicationController
	router.POST("/api/v1/replicationcontrollers",rc.CreateRc)
	router.POST("/api/v1/namespaces/:namespace/replicationcontrollers", rc.CreateRcInNS)
	router.GET("/api/v1/namespaces/:namespace/replicationcontrollers/:name",rc.GetRcFromNS)
	router.GET("/api/v1/replicationcontrollers", rc.GetAllRc)
	router.GET("/api/v1/namespaces/:namespace/replicationcontrollers", rc.GetAllRcFromNS)
	router.GET("/api/v1/namespaces/:namespace/replicationcontrollers/:name/scale",rc.GetScaleRcFromNS)
	router.GET("/apis/extensions/v1beta1/namespaces/:namespace/replicationcontrollers/:name/scale",rc.GetExScaleRcFromNS)
	router.GET("/api/v1/namespaces/:namespace/replicationcontrollers/:name/status",rc.GetStatusRcFromNS)
	router.PUT("/api/v1/namespaces/:namespace/replicationcontrollers/:name", rc.UpdataRcFromNS)
	router.PUT("/api/v1/namespaces/:namespace/replicationcontrollers/:name/scale", rc.UpdataScaleRcFromNS)
	router.PUT("/apis/extensions/v1beta1/namespaces/:namespace/replicationcontrollers/:name/scale", rc.UpdataExScaleRcFromNS)
	router.PUT("/api/v1/namespaces/:namespace/replicationcontrollers/:name/proxy/:path", rc.UpdataStatusRcFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/replicationcontrollers/:name", rc.PatchRcFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/replicationcontrollers/:name/sacle", rc.PatchScaleRcFromNS)
	router.PATCH("/apis/extensions/v1beta1/namespaces/:namespace/replicationcontrollers/:name/scale", rc.PatchExScaleFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/replicationcontrollers/:name/proxy/:path", rc.PatchStatusRcFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/replicationcontrollers/:name", rc.DeleteRcFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/replicationcontrollers", rc.DeleteAllRcFromNS)

	//v1.ResourceQuota
	router.POST("/api/v1/resourcequotas",rq.CreateRq)
	router.POST("/api/v1/namespaces/:namespace/resourcequotas", rq.CreateRqInNS)
	router.GET("/api/v1/namespaces/:namespace/resourcequotas/:name",rq.GetRqFromNS)
	router.GET("/api/v1/resourcequotas", rq.GetAllRq)
	router.GET("/api/v1/namespaces/:namespace/resourcequotas", rq.GetAllRqFromNS)
	router.GET("/api/v1/namespaces/:namespace/resourcequotas/:name/status",rq.GetStuRqFromNS)
	router.PUT("/api/v1/namespaces/:namespace/resourcequotas/:name", rq.UpdataRqFromNS)
	router.PUT("/api/v1/namespaces/:namespace/resourcequotas/:name/status",rq.UpdataStuRqFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/resourcequotas/:name", rq.PatchRqFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/resourcequotas/:name/status",rq.PatchStuRqFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/resourcequotas/:name", rq.DeleteRqFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/resourcequotas", rq.DeleteAllRqFromNS)

	//v1.Secret
	router.POST("/api/v1/secrets",secret.CreateSecret)
	router.POST("/api/v1/namespaces/:namespace/secrets", secret.CreateSecretInNS)
	router.GET("/api/v1/namespaces/:namespace/secrets/:name",secret.GetSecretFromNS)
	router.GET("/api/v1/secrets", secret.GetAllSecret)
	router.GET("/api/v1/namespaces/:namespace/secrets", secret.GetAllSecretFromNS)
	router.PUT("/api/v1/namespaces/:namespace/secrets/:name", secret.UpdataSecretFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/secrets/:name", secret.PatchSecretFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/secrets/:name", secret.DeleteSecretFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/secrets", secret.DeleteAllSecretFromNS)

	//v1.Service
	router.POST("/api/v1/services",service.CreateService)
	router.POST("/api/v1/namespaces/:namespace/services", service.CreateServiceInNS)
	router.POST("/api/v1/namespaces/:namespace/services/:name/proxy", service.CreateProxysInNS)
	router.POST("/api/v1/namespaces/:namespace/services/:name/proxy/:path", service.CreateProxysPathInNS)
	router.HEAD("/api/v1/namespaces/:namespace/services/:name/proxy", service.HeadProxysInNS)
	router.HEAD("/api/v1/namespaces/:namespace/services/:name/proxy/:path", service.HeadProxysPathInNS)
	router.GET("/api/v1/namespaces/:namespace/services/:name",service.GetServiceFromNS)
	router.GET("/api/v1/service", service.GetAllServices)
	router.GET("/api/v1/namespaces/:namespace/services", service.GetAllServicesFromNS)
	router.GET("/api/v1/namespaces/:namespace/services/:name/status",service.GetStuServiceFromNS)
	router.GET("/api/v1/namespaces/:namespace/services/:name/proxy",service.GetProServiceFromNS)
	router.GET("/api/v1/namespaces/:namespace/services/:name/proxy/:path",service.GetProPathServiceFromNS)
	router.PUT("/api/v1/namespaces/:namespace/services/:name", service.UpdataServicesFromNS)
	router.PUT("/api/v1/namespaces/:namespace/services/:name/status", service.UpdataStuServicesFromNS)
	router.PUT("/api/v1/namespaces/:namespace/services/:name/proxy", service.UpdataProServicesFromNS)
	router.PUT("/api/v1/namespaces/:namespace/services/:name/proxy/:path", service.UpdataProPathServicesFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/services/:name", service.PatchServicesFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/services/:name/status", service.PatchStuServicesFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/services/:name/proxy", service.PatchProServicesFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/services/:name/proxy/:path", service.PatchProPathServicesFromNS)
	router.OPTIONS("/api/v1/namespaces/:namespace/services/:name/proxy", service.OptionsServicesFromNS)
	router.OPTIONS("/api/v1/namespaces/:namespace/services/:name/proxy/:path", service.OptionsPathServicesFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/services/:name/proxy", service.DeleteProServicesFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/services/:name/proxy/:path", service.DeleteProPathServicesFromNS)


	return
}