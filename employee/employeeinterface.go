package employee

type EmployeeInterface interface {
	DeleteEmployee(id string, headers map[string]string) (map[string]interface{}, string, error)
	FetchEmployee(id string, headers map[string]string) (map[string]interface{}, string, error)
}
