/**
 * Service Layer implementation goes here.
 *
**/

package services

import (
	"UserManagement/client"
	common "UserManagement/common"
	employee_repository "UserManagement/repository/employees"
	user_repository "UserManagement/repository/users"

	"github.com/go-chassis/openlog"
)

type TemplateService struct {
	User_Repo     user_repository.TemplateRepositoryInterface
	Employee_Repo employee_repository.TemplateRepositoryInterface
}

// sample service implementation of user
func (ts *TemplateService) CreateUser(input common.CreateUserInput) common.Response {

	usernotexits, errcode, _ := ts.User_Repo.IsNameNotExists(input.Metadata["name"].(string))
	if usernotexits {
		id, err := ts.User_Repo.Insert(input.Metadata)

		if err != nil {
			openlog.Error("Error occured while creating user")
			return common.ResponseHandler("702", "en", 0, nil)
		}
		user, errcode, err := ts.User_Repo.Fetch(id)
		if err != nil {
			return common.ResponseHandler(errcode, "en", 0, err.Error())
		}

		return common.ResponseHandler("703", "en", 1, user)
	}
	return common.ResponseHandler(errcode, "en", 0, nil)

}

func (ts *TemplateService) FetchAllUsers(input common.FetchAllUsersInput) common.Response {
	res, errcode, tcount, err := ts.User_Repo.FetchAll(input.Filters, input.Page, input.Size)
	if err != nil {
		openlog.Error("Error occured while fetching all users")
		return common.ResponseHandler(errcode, "en", 0, err.Error())
	}
	return common.ResponseHandler(errcode, "en", tcount, res)
}

func (ts *TemplateService) FetchUser(input common.FetchUserInput) common.Response {
	res1, errcode, err := ts.User_Repo.Fetch(input.ID)
	if err != nil {
		openlog.Error("Error occured while fetching user")
		return common.ResponseHandler(errcode, "en", 0, err.Error())
	}
	res2, errcode, err := ts.Employee_Repo.Fetch(input.ID)
	if err != nil {
		openlog.Error("Error occured while fetching Employee")
		return common.ResponseHandler(errcode, "en", 0, err.Error())
	}
	res1["user_details"] = res2
	return common.ResponseHandler("715", "en", 1, res2)
}
func (ts *TemplateService) UpdateUser(input common.UpdateUserInput) common.Response {

	userr, errcode, err := ts.User_Repo.Fetch(input.ID)
	if err == nil {
		if input.Metadata["name"] != nil {
			if userr["name"].(string) != input.Metadata["name"].(string) {
				IsNotExistingUser, errcode, res := ts.User_Repo.IsNameNotExists(input.Metadata["name"].(string))
				if IsNotExistingUser {
					delete(input.Metadata, "_id")
					_, update_errcode, err := ts.User_Repo.Update(input.ID, input.Metadata)
					if err != nil {
						return common.ResponseHandler(update_errcode, "en", 0, err.Error())
					}
					user, errcode, err := ts.User_Repo.Fetch(input.ID)
					if err != nil {
						return common.ResponseHandler(errcode, "en", 0, err.Error())
					}
					return common.ResponseHandler("717", "en", 1, user)
				}
				return common.ResponseHandler(errcode, "en", 0, res)
			}
			return common.ResponseHandler("722", "en", 0, nil)
		}

		delete(input.Metadata, "_id")
		_, errcode, err := ts.User_Repo.Update(input.ID, input.Metadata)
		if err != nil {
			return common.ResponseHandler(errcode, "en", 0, err.Error())
		}
		user, errcode, err := ts.User_Repo.Fetch(input.ID)
		if err != nil {
			return common.ResponseHandler(errcode, "en", 0, err.Error())
		}
		return common.ResponseHandler("717", "en", 1, user)

	}
	return common.ResponseHandler(errcode, "en", 0, err.Error())
}

// func (ts *TemplateService) DeleteUser(input common.DeleteUserInput) common.Response {

// 	errcode, err := ts.User_Repo.Delete(input.ID)
// 	if err != nil {
// 		openlog.Error("Error occured while deleting user")
// 		return common.ResponseHandler("718", "en", 0, err.Error())
// 	}
// 	return common.ResponseHandler(errcode, "en", 1, nil)
// }

func (ts *TemplateService) DeleteUser(input common.DeleteUserInput) common.Response {
	res, _, errrr := client.MakeRequest("http://61c2ffca70d48fba53ceea1d_ContainerManager/deletebyuid/"+input.ID, "DELETE", nil, nil)
	var stat float64 = 200
	if res["status"].(float64) == stat {
		errcode, err := ts.User_Repo.Delete(input.ID)
		if err != nil {
			openlog.Error("Error occured while deleting user")
			return common.ResponseHandler(errcode, "en", 0, err.Error())
		}
		return common.ResponseHandler(errcode, "en", 1, nil)
	}
	return common.ResponseHandler("818", "en", 0, errrr)
}

func (ts *TemplateService) DeleteAllUsers(input common.DeleteAllUsersInput) common.Response {
	errcode, err := ts.User_Repo.DeleteAll(input.Token)
	if err != nil {
		openlog.Error("Error occured while deleting all users")
		return common.ResponseHandler(errcode, "en", 0, err.Error())
	}
	return common.ResponseHandler(errcode, "en", 1, nil)
}
