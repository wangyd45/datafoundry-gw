package lapi

import (
	"encoding/json"
	"errors"
	oapi "github.com/asiainfoldp/datafoundry-gw/apirequest"
	"github.com/asiainfoldp/datafoundry-gw/pkg"
	"github.com/gin-gonic/gin"
	projectapi "github.com/openshift/project/api/v1"
	rolebindingapi "github.com/openshift/rolebinding/api/v1"
	userapi "github.com/openshift/user/api/v1"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	kapi "k8s.io/kubernetes/pkg/api/v1"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var log lager.Logger

func init() {
	log = lager.NewLogger("lapi_v1.log")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) //默认日志级别
}

type OrgMember struct {
	Name         string            `json:"member_name"`
	IsAdmin      bool              `json:"admin"`
	PrivilegedBy string            `json:"privileged_by"`
	JoinedAt     string            `json:"joined_at"`
	Status       MemberStatusPhase `json:"status"`
}

type Orgnazition struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	CreateBy    string      `json:"create_by"`
	CreateTime  string      `json:"creation_time"`
	MemberList  []OrgMember `json:"members"`
	Status      string      `json:"status"`
	RoleBinding bool        `json:"rolebinding"`
	Reason      string      `json:"reason,omitempty"`
}

type MemberStatusPhase string

func CreateProject(c *gin.Context) {
	org := new(Orgnazition)
	if err := parseRequestBody(c.Request, org); err != nil {
		log.Error("read request body error ", err)
		return
	}
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	if "" == token {
		log.Error("get token error ", nil)
		rstring := "{\"messages\": \"Unauthorized\",\"reason\": \"TokenNull\"}"
		c.Data(403, "application/json", []byte(rstring))
		return
	}
	user, err := authDF(host,token)
	if err != nil {
		log.Error("get user error!", err)
		return
	}

	projRequest := new(projectapi.ProjectRequest)
	{
		projRequest.DisplayName = org.Name
		projRequest.Name = user + "-org-" + genRandomName(8)
	}
	rBody, err := json.Marshal(projRequest)
	if err != nil {
		log.Error("json Masrshal error ", err)
		return
	}
	req, err := oapi.GenRequest("POST",host + "/oapi/v1/projectrequests", token, rBody)

	if err != nil {
		log.Error("Create A Project Fail", err)
		return
	}
	//返回结果JSON格式
	log.Info("Create Project ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateProject Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)
}

func ListMembers(c *gin.Context) {
	project := c.Param("project")
	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	if "" == token {
		log.Error("get token error ", nil)
		rstring := "{\"messages\": \"Unauthorized\",\"reason\": \"TokenNull\"}"
		c.Data(403, "application/json", []byte(rstring))
		return
	}
	req, err := oapi.GenRequest("GET",host + "/oapi/v1/namespaces/"+project+"/rolebindings", token, nil)

	if err != nil {
		log.Error("ListMembers error ", err)
	}
	log.Info("List Members ", map[string]interface{}{"user": pkg.GetUserFromToken(pkg.SliceToken(token)), "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("ListMembers Read req.Body error", err)
	}
	defer req.Body.Close()
	roles := new(rolebindingapi.RoleBindingList)
	rolesResult := new(rolebindingapi.RoleBindingList)
	err = json.Unmarshal(result, roles)
	if err != nil {
		log.Error("json Unmarshal error ", err)
		return
	}
	for _, role := range roles.Items {
		if role.Name == "view" || role.Name == "admin" || role.Name == "edit" {
			rolesResult.Items = append(rolesResult.Items, role)
		} else {
			for _, subject := range role.Subjects {
				if subject.Kind == "User" {
					if role.RoleRef.Name == "view" || role.RoleRef.Name == "admin" ||
						role.RoleRef.Name == "edit" {
						//clog.Debugf("%#v", role)
						rolesResult.Items = append(rolesResult.Items, role)
						break
					}
				}
			}
		}
	}
	res, err := json.Marshal(rolesResult)
	if err != nil {
		log.Error("json Masrshal error ", err)
		return
	}
	c.Data(req.StatusCode, "application/json", res)

}

func InviteMember(c *gin.Context) {

	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	if "" == token {
		log.Error("get token error ", nil)
		rstring := "{\"messages\": \"Unauthorized\",\"reason\": \"TokenNull\"}"
		c.Data(403, "application/json", []byte(rstring))
		return
	}
	datainfo := make(map[string]interface{})
	datainfo["RemoteAddr"] = c.Request.RemoteAddr
	datainfo["Method"] = c.Request.Method
	datainfo["URL"] = c.Request.URL.RequestURI()
	datainfo["Proto"] = c.Request.Proto
	log.Info("from", datainfo)

	member := new(OrgMember)

	if err := parseRequestBody(c.Request, member); err != nil {
		log.Error("InviteMember read request body error.", err)
		return
	}

	project := c.Param("project")
	req, err := roleAdd(host,token, project, member.Name, member.IsAdmin)

	if err != nil {
		log.Error("InviteMember error ", err)
		rstring := "{\"messages\": \"" + err.Error() + "\"}"
		c.Data(409, "application/json", []byte(rstring))
		return
	}
	log.Info("Invite Members ", map[string]interface{}{"user": "", "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("ListMembers Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, "application/json", result)

}

func RemoveMember(c *gin.Context) {

	token := pkg.GetToken(c)
	host := pkg.GetHost(c)
	if "" == token {
		log.Error("get token error ", nil)
		rstring := "{\"messages\": \"Unauthorized\",\"reason\": \"TokenNull\"}"
		c.Data(403, "application/json", []byte(rstring))
		return
	}
	datainfo := make(map[string]interface{})
	datainfo["RemoteAddr"] = c.Request.RemoteAddr
	datainfo["Method"] = c.Request.Method
	datainfo["URL"] = c.Request.URL.RequestURI()
	datainfo["Proto"] = c.Request.Proto
	log.Info("from", datainfo)

	member := new(OrgMember)

	if err := parseRequestBody(c.Request, member); err != nil {
		log.Error("read request body error.", err)
		return
	}

	project := c.Param("project")

	req, err := roleRemove(host,token, project, member.Name)
	if err != nil {
		log.Error("RemoveMember error ", err)
		rstring := "{\"messages\": \"" + err.Error() + "\"}"
		c.Data(404, "application/json", []byte(rstring))
		return
	}
	log.Info("Remove Members ", map[string]interface{}{"user": "", "time": pkg.GetTimeNow(), "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Error("ListMembers Read req.Body error", err)
	}
	c.Data(req.StatusCode, "application/json", result)
}

//获取长度为strlen的随机码
func genRandomName(strlen int) (name string) {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

//根据Token获取用户名称
func authDF(host,token string) (string, error) {
	u := &userapi.User{}
	req, err := oapi.GenRequest("GET",host + "/oapi/v1/users/~", token, []byte{})

	if err != nil {
		log.Error("GetUser error ", err)
	}
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetUser ioutil.ReadAll error ", err)
	}
	defer req.Body.Close()
	err = json.Unmarshal(result, u)
	if err != nil {
		log.Error("json unmarshal error ", err)
		return "", err
	}
	return u.Name, nil
}

//检验请求体正确性
func parseRequestBody(r *http.Request, v interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	log.Debug("Request Body:" + string(b))
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	return nil
}

//移除rolebindings
func roleRemove(host,token, project, name string) (*http.Response, error) {

	var req *http.Response
	var err error

	if name == "" || project == "" {
		return nil, errors.New("namespace or username is null")
	}

	println("project:" + project)
	roleList, _ := getListRoles(host,token, project)
	if err != nil {
		return nil, err
	}

	role := findUserInRoles(roleList, name)
	if role == nil {
		return nil, errors.New("can't find user '" + name + "' from roles in project '" + project + "'")
	} else {
		role = removeUserInRole(role, name)
		body, _ := json.Marshal(role)
		req, err = oapi.GenRequest("PUT",host + "/oapi/v1/namespaces/"+project+"/rolebindings/"+role.Name, token, body)

		return req, err
	}

}

//创建或更新rolebindings
func roleAdd(host,token, project, name string, admin bool) (*http.Response, error) {

	var req *http.Response
	var err error

	if name == "" || project == "" {
		return nil, errors.New("namespace or username is null")
	}

	roleList, err := getListRoles(host,token, project)
	if err != nil {
		return nil, err
	}

	if exist := findUserInRoles(roleList, name); exist != nil {
		return nil, errors.New("duplicate user: " + name + ", role: " + exist.RoleRef.Name + ", project: " + project + ", reason: exist")
	}

	roleRef := "edit"
	if admin {
		roleRef = "admin"
	}

	role := findRole(roleList, roleRef)
	create := false

	if role == nil { //post else put

		create = true
		role = new(rolebindingapi.RoleBinding)
		role.Name = roleRef
		role.RoleRef.Name = roleRef
	}

	subject := kapi.ObjectReference{Kind: "User", Name: name}
	role.Subjects = append(role.Subjects, subject)
	role.UserNames = append(role.UserNames, name)

	if role.Annotations == nil {
		role.Annotations = make(map[string]string)
	}
	role.Annotations["joinedTime/"+name] = time.Now().Format(time.RFC3339)

	body, _ := json.Marshal(role)
	if create {
		req, err = oapi.GenRequest("POST",host + "/oapi/v1/namespaces/"+project+"/rolebindings", token, body)


	} else {
		req, err = oapi.GenRequest("PUT",host + "/oapi/v1/namespaces/"+project+"/rolebindings/"+roleRef, token, body)

	}

	return req, err
}

//获取rolebindings列表
func getListRoles(host,token, project string) (*rolebindingapi.RoleBindingList, error) {

	var err error

	roles := new(rolebindingapi.RoleBindingList)

	resp, err := oapi.GenRequest("GET",host + "/oapi/v1/namespaces/"+project+"/rolebindings", token, nil)

	if err != nil {
		log.Error("Get All RoleBindings In A Namespace Fail", err)
		return nil, err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(roles)

	rolesResult := new(rolebindingapi.RoleBindingList)

	for _, role := range roles.Items {
		if role.Name == "view" || role.Name == "admin" || role.Name == "edit" {
			rolesResult.Items = append(rolesResult.Items, role)
		} else {
			for _, subject := range role.Subjects {
				if subject.Kind == "User" {
					if role.RoleRef.Name == "view" || role.RoleRef.Name == "admin" ||
						role.RoleRef.Name == "edit" {
						//clog.Debugf("%#v", role)
						rolesResult.Items = append(rolesResult.Items, role)
						break
					}
				}
			}
		}
	}
	return rolesResult, nil
}

//根据Role名称查询rolebindings
func findRole(roles *rolebindingapi.RoleBindingList, roleRef string) *rolebindingapi.RoleBinding {
	for _, role := range roles.Items {
		if role.Name == roleRef {
			return &role
		}
	}
	return nil
}

//根据用户名查询rolebindings
func findUserInRoles(roles *rolebindingapi.RoleBindingList, username string) *rolebindingapi.RoleBinding {
	for _, role := range roles.Items {
		for _, v := range role.UserNames {
			if username == v {
				return &role
			}
		}
	}
	return nil
}

//根据用户名删除rolebindings
func removeUserInRole(role *rolebindingapi.RoleBinding, user string) *rolebindingapi.RoleBinding {
	for idx, userName := range role.UserNames {
		if userName == user {
			role.UserNames = append(role.UserNames[:idx], role.UserNames[idx+1:]...)
		}
	}
	for idx, subject := range role.Subjects {
		if subject.Name == user {
			role.Subjects = append(role.Subjects[:idx], role.Subjects[idx+1:]...)
		}
	}

	delete(role.Annotations, "joinedTime/"+user)

	return role
}
