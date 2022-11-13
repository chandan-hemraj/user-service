/**
 * Inteface for service layer
 *
**/

package services

import (
	common "user-service/common"
)

type TemplateServiceInterface interface {
	CreateUser(common.CreateUserInput) common.Response
	FetchAllUsers(common.FetchAllUsersInput) common.Response
	FetchUser(common.FetchUserInput) common.Response
	UpdateUser(common.UpdateUserInput) common.Response
	DeleteUser(common.DeleteUserInput) common.Response
	DeleteAllUsers(common.DeleteAllUsersInput) common.Response
}
