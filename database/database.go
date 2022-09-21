/**
 * Database maintain the connection with mongodb.
 * for reference visit -
 *
**/

package database

import (
	"context"
	"sync"

	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/openlog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Database struct{}

var client mongo.Client
var once sync.Once

// Reads config from database.yaml
// * Note it has to be add to archaius.
// database.mongodb.uri => represents mongodb url - Mandatory
// database.mongodb.poolsize => represents pool size - Non Mandatory.
func Connect() error {
	uri := archaius.GetString("database.mongodb.uri", "")
	poolsize := archaius.GetInt64("database.mongodb.poolsize", 15) // by defaylt pool size is 15.
	clientOptions := options.Client().ApplyURI(uri).SetMaxPoolSize(uint64(poolsize))
	clientlocal, err := mongo.Connect(context.TODO(), clientOptions)
	client = *clientlocal
	if err != nil {
		return err
	}
	err = clientlocal.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	openlog.Info("Connected to Mongodb")
	return nil
}

func GetClient() *mongo.Client { return &client }
