package imagestream

import (
	"os"
	//"fmt"
	oapi "github.com/asiainfoLDP/datafoundry-gw/apirequest"
	"github.com/asiainfoLDP/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
)

const (
	IMAGE       = "/oapi/v1/imagestreams"
	IMAGENAME   = "/oapi/v1/namespaces/"
	IMAGECONFIG = "/imagestreams/"
	WATCH       = "/oapi/v1/watch/namespaces/"
	WATCHALL    = "/oapi/v1/watch/imagestreams"
	JSON        = "application/json"
)

var log lager.Logger

func init() {
	log = lager.NewLogger("oapi_v1_ImageStream")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
}

func CreateIS(c *gin.Context) {
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateIS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("POST", IMAGE, token, rBody)
	if err != nil {
		log.Error("CreateIS error ", err)
	}
	log.Info("Create IS", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateIS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func CreateImageInNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateImageInNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("POST", IMAGENAME+namespace+"/imagestreams", token, rBody)
	if err != nil {
		log.Error("CreateBuildConfigInNameSpace error ", err)
	}
	log.Info("Create Image In NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateImageInNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetImageFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchImageFromNS(c)
	} else {
		namespace := c.Param("namespace")
		name := c.Param("name")
		token := pkg.GetToken(c)
		req, err := oapi.GenRequest("GET", IMAGENAME+namespace+IMAGECONFIG+name, token, []byte{})
		if err != nil {
			log.Error("GetImageFromNS error ", err)
		}
		log.Info("Cet Image From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetImageFromNS Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllImage(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllImage(c)
	} else {
		token := pkg.GetToken(c)
		req, err := oapi.GenRequest("GET", IMAGE, token, []byte{})
		if err != nil {
			log.Error("GetAllImage error ", err)
		}
		log.Info("List Image  ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetAllImage Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetAllImageFromNS(c *gin.Context) {
	if pkg.IsWebsocket(c) {
		watchAllImageFromNS(c)
	} else {
		namespace := c.Param("namespace")
		token := pkg.GetToken(c)
		req, err := oapi.GenRequest("GET", IMAGENAME+namespace+"/imagestreams", token, []byte{})
		if err != nil {
			log.Error("GetAllImageFromNS error ", err)
		}
		log.Info("List Image From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error("GetAllImageFromNS Read req.Body error", err)
		}
		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}

func GetSecImageFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := oapi.GenRequest("GET", IMAGENAME+namespace+IMAGECONFIG+name+"/secrets", token, []byte{})
	if err != nil {
		log.Error("GetSecImageFromNS error ", err)
	}
	log.Info("Get Sec Image From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetSecImageFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func GetStaImageFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	req, err := oapi.GenRequest("GET", IMAGENAME+namespace+IMAGECONFIG+name+"/status", token, []byte{})
	if err != nil {
		log.Error("GetStaImageFromNS error ", err)
	}
	log.Info("Get Sta Image From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetStaImageFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func watchImageFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetWSToken(c)
	log.Info("Watch A Image From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	oapi.WSRequest(WATCH+namespace+IMAGECONFIG+name, token, c.Writer, c.Request)
	log.Info("Watch A Image From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func watchAllImage(c *gin.Context) {
	token := pkg.GetWSToken(c)
	log.Info("Watch collection Image", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	oapi.WSRequest(WATCHALL, token, c.Writer, c.Request)
	log.Info("Watch collection Image", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func watchAllImageFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetWSToken(c)
	log.Info("Watch collection Image From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "start watch"})
	oapi.WSRequest(WATCH+namespace+"/imagestreams", token, c.Writer, c.Request)
	log.Info("Watch collection Image From NameSpace", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": "end watch"})
}

func UpdataImageFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataImageFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("PUT", IMAGENAME+namespace+IMAGECONFIG+name, token, rBody)
	if err != nil {
		log.Error("UpdataImageFromNS error ", err)
	}
	log.Info("Updata Image From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataImageFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func UpdataStaImageFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("UpdataStaImageFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("PUT", IMAGENAME+namespace+IMAGECONFIG+name+"/status", token, rBody)
	if err != nil {
		log.Error("UpdataStaImageFromNS error ", err)
	}
	log.Info("Updata Status Image From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdataStaImageFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchImageFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchImageFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("PATCH", IMAGENAME+namespace+IMAGECONFIG+name, token, rBody)
	if err != nil {
		log.Error("PatchImageFromNS error ", err)
	}
	log.Info("Patch Status Image From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchImageFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func PatchStaImageFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("PatchStaImageFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("PATCH", IMAGENAME+namespace+IMAGECONFIG+name+"/status", token, rBody)
	if err != nil {
		log.Error("PatchStaImageFromNS error ", err)
	}
	log.Info("Patch Status Image From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("PatchStaImageFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteImageFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteImageFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("DELETE", IMAGENAME+namespace+IMAGECONFIG+name, token, rBody)
	if err != nil {
		log.Error("DeleteImageFromNS error ", err)
	}
	log.Info("Delete Image From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteImageFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func DeleteAllImageFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	token := pkg.GetToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("DeleteAllImageFromNS Read Request.Body error", err)
	}
	defer c.Request.Body.Close()
	req, err := oapi.GenRequest("DELETE", IMAGENAME+namespace+"/imagestreams", token, rBody)
	if err != nil {
		log.Error("DeleteAllImageFromNS error ", err)
	}
	log.Info("Delete Collection Image From NameSpace ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteAllImageFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
