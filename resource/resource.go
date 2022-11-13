/**
 * Resource definition will go here
 * for reference visit
 * - HTTP Router definition - https://github.com/go-chassis/go-chassis/blob/v2.1.0/server/restful/router.go
 * - Context - https://github.com/go-chassis/go-chassis/blob/v2.1.0/server/restful/context.go
 **/
package resource

import (
	"net/http"

	"user-service/common"
	user_services "user-service/services/users"

	"github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
)

type TemplateResource struct {
	us user_services.TemplateServiceInterface
}

// Services should get injected before using.
func (tr *TemplateResource) Inject(us user_services.TemplateServiceInterface) {
	tr.us = us
}

// creates the user
func (r *TemplateResource) CreateUser(context *restful.Context) {
	openlog.Info("Got a request to create user")
	// create a user map for payload
	user := make(map[string]interface{})
	// Read the payload into user from context
	err := context.ReadEntity(&user)
	if err != nil {
		openlog.Error(err.Error())
		// Send error response
		context.WriteHeaderAndJSON(400, common.ResponseHandler("701", "en", 0, nil), "application/json")
		return
	}
	// call service layer
	ip := common.CreateUserInput{Metadata: user, Language: "en"}
	res := r.us.CreateUser(ip)
	context.WriteHeaderAndJSON(res.Status, res, "application/json")
}

func (r *TemplateResource) FetchAllUsers(context *restful.Context) {
	openlog.Info("Got a request to get all users")
	//Get the filters from context
	filters, page, limit := context.ReadQueryParameter("filters"), context.ReadQueryParameter("page"), context.ReadQueryParameter("size")

	//call service layer
	ip := common.FetchAllUsersInput{Filters: filters, Page: page, Size: limit, Language: "en"}
	res := r.us.FetchAllUsers(ip)
	context.WriteHeaderAndJSON(res.Status, res, "application/json")
}

func (r *TemplateResource) FetchUser(context *restful.Context) {
	openlog.Info("Got a request to get user")
	// Get the user id from context
	id := context.ReadPathParameter("id")
	// call service layer
	ip := common.FetchUserInput{ID: id, Language: "en"}
	res := r.us.FetchUser(ip)
	//time.Sleep(35 * time.Second)
	context.WriteHeaderAndJSON(res.Status, res, "application/json")
}

func (r *TemplateResource) UpdateUser(context *restful.Context) {
	openlog.Info("Got a request to update user")
	// create a user map for payload
	user := make(map[string]interface{})
	// Read the payload into user from context
	err := context.ReadEntity(&user)
	if err != nil {
		openlog.Error(err.Error())
		// Send error response
		context.WriteHeaderAndJSON(400, common.ResponseHandler("701", "en", 0, nil), "application/json")
		return
	}

	id := context.ReadPathParameter("id")

	// call service layer
	ip := common.UpdateUserInput{ID: id, Metadata: user, Language: "en"}
	res := r.us.UpdateUser(ip)
	context.WriteHeaderAndJSON(res.Status, res, "application/json")
}

func (r *TemplateResource) DeleteUser(context *restful.Context) {
	openlog.Info("Got a request to delete user")
	// Get the user id from context
	id := context.ReadPathParameter("id")
	// call service layer
	ip := common.DeleteUserInput{ID: id, Language: "en"}
	res := r.us.DeleteUser(ip)
	//time.Sleep(35 * time.Second)
	context.WriteHeaderAndJSON(res.Status, res, "application/json")
}

func (r *TemplateResource) DeleteAllUsers(context *restful.Context) {
	openlog.Info("Got a request to get user")
	// Get the user id from context
	key := context.ReadPathParameter("key")
	// call service layer
	ip := common.DeleteAllUsersInput{Token: key, Language: "en"}
	res := r.us.DeleteAllUsers(ip)
	context.WriteHeaderAndJSON(res.Status, res, "application/json")
}

func (r *TemplateResource) Info(context *restful.Context) {
	openlog.Info("Got a request to get info")
	type Info struct {
		ServiceName string `json:"service_name"`
		Version     string `json:"version"`
	}
	res := Info{ServiceName: "User Service", Version: "1.0.0"}
	context.WriteHeaderAndJSON(200, res, "application/json")
}

// Define all APIs here.
func (r *TemplateResource) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodGet, Path: "/info", ResourceFunc: r.Info},
		{Method: http.MethodPost, Path: "/createUser", ResourceFunc: r.CreateUser, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
		{Method: http.MethodGet, Path: "/getAllUsers", ResourceFunc: r.FetchAllUsers, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
		{Method: http.MethodGet, Path: "/getUser/{id}", ResourceFunc: r.FetchUser, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
		{Method: http.MethodPut, Path: "/updateUser/{id}", ResourceFunc: r.UpdateUser, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
		{Method: http.MethodDelete, Path: "/deleteUser/{id}", ResourceFunc: r.DeleteUser, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
		{Method: http.MethodDelete, Path: "/deleteAllUsers/{key}", ResourceFunc: r.DeleteAllUsers, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},

		// {Method: http.MethodPost, Path: "/createEmployee", ResourceFunc: r.CreateEmployee, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
		// {Method: http.MethodGet, Path: "/getAllEmployees", ResourceFunc: r.FetchAllEmployees, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
		// {Method: http.MethodGet, Path: "/getEmployee/{id}", ResourceFunc: r.FetchEmployee, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
		// {Method: http.MethodPut, Path: "/updateEmployee/{id}", ResourceFunc: r.UpdateEmployee, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
		// {Method: http.MethodDelete, Path: "/deleteEmployee/{id}", ResourceFunc: r.DeleteEmployee, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
	}
}
