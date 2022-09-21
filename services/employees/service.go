/**
 * Service Layer implementation goes here.
 *
**/

package services

import (
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
func (ts *TemplateService) CreateEmployee(input common.CreateEmployeeInput) common.Response {

	IdExists, iderrcode, _ := ts.Employee_Repo.IsIdExists(input.Metadata["user_id"].(string))
	if IdExists {
		usernotexits, errcode, _ := ts.Employee_Repo.IsUserNotExists(input.Metadata["user_id"].(string))
		if usernotexits {
			id, err := ts.Employee_Repo.Insert(input.Metadata)
			if err != nil {
				openlog.Error("Error occured while creating employee")
				return common.ResponseHandler("803", "en", 0, nil)
			}
			user, errcode, err := ts.Employee_Repo.Fetch(id)
			if err != nil {
				return common.ResponseHandler(errcode, "en", 0, err.Error())
			}

			return common.ResponseHandler("806", "en", 1, user)
		}
		return common.ResponseHandler(errcode, "en", 0, nil)

	}
	return common.ResponseHandler(iderrcode, "en", 0, nil)
}

func (ts *TemplateService) FetchAllEmployees(input common.FetchAllEmployeesInput) common.Response {
	res, errcode, tcount, err := ts.Employee_Repo.FetchAll(input.Filters, input.Page, input.Size)
	if err != nil {
		openlog.Error("Error occured while fetching all users")
		return common.ResponseHandler(errcode, "en", 0, err.Error())
	}
	return common.ResponseHandler(errcode, "en", tcount, res)
}

func (ts *TemplateService) FetchEmployee(input common.FetchEmployeeInput) common.Response {
	res, errcode, err := ts.Employee_Repo.Fetch(input.ID)
	if err != nil {
		openlog.Error("Error occured while fetching employee")
		return common.ResponseHandler(errcode, "en", 0, err.Error())
	}
	return common.ResponseHandler("812", "en", 1, res)
}

func (ts *TemplateService) UpdateEmployee(input common.UpdateEmployeeInput) common.Response {

	_, errcode, err := ts.Employee_Repo.Fetch(input.ID)
	if err == nil {
		_, errcode, err := ts.Employee_Repo.Update(input.ID, input.Metadata)
		if err != nil {
			openlog.Error("Error occured while updating user")
			return common.ResponseHandler(errcode, "en", 0, err.Error())
		}
		res, errcode, err := ts.Employee_Repo.Fetch(input.ID)
		if err != nil {
			return common.ResponseHandler(errcode, "en", 0, err.Error())
		}
		return common.ResponseHandler("808", "en", 1, res)
	}
	return common.ResponseHandler(errcode, "en", 0, err.Error())
}

func (ts *TemplateService) DeleteEmployee(input common.DeleteEmployeeInput) common.Response {
	errcode, err := ts.Employee_Repo.Delete(input.ID)
	if err != nil {
		openlog.Error("Error occured while deleting user")
		return common.ResponseHandler(errcode, "en", 0, err.Error())
	}
	return common.ResponseHandler("810", "en", 1, nil)
}
