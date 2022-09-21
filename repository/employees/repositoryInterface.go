/**
 * Interface for implementation of repositories.
**/

package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type TemplateRepositoryInterface interface {
	Insert(s map[string]interface{}) (string, error)
	IsIdExists(id string) (bool, string, error)
	IsUserNotExists(id string) (bool, string, error)
	FetchAll(filters string, page string, size string) ([]primitive.M, string, int, error)
	Fetch(id string) (map[string]interface{}, string, error)
	Update(id string, data map[string]interface{}) (map[string]interface{}, string, error)
	Delete(id string) (string, error)
}
