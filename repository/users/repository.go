/**
 * Used for Database operations
 * for reference visit :
 * 		1- converting mongdb id into string - https://stackoverflow.com/questions/60864873/primitive-objectid-to-string
 *		2- Mongodb driver API reference - https://godoc.org/go.mongodb.org/mongo-driver/mongo
**/

package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/go-chassis/openlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TemplateRepository struct {
	DbClient     *mongo.Client
	DatabaseName string
}

func (tr *TemplateRepository) IsNameNotExists(name string) (bool, string, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("users")
	res := collection.FindOne(context.Background(), bson.M{"name": name})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return true, "", nil
		}
		return false, "704", nil
	}
	return false, "705", errors.New("user already exists, Please change name")
}

// insert interface and return mongodb id if success else error
func (tr *TemplateRepository) Insert(s map[string]interface{}) (string, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("users")
	res, err := collection.InsertOne(context.Background(), s)
	if err != nil {
		return "", err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex() // converting mongdb into string
	return id, nil
}

func (tr *TemplateRepository) FetchAll(filters string, page string, limit string) ([]primitive.M, string, int, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("users")
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}
	var results []primitive.M
	pag, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		openlog.Error(err.Error())
		return nil, "707", 0, err
	}
	lim, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		log.Println(err)
		return nil, "708", 0, err
	}
	skip := (pag - 1) * lim
	opts := options.FindOptions{Skip: &skip, Limit: &lim}
	filter := map[string]interface{}{}
	err = json.Unmarshal([]byte(filters), &filter)
	var cur *mongo.Cursor
	var errr error
	var count int64
	var errrr error
	if err != nil && filters != "" {
		count, errrr = collection.CountDocuments(context.TODO(), bson.M{})
		if errrr != nil {
			if errrr == mongo.ErrNoDocuments {
				log.Println(errrr)
				return nil, "709", 0, errrr
			}
			log.Println(errrr)
			return nil, "710", 0, errrr
		}
		cur, errr = collection.Find(context.TODO(), bson.M{}, &opts)
	} else {
		count, errrr = collection.CountDocuments(context.TODO(), filter)
		if errrr != nil {
			if errrr == mongo.ErrNoDocuments {
				log.Println(errrr)
				return nil, "709", 0, errrr
			}
			fmt.Println(errrr)
			return nil, "710", 0, errrr
		}
		cur, errr = collection.Find(context.TODO(), filter, &opts)
	}
	if count == 0 {
		log.Println("No Documents found")
		return nil, "724", 0, errors.New("no documents found")
	}
	if errr != nil {
		if errr == mongo.ErrNoDocuments {
			log.Println(errr)
			return nil, "709", 0, errr
		}
		log.Println(errr)
		return nil, "710", 0, errr
	}
	for cur.Next(context.TODO()) {
		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
			return nil, "711", 0, err
		}
		results = append(results, elem)
	}
	cur.Close(context.TODO())
	return results, "712", int(count), nil
}

func (tr *TemplateRepository) Fetch(id string) (map[string]interface{}, string, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, "713", err
	}
	filter := map[string]interface{}{"_id": objID}
	res := collection.FindOne(context.Background(), filter)
	result := make(map[string]interface{})
	err = res.Decode(&result)
	if err != nil {
		return nil, "714", err
	}
	return result, "", nil
}

func (tr *TemplateRepository) Update(id string, user map[string]interface{}) (map[string]interface{}, string, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("users")
	result := make(map[string]interface{})
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, "713", err
	}
	err = collection.FindOneAndUpdate(context.Background(), map[string]interface{}{"_id": _id}, map[string]interface{}{"$set": user}).Decode(&result)
	if err != nil {
		return result, "716", err
	}
	return result, "717", nil
}
func (tr *TemplateRepository) Delete(id string) (string, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("users")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "713", err
	}
	res := collection.FindOneAndDelete(context.Background(), bson.M{"_id": oid})
	result := make(map[string]interface{})
	err = res.Decode(&result)
	if err != nil {
		return "718", err
	}
	return "719", nil
}

func (tr *TemplateRepository) DeleteAll(key string) (string, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("users")
	if key == "admin" {
		_, err := collection.DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return "725", err
		}
		return "727", nil
	}
	return "728", errors.New("You are not authorized to delete all records")
}
