/**
 * Server starts here
 * for reference
 * - openlog - https://github.com/go-chassis/openlog/blob/master/openlog.go
 * - Open architecture - https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047
**/

package main

import (
	"encoding/json"
	_ "user-service/chassisHandlers"
	"user-service/common"
	"user-service/database"
	user_repository "user-service/repository/users"
	resource "user-service/resource"
	user_services "user-service/services/users"

	"io/ioutil"
	"log"

	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"
)

func getService() *user_services.TemplateService {
	user_repo := user_repository.TemplateRepository{DbClient: database.GetClient(), DatabaseName: "chandan"}
	return &user_services.TemplateService{User_Repo: &user_repo}
}

func LoadErrors(errors []map[string]interface{}) {
	res := make(map[string]interface{})
	for _, err := range errors {
		res[err["errcode"].(string)] = err
	}
	common.ErrorMessages = res
}

func main() {
	bytes, err := ioutil.ReadFile("conf/errcodes.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	errors := make([]map[string]interface{}, 0)
	json.Unmarshal(bytes, &errors)
	LoadErrors(errors)
	temp_resource := resource.TemplateResource{}
	chassis.RegisterSchema("rest", &temp_resource)
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}

	// Add database configurations to archaius
	if err := archaius.AddFile("./conf/database.yaml"); err != nil {
		openlog.Error("add props configurations failed." + err.Error())
		return
	}
	// Add schema paths configurations to archaius
	if err := archaius.AddFile("./conf/payloadSchemas.yaml"); err != nil {
		openlog.Error("add props configurations failed." + err.Error())
		return
	}

	// Server will not start if error occurs.
	if err := database.Connect(); err != nil {
		openlog.Fatal("Error occured while connecting to database")
		return
	}

	// Inject service into resource
	temp_resource.Inject(getService())

	chassis.Run()
}
