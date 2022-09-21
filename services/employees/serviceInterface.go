/**
 * Inteface for service layer
 *
**/

package services

import (
	common "UserManagement/common"
)

type TemplateServiceInterface interface {
	CreateEmployee(common.CreateEmployeeInput) common.Response
	FetchAllEmployees(common.FetchAllEmployeesInput) common.Response
	FetchEmployee(common.FetchEmployeeInput) common.Response
	UpdateEmployee(common.UpdateEmployeeInput) common.Response
	DeleteEmployee(common.DeleteEmployeeInput) common.Response
}
