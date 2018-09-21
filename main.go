package main

import (

	//_ "net/http/pprof"
	//"github.com/DeanThompson/ginpprof"
	"github.com/asiainfoldp/datafoundry-gw/apis"
	"github.com/asiainfoldp/datafoundry-gw/hawkular"
	"github.com/asiainfoldp/datafoundry-gw/k8sapi/configmap"
	"github.com/asiainfoldp/datafoundry-gw/k8sapi/endpoints"
	"github.com/asiainfoldp/datafoundry-gw/k8sapi/event"
	"github.com/asiainfoldp/datafoundry-gw/k8sapi/limitrange"
	"github.com/asiainfoldp/datafoundry-gw/k8sapi/namespace"
	"github.com/asiainfoldp/datafoundry-gw/k8sapi/node"
	"github.com/asiainfoldp/datafoundry-gw/k8sapi/persistentvolume"
	"github.com/asiainfoldp/datafoundry-gw/k8sapi/persistentvolumeclaim"
	"github.com/asiainfoldp/datafoundry-gw/k8sapi/pod"
	rc "github.com/asiainfoldp/datafoundry-gw/k8sapi/replicationcontroller"
	rq "github.com/asiainfoldp/datafoundry-gw/k8sapi/resourcequota"
	"github.com/asiainfoldp/datafoundry-gw/k8sapi/secret"
	"github.com/asiainfoldp/datafoundry-gw/k8sapi/service"
	"github.com/asiainfoldp/datafoundry-gw/lapi"
	"github.com/asiainfoldp/datafoundry-gw/oapi/build"
	"github.com/asiainfoldp/datafoundry-gw/oapi/buildconfig"
	dep "github.com/asiainfoldp/datafoundry-gw/oapi/deploymentconfig"
	image "github.com/asiainfoldp/datafoundry-gw/oapi/imagestream"
	"github.com/asiainfoldp/datafoundry-gw/oapi/imagestreamimport"
	tag "github.com/asiainfoldp/datafoundry-gw/oapi/imagestreamtag"
	"github.com/asiainfoldp/datafoundry-gw/oapi/netnamespace"
	"github.com/asiainfoldp/datafoundry-gw/oapi/processedtemplate"
	"github.com/asiainfoldp/datafoundry-gw/oapi/project"
	"github.com/asiainfoldp/datafoundry-gw/oapi/role"
	"github.com/asiainfoldp/datafoundry-gw/oapi/rolebinding"
	"github.com/asiainfoldp/datafoundry-gw/oapi/route"
	"github.com/asiainfoldp/datafoundry-gw/oapi/template"
	"github.com/asiainfoldp/datafoundry-gw/oapi/user"
	"github.com/asiainfoldp/datafoundry-gw/others"
	"github.com/asiainfoldp/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"net/http"
	"os"
	"time"
)

//定义日志以及其他变量
var logger lager.Logger
var ocAPIName string = "ocapi"

func main() {
	//初始化日志对象，日志输出到stdout
	logger = lager.NewLogger(ocAPIName)
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.INFO)) //默认日志级别
	router := handle()
	//运行测试调优代码需打开依赖包及下面一条语句，并设置WriteTimeout需大约测试设置的时间或直接注释掉
	//ginpprof.Wrap(router)
	s := &http.Server{
		Addr:           ":10012",
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 0,
	}
	logger.Info("Service starting ...", map[string]interface{}{"user": "master", "time": pkg.GetTimeNow()})
	//监听端口
	s.ListenAndServe()
}

func handle() (router *gin.Engine) {
	//设置全局环境：1.开发环境（gin.DebugMode） 2.线上环境（gin.ReleaseMode）
	gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)
	//获取路由实例
	router = gin.Default()

	//hawkular
	router.POST("/hawkular/cpu", hawkular.GainCpu)
	router.POST("/hawkular/memory", hawkular.GainMemory)
	router.POST("/hawkular/network/:sigin", hawkular.GainNetwork)

	//OpenShift API
	//v1.user
	router.Group("/oapi/v1/users/~", user.GetSelf)
	router.GET("/oapi/v1/users/:name", user.GetUser)
	router.GET("/oapi/v1/users/", user.GetAllUser)
	router.GET("/oapi/v1/users", user.GetAllUser)
	router.DELETE("/oapi/v1/users/:name", user.DeleteUser)

	//v1.project
	router.POST("/oapi/v1/projectrequests", project.CreateProject)
	router.GET("/oapi/v1/projects/:name", project.GorWProject)
	router.GET("/oapi/v1/projects", project.GorWAllProjects)
	router.GET("/oapi/v1/projects/", project.GorWAllProjects)
	router.PUT("/oapi/v1/projects/:name", project.UpdateProject)
	router.PATCH("/oapi/v1/projects/:name", project.PatchAProject)
	router.DELETE("/oapi/v1/projects/:name", project.DeleteProject)

	//v1.Build NS -> NameSpace
	router.POST("/oapi/v1/builds", build.CreateBuild)
	router.POST("/oapi/v1/namespaces/:namespace/builds", build.CreateBuildInNS)
	router.POST("/oapi/v1/namespaces/:namespace/builds/:name/clone", build.CreateCloneInNS)
	router.GET("/oapi/v1/namespaces/:namespace/builds/:name", build.GetBuildFromNS)
	router.GET("/oapi/v1/builds", build.GetAllBuilds)
	router.GET("/oapi/v1/namespaces/:namespace/builds", build.GetAllBuildFromNS)
	router.GET("/oapi/v1/namespaces/:namespace/builds/:name/log", build.GetLogBuildFromNS)
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
	router.GET("/oapi/v1/namespaces/:namespace/buildconfigs/", buildconfig.GetAllBCFromNS)
	router.PUT("/oapi/v1/namespaces/:namespace/buildconfigs/:name", buildconfig.UpdataBCFromNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/buildconfigs/:name", buildconfig.PatchBCFromNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/buildconfigs/:name", buildconfig.DeleteBCFromNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/buildconfigs", buildconfig.DeleteAllBuildConFromNS)

	//v1.DeploymentConfig DC -> DeploymentConfig Dep -> Deployment
	router.POST("/oapi/v1/deploymentconfigs", dep.CreateDC)
	router.POST("/oapi/v1/namespaces/:namespace/deploymentconfigs", dep.CreateDCInNS)
	router.POST("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name/instantiate", dep.CreateInsInNS)
	router.POST("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name/rollback", dep.CreateRollBackInNS)
	router.GET("/oapi/v1/namespaces/:namespace/deploymentconfigs/:name", dep.GetDCFromNS)
	router.GET("/oapi/v1/deploymentconfigs", dep.GetAllDC)
	router.GET("/oapi/v1/namespaces/:namespace/deploymentconfigs", dep.GetAllDCFromNS)
	router.GET("/oapi/v1/namespaces/:namespace/deploymentconfigs/", dep.GetAllDCFromNS)
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
	router.GET("/oapi/v1/namespaces/:namespace/imagestreams/", image.GetAllImageFromNS)
	router.GET("/oapi/v1/namespaces/:namespace/imagestreams/:name/secrets", image.GetSecImageFromNS)
	router.GET("/oapi/v1/namespaces/:namespace/imagestreams/:name/status", image.GetStaImageFromNS)
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
	router.GET("/oapi/v1/namespaces/:namespace/imagestreamtags/", tag.GetAllImageTagFromNS)
	router.PUT("/oapi/v1/namespaces/:namespace/imagestreamtags/:name", tag.UpdataImageTagFromNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/imagestreamtags/:name", tag.PatchImageTagFromNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/imagestreamtags/:name", tag.DeleteImageTagFromNS)

	//v1.NetNamespace
	router.POST("/oapi/v1/netnamespaces", netnamespace.CreateNetNamespace)
	router.GET("/oapi/v1/netnamespaces/:name", netnamespace.GorWNetNamespace)
	router.GET("/oapi/v1/netnamespaces", netnamespace.GorWAllNetNamespaces)
	router.GET("/oapi/v1/netnamespaces/", netnamespace.GorWAllNetNamespaces)
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
	router.GET("/oapi/v1/namespaces/:namespace/roles/", role.GetRolesInNS)
	router.PUT("/oapi/v1/namespaces/:namespace/roles/:name", role.UpdateRoleInNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/roles/:name", role.PatchRoleInNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/roles/:name", role.DeleteRoleInNS)

	//v1.RoleBinding
	router.POST("/oapi/v1/rolebindings", rolebinding.CreateRoleBinding)
	router.POST("/oapi/v1/namespaces/:namespace/rolebindings", rolebinding.CreateRoleBindingInNS)
	router.GET("/oapi/v1/namespaces/:namespace/rolebindings/:name", rolebinding.GetRoleBindingInNS)
	router.GET("/oapi/v1/rolebindings", rolebinding.GetAllRoleBindings)
	router.GET("/oapi/v1/namespaces/:namespace/rolebindings", rolebinding.GetRoleBindingsInNS)
	router.GET("/oapi/v1/namespaces/:namespace/rolebindings/", rolebinding.GetRoleBindingsInNS)
	router.PUT("/oapi/v1/namespaces/:namespace/rolebindings/:name", rolebinding.UpdateRoleBindingInNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/rolebindings/:name", rolebinding.PatchRoleBindingInNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/rolebindings/:name", rolebinding.DeleteRoleBindingInNS)

	//v1.Route
	router.POST("/oapi/v1/routes", route.CreateRoute)
	router.POST("/oapi/v1/namespaces/:namespace/routes", route.CreateRouteInNS)
	router.GET("/oapi/v1/namespaces/:namespace/routes/:name", route.GorWRouteInNS)
	router.GET("/oapi/v1/routes", route.GorWAllRoutes)
	router.GET("/oapi/v1/namespaces/:namespace/routes", route.GorWAllRoutesInNS)
	router.GET("/oapi/v1/namespaces/:namespace/routes/", route.GorWAllRoutesInNS)
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
	router.GET("/oapi/v1/namespaces/:namespace/templates/", template.GorWAllTemplatesInNS)
	router.PUT("/oapi/v1/namespaces/:namespace/templates/:name", template.UpdateTemplateInNS)
	router.PATCH("/oapi/v1/namespaces/:namespace/templates/:name", template.PatchTemplateInNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/templates/:name", template.DeleteTemplateInNS)
	router.DELETE("/oapi/v1/namespaces/:namespace/templates", template.DeleteAllTemplatesInNS)

	//v1.ProcessedTemplate
	//自定义接口
	router.PUT("/oapi/v1/namespaces/:namespace/processedtemplates", processedtemplate.CDProcessedTemplate)

	//K8S API
	//v1.Pod
	router.POST("/api/v1/pods", pod.CreatePod)
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
	router.GET("/api/v1/namespaces/:namespace/pods/:name", pod.GetPodFromNS)
	router.GET("/api/v1/pods", pod.GetAllPod)
	router.GET("/api/v1/namespaces/:namespace/pods", pod.GetAllPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/", pod.GetAllPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/attach", pod.GetAtaPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/exec", pod.GetExecPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/log", pod.GetLogPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/portforward", pod.GetPortPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/status", pod.GetStatusPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/proxy", pod.GetProxyPodFromNS)
	router.GET("/api/v1/namespaces/:namespace/pods/:name/proxy/:path", pod.GetProxyPathPodFromNS)
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
	router.POST("/api/v1/replicationcontrollers", rc.CreateRc)
	router.POST("/api/v1/namespaces/:namespace/replicationcontrollers", rc.CreateRcInNS)
	router.GET("/api/v1/namespaces/:namespace/replicationcontrollers/:name", rc.GetRcFromNS)
	router.GET("/api/v1/replicationcontrollers", rc.GetAllRc)
	router.GET("/api/v1/namespaces/:namespace/replicationcontrollers", rc.GetAllRcFromNS)
	router.GET("/api/v1/namespaces/:namespace/replicationcontrollers/", rc.GetAllRcFromNS)
	router.GET("/api/v1/namespaces/:namespace/replicationcontrollers/:name/scale", rc.GetScaleRcFromNS)
	router.GET("/apis/extensions/v1beta1/namespaces/:namespace/replicationcontrollers/:name/scale", rc.GetExScaleRcFromNS)
	router.GET("/api/v1/namespaces/:namespace/replicationcontrollers/:name/status", rc.GetStatusRcFromNS)
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
	router.POST("/api/v1/resourcequotas", rq.CreateRq)
	router.POST("/api/v1/namespaces/:namespace/resourcequotas", rq.CreateRqInNS)
	router.GET("/api/v1/namespaces/:namespace/resourcequotas/:name", rq.GetRqFromNS)
	router.GET("/api/v1/resourcequotas", rq.GetAllRq)
	router.GET("/api/v1/namespaces/:namespace/resourcequotas", rq.GetAllRqFromNS)
	router.GET("/api/v1/namespaces/:namespace/resourcequotas/", rq.GetAllRqFromNS)
	router.GET("/api/v1/namespaces/:namespace/resourcequotas/:name/status", rq.GetStuRqFromNS)
	router.PUT("/api/v1/namespaces/:namespace/resourcequotas/:name", rq.UpdataRqFromNS)
	router.PUT("/api/v1/namespaces/:namespace/resourcequotas/:name/status", rq.UpdataStuRqFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/resourcequotas/:name", rq.PatchRqFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/resourcequotas/:name/status", rq.PatchStuRqFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/resourcequotas/:name", rq.DeleteRqFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/resourcequotas", rq.DeleteAllRqFromNS)

	//v1.Secret
	router.POST("/api/v1/secrets", secret.CreateSecret)
	router.POST("/api/v1/namespaces/:namespace/secrets", secret.CreateSecretInNS)
	router.GET("/api/v1/namespaces/:namespace/secrets/:name", secret.GetSecretFromNS)
	router.GET("/api/v1/secrets", secret.GetAllSecret)
	router.GET("/api/v1/namespaces/:namespace/secrets", secret.GetAllSecretFromNS)
	router.GET("/api/v1/namespaces/:namespace/secrets/", secret.GetAllSecretFromNS)
	router.PUT("/api/v1/namespaces/:namespace/secrets/:name", secret.UpdataSecretFromNS)
	router.PATCH("/api/v1/namespaces/:namespace/secrets/:name", secret.PatchSecretFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/secrets/:name", secret.DeleteSecretFromNS)
	router.DELETE("/api/v1/namespaces/:namespace/secrets", secret.DeleteAllSecretFromNS)

	//v1.Service
	router.POST("/api/v1/services", service.CreateService)
	router.POST("/api/v1/namespaces/:namespace/services", service.CreateServiceInNS)
	router.POST("/api/v1/namespaces/:namespace/services/:name/proxy", service.CreateProxysInNS)
	router.POST("/api/v1/namespaces/:namespace/services/:name/proxy/:path", service.CreateProxysPathInNS)
	router.HEAD("/api/v1/namespaces/:namespace/services/:name/proxy", service.HeadProxysInNS)
	router.HEAD("/api/v1/namespaces/:namespace/services/:name/proxy/:path", service.HeadProxysPathInNS)
	router.GET("/api/v1/namespaces/:namespace/services/:name", service.GetServiceFromNS)
	router.GET("/api/v1/services", service.GetAllServices)
	router.GET("/api/v1/namespaces/:namespace/services", service.GetAllServicesFromNS)
	router.GET("/api/v1/namespaces/:namespace/services/", service.GetAllServicesFromNS)
	router.GET("/api/v1/namespaces/:namespace/services/:name/status", service.GetStuServiceFromNS)
	router.GET("/api/v1/namespaces/:namespace/services/:name/proxy", service.GetProServiceFromNS)
	router.GET("/api/v1/namespaces/:namespace/services/:name/proxy/:path", service.GetProPathServiceFromNS)
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

	//v1.ConfigMap
	router.POST("/api/v1/configmaps", configmap.CreateConfigMap)
	router.POST("/api/v1/namespaces/:namespace/configmaps", configmap.CreateConfigMapNS)
	router.GET("/api/v1/namespaces/:namespace/configmaps/:name", configmap.GorWConfigMapNS)
	router.GET("/api/v1/configmaps", configmap.GorWAllConfigMap)
	router.GET("/api/v1/namespaces/:namespace/configmaps", configmap.GorWAllConfigMapNS)
	router.GET("/api/v1/namespaces/:namespace/configmaps/", configmap.GorWAllConfigMapNS)
	router.PUT("/api/v1/namespaces/:namespace/configmaps/:name", configmap.UpdateConfigMapNS)
	router.PATCH("/api/v1/namespaces/:namespace/configmaps/:name", configmap.PatchConfigMapNS)
	router.DELETE("/api/v1/namespaces/:namespace/configmaps/:name", configmap.DeleteConfigMapNS)
	router.DELETE("/api/v1/namespaces/:namespace/configmaps", configmap.DeleteAllConfigMapNS)

	//v1.Endpoints
	router.POST("/api/v1/endpoints", endpoints.CreateEndpoints)
	router.POST("/api/v1/namespaces/:namespace/endpoints", endpoints.CreateEndpointsNS)
	router.GET("/api/v1/namespaces/:namespace/endpoints/:name", endpoints.GorWEndpointsNS)
	router.GET("/api/v1/endpoints", endpoints.GorWAllEndpoints)
	router.GET("/api/v1/namespaces/:namespace/endpoints", endpoints.GorWAllEndpointsNS)
	router.GET("/api/v1/namespaces/:namespace/endpoints/", endpoints.GorWAllEndpointsNS)
	router.PUT("/api/v1/namespaces/:namespace/endpoints/:name", endpoints.UpdateEndpointsNS)
	router.PATCH("/api/v1/namespaces/:namespace/endpoints/:name", endpoints.PatchEndpointsNS)
	router.DELETE("/api/v1/namespaces/:namespace/endpoints/:name", endpoints.DeleteEndpointsNS)
	router.DELETE("/api/v1/namespaces/:namespace/endpoints", endpoints.DeleteAllEndpointsNS)

	//v1.Event
	router.POST("/api/v1/events", event.CreateEvent)
	router.POST("/api/v1/namespaces/:namespace/events", event.CreateEventNS)
	router.GET("/api/v1/namespaces/:namespace/events/:name", event.GorWEventNS)
	router.GET("/api/v1/events", event.GorWAllEvents)
	router.GET("/api/v1/namespaces/:namespace/events", event.GorWAllEventsNS)
	router.GET("/api/v1/namespaces/:namespace/events/", event.GorWAllEventsNS)
	router.PUT("/api/v1/namespaces/:namespace/events/:name", event.UpdateEventNS)
	router.PATCH("/api/v1/namespaces/:namespace/events/:name", event.PatchEventNS)
	router.DELETE("/api/v1/namespaces/:namespace/events/:name", event.DeleteEventNS)
	router.DELETE("/api/v1/namespaces/:namespace/events", event.DeleteAllEventNS)

	//v1.LimitRange
	router.POST("/api/v1/limitranges", limitrange.CreateLimitRange)
	router.POST("/api/v1/namespaces/:namespace/limitranges", limitrange.CreateLimitRangeNS)
	router.GET("/api/v1/namespaces/:namespace/limitranges/:name", limitrange.GorWLimitRangeNS)
	router.GET("/api/v1/limitranges", limitrange.GorWAllLimitRanges)
	router.GET("/api/v1/namespaces/:namespace/limitranges", limitrange.GorWAllLimitRangesNS)
	router.GET("/api/v1/namespaces/:namespace/limitranges/", limitrange.GorWAllLimitRangesNS)
	router.PUT("/api/v1/namespaces/:namespace/limitranges/:name", limitrange.UpdateLimitRangeNS)
	router.PATCH("/api/v1/namespaces/:namespace/limitranges/:name", limitrange.PatchLimitRangeNS)
	router.DELETE("/api/v1/namespaces/:namespace/limitranges/:name", limitrange.DeleteLimitRangeNS)
	router.DELETE("/api/v1/namespaces/:namespace/limitranges", limitrange.DeleteAllLimitRangeNS)

	//v1.Namespace
	router.POST("/api/v1/namespaces", namespace.CreateNamespace)
	router.GET("/api/v1/namespaces/:namespace", namespace.GorWNamespace)
	router.GET("/api/v1/namespaces", namespace.GorWAllNamespaces)
	router.GET("/api/v1/namespaces/", namespace.GorWAllNamespaces)
	router.PUT("/api/v1/namespaces/:namespace", namespace.UpdateNamespace)
	router.PATCH("/api/v1/namespaces/:namespace", namespace.PatchNamespace)
	router.DELETE("/api/v1/namespaces/:namespace", namespace.DeleteNamespace)
	router.PUT("/api/v1/namespaces/:namespace/finalize", namespace.UpdatefinalizeofNS)
	router.GET("/api/v1/namespaces/:namespace/status", namespace.GetstatusofNS)
	router.PUT("/api/v1/namespaces/:namespace/status", namespace.UpdatestatusofNS)
	router.PATCH("/api/v1/namespaces/:namespace/status", namespace.PatchstatusofNS)

	//v1.Node
	router.POST("/api/v1/nodes", node.CreateNode)
	router.GET("/api/v1/nodes/:name", node.GorWNode)
	router.GET("/api/v1/nodes", node.GorWAllNodes)
	router.GET("/api/v1/nodes/", node.GorWAllNodes)
	router.PUT("/api/v1/nodes/:name", node.UpdateNode)
	router.PATCH("/api/v1/nodes/:name", node.PatchNode)
	router.DELETE("/api/v1/nodes/:name", node.DeleteNode)
	router.DELETE("/api/v1/nodes", node.DeleteAllNodes)
	router.GET("/api/v1/nodes/:name/status", node.GetStatusOfNode)
	router.PUT("/api/v1/nodes/:name/status", node.UpdateStatusOfNode)
	router.PATCH("/api/v1/nodes/:name/status", node.PatchStatusOfNode)
	router.OPTIONS("/api/v1/nodes/:name/proxy", node.ProxyOpnReqToNode)
	router.POST("/api/v1/nodes/:name/proxy", node.ProxyPostReqToNode)
	router.HEAD("/api/v1/nodes/:name/proxy", node.ProxyHeadReqToNode)
	router.GET("/api/v1/nodes/:name/proxy", node.ProxyGetReqToNode)
	router.PUT("/api/v1/nodes/:name/proxy", node.ProxyPutReqToNode)
	router.PATCH("/api/v1/nodes/:name/proxy", node.ProxyPatchReqToNode)
	router.DELETE("/api/v1/nodes/:name/proxy", node.ProxyDelReqToNode)
	router.OPTIONS("/api/v1/nodes/:name/proxy/:path", node.ProxyOpnReqToNodeP)
	router.POST("/api/v1/nodes/:name/proxy/:path", node.ProxyPostReqToNodeP)
	router.HEAD("/api/v1/nodes/:name/proxy/:path", node.ProxyHeadReqToNodeP)
	router.GET("/api/v1/nodes/:name/proxy/:path", node.ProxyGetReqToNodeP)
	router.PUT("/api/v1/nodes/:name/proxy/:path", node.ProxyPutReqToNodeP)
	router.PATCH("/api/v1/nodes/:name/proxy/:path", node.ProxyPatchReqToNodeP)
	router.DELETE("/api/v1/nodes/:name/proxy/:path", node.ProxyDelReqToNodeP)

	//v1.PersistentVolume
	router.POST("/api/v1/persistentvolumes", persistentvolume.CreatePV)
	router.GET("/api/v1/persistentvolumes/:name", persistentvolume.GorWPV)
	router.GET("/api/v1/persistentvolumes", persistentvolume.GorWAllPVs)
	router.GET("/api/v1/persistentvolumes/", persistentvolume.GorWAllPVs)
	router.PUT("/api/v1/persistentvolumes/:name", persistentvolume.UpdatePV)
	router.PATCH("/api/v1/persistentvolumes/:name", persistentvolume.PatchPV)
	router.DELETE("/api/v1/persistentvolumes/:name", persistentvolume.DeletePV)
	router.DELETE("/api/v1/persistentvolumes", persistentvolume.DeleteAllPVs)
	router.GET("/api/v1/persistentvolumes/:name/status", persistentvolume.GetstatusofPV)
	router.PUT("/api/v1/persistentvolumes/:name/status", persistentvolume.UpdatestatusofPV)
	router.PATCH("/api/v1/persistentvolumes/:name/status", persistentvolume.PatchstatusofPV)

	//v1.PersistentVolumeClaim
	router.POST("/api/v1/persistentvolumeclaims", persistentvolumeclaim.CreatePVC)
	router.POST("/api/v1/namespaces/:namespace/persistentvolumeclaims", persistentvolumeclaim.CreatePVCns)
	router.GET("/api/v1/namespaces/:namespace/persistentvolumeclaims/:name", persistentvolumeclaim.GorWPVCns)
	router.GET("/api/v1/persistentvolumeclaims", persistentvolumeclaim.GorWAllPVC)
	router.GET("/api/v1/namespaces/:namespace/persistentvolumeclaims", persistentvolumeclaim.GorWAllPVCns)
	router.GET("/api/v1/namespaces/:namespace/persistentvolumeclaims/", persistentvolumeclaim.GorWAllPVCns)
	router.PUT("/api/v1/namespaces/:namespace/persistentvolumeclaims/:name", persistentvolumeclaim.UpdatePVCns)
	router.PATCH("/api/v1/namespaces/:namespace/persistentvolumeclaims/:name", persistentvolumeclaim.PatchPVCns)
	router.DELETE("/api/v1/namespaces/:namespace/persistentvolumeclaims/:name", persistentvolumeclaim.DeletePVCns)
	router.DELETE("/api/v1/namespaces/:namespace/persistentvolumeclaims", persistentvolumeclaim.DeleteAllPVCns)
	router.GET("/api/v1/namespaces/:namespace/persistentvolumeclaims/:name/status", persistentvolumeclaim.GetstatusofPVCns)
	router.PUT("/api/v1/namespaces/:namespace/persistentvolumeclaims/:name/status", persistentvolumeclaim.UpdatestatusofPVCns)
	router.PATCH("/api/v1/namespaces/:namespace/persistentvolumeclaims/:name/status", persistentvolumeclaim.PatchstatusofPVCns)

	//lapi
	router.POST("/lapi/v1/orgs", lapi.CreateProject)
	router.GET("/lapi/v1/orgs/:project/roles", lapi.ListMembers)
	router.PUT("/lapi/v1/orgs/:project/invite", lapi.InviteMember)
	router.PUT("/lapi/v1/orgs/:project/remove", lapi.RemoveMember)

	//apis
	router.GET("/apis/autoscaling/v1/namespaces/:namespace/horizontalpodautoscalers/:name", apis.GetHPAns)
	router.PUT("/apis/autoscaling/v1/namespaces/:namespace/horizontalpodautoscalers/:name", apis.UpdateHPAns)
	router.PATCH("/apis/autoscaling/v1/namespaces/:namespace/horizontalpodautoscalers/:name", apis.PatchHPAns)
	router.DELETE("/apis/autoscaling/v1/namespaces/:namespace/horizontalpodautoscalers/:name", apis.DeleteHPAns)
	router.GET("/apis/apps/v1beta1/namespaces/:namespace/statefulsets/:name", apis.GetSFSns)
	router.PUT("/apis/apps/v1beta1/namespaces/:namespace/statefulsets/:name", apis.UpdateSFSns)
	router.PATCH("/apis/apps/v1beta1/namespaces/:namespace/statefulsets/:name", apis.PatchSFSns)
	router.DELETE("/apis/apps/v1beta1/namespaces/:namespace/statefulsets/:name", apis.DeleteSFSns)
	router.GET("/apis/extensions/v1beta1/namespaces/:namespace/deployments/:name", apis.GetDeploymentns)
	router.PUT("/apis/extensions/v1beta1/namespaces/:namespace/deployments/:name", apis.UpdateDeploymentns)
	router.PATCH("/apis/extensions/v1beta1/namespaces/:namespace/deployments/:name", apis.PatchDeploymentns)
	router.DELETE("/apis/extensions/v1beta1/namespaces/:namespace/deployments/:name", apis.DeleteDeploymentns)
	router.GET("/apis/extensions/v1beta1/namespaces/:namespace/deployments/:name/scale", apis.GetDeploymentScalens)
	router.PUT("/apis/extensions/v1beta1/namespaces/:namespace/deployments/:name/scale", apis.UpdateDeploymentScalens)
	router.PATCH("/apis/extensions/v1beta1/namespaces/:namespace/deployments/:name/scale", apis.PatchDeploymentScalens)
	router.GET("/apis/extensions/v1beta1/namespaces/:namespace/replicasets/:name", apis.GetRSns)
	router.PUT("/apis/extensions/v1beta1/namespaces/:namespace/replicasets/:name", apis.UpdateRSns)
	router.PATCH("/apis/extensions/v1beta1/namespaces/:namespace/replicasets/:name", apis.PatchRSns)
	router.DELETE("/apis/extensions/v1beta1/namespaces/:namespace/replicasets/:name", apis.DeleteRSns)
	router.GET("/apis/extensions/v1beta1/namespaces/:namespace/replicasets", apis.WatchAllRSns)
	router.GET("/apis/extensions/v1beta1/namespaces/:namespace/deployments", apis.WatchAllDeployns)

	// 未知路由处理
	router.NoRoute(others.AnyRequest)
	// 未知调用方式
	router.NoMethod(func(context *gin.Context) {
		context.String(404, "Not method")
	})
	return
}
