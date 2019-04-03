package processedtemplate

import (
	"encoding/json"
	oapi "github.com/asiainfoldp/datafoundry-gw/apirequest"
	"github.com/asiainfoldp/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var logger lager.Logger
var cdStatusMap map[string]int

func init() {
	logger = lager.NewLogger("oapi_v1_ProcessedTemplate")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func CDProcessedTemplate(c *gin.Context) {
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	namespace := c.Param("namespace")
	urlParas := pkg.SliceURL(c.Request.URL.String())
	rBody, _ := ioutil.ReadAll(c.Request.Body)

	cdptMap, err := pkg.BreakBody(rBody)
	if err != nil {
		logger.Error("CDProcessedTemplate BreakBody error ", err)
	}
	cdStatusMap = make(map[string]int)

	for k, v := range cdptMap {
		req, err := oapi.GenRequest("POST", host+"/oapi/v1/namespaces/"+namespace+"/"+strings.ToLower(k)+"s"+urlParas, token, v)

		if err != nil {
			logger.Error("Create A "+k+" Fail", err)
		}
		logger.Info("Create "+k, map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
		defer req.Body.Close()
		cdStatusMap[k] = req.StatusCode
	}
	b, _ := json.Marshal(cdStatusMap)

	c.Data(http.StatusOK, "application/json", b)
}
