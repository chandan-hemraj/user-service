package common

type CreateUserInput struct {
	Metadata map[string]interface{}
	Language string
}

type FetchAllUsersInput struct {
	Filters  string
	Page     string
	Size     string
	Language string
}

type FetchUserInput struct {
	ID       string
	Language string
}

type UpdateUserInput struct {
	ID       string
	Metadata map[string]interface{}
	Language string
}

type DeleteUserInput struct {
	ID       string
	Language string
}

type DeleteAllUsersInput struct {
	Token    string
	Language string
}

type CreateEmployeeInput struct {
	Metadata map[string]interface{}
	Language string
}

type FetchAllEmployeesInput struct {
	Filters  string
	Page     string
	Size     string
	Language string
}

type FetchEmployeeInput struct {
	ID       string
	Language string
}

type UpdateEmployeeInput struct {
	ID       string
	Metadata map[string]interface{}
	Language string
}

type DeleteEmployeeInput struct {
	ID       string
	Language string
}

// type DeleteAllEmployeeInput struct {
// 	Token    string
// 	Language string
// }

var ErrorMessages map[string]interface{}
