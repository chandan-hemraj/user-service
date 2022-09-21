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

func (tr *TemplateRepository) IsIdExists(id string) (bool, string, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, "804", err
	}
	res := collection.FindOne(context.TODO(), bson.M{"_id": objID})
	fmt.Println(res.Err())
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return false, "801", errors.New("no documents found")
		}
		return false, "802", errors.New("internal server error")
	}
	return true, "", nil
}

func (tr *TemplateRepository) IsUserNotExists(id string) (bool, string, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("employees")
	res := collection.FindOne(context.TODO(), bson.M{"user_id": id})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return true, "", nil
		}
		return false, "802", errors.New("internal server error")
	}
	return false, "807", errors.New("employee already exists")
}

// insert interface and return mongodb id if success else error
func (tr *TemplateRepository) Insert(s map[string]interface{}) (string, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("employees")
	delete(s, "_id")
	//delete(s, "user_id")
	res, err := collection.InsertOne(context.Background(), s)
	if err != nil {
		return "", err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex() // converting mongdb into string
	return id, nil
}

func (tr *TemplateRepository) Fetch(id string) (map[string]interface{}, string, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("employees")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, "804", err
	}
	filter := map[string]interface{}{"_id": objID}
	res := collection.FindOne(context.Background(), filter)
	result := make(map[string]interface{})
	err = res.Decode(&result)
	if err != nil {
		return nil, "805", err
	}

	uid := result["user_id"].(string)
	ID, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return nil, "804", err
	}
	collection = tr.DbClient.Database(tr.DatabaseName).Collection("users")
	filter = map[string]interface{}{"_id": ID}
	r := collection.FindOne(context.Background(), filter)
	resul := make(map[string]interface{})
	er := r.Decode(&resul)
	if er != nil {
		return nil, "805", er
	}
	result["user_details"] = resul
	delete(result, "user_id")
	return result, "", nil
}

// func (tr *TemplateRepository) Fetch(id string) (map[string]interface{}, string, error) {
// 	collection := tr.DbClient.Database(tr.DatabaseName).Collection("employees")
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, "804", err
// 	}
// 	filter := map[string]interface{}{"_id": objID}
// 	stage1 := bson.D{{"$match", filter}}
// 	stage2 := bson.D{{"$lookup", bson.D{{"from", "users"}, {"localField", "user_id"}, {"foreignField", "_id"}, {"as", "user_details"}}}}
// 	ress, errr := collection.Aggregate(context.Background(), mongo.Pipeline{
// 		stage1,
// 		stage2,
// 	})
// 	if errr != nil {
// 		return nil, "805", errr
// 	}
// 	result := make(map[string]interface{})
// 	fmt.Println("herererreeeeeeeeeeee")
// 	err = ress.Decode(&result)
// 	if err != nil {
// 		return nil, "805", err
// 	}
// 	return result, "", nil
// }

func (tr *TemplateRepository) FetchAll(filters string, page string, limit string) ([]primitive.M, string, int, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("employees")
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
				return nil, "813", 0, errrr
			}
			log.Println(errrr)
			return nil, "814", 0, errrr
		}
		cur, errr = collection.Find(context.TODO(), bson.M{}, &opts)
	} else {
		count, errrr = collection.CountDocuments(context.TODO(), filter)
		if errrr != nil {
			if errrr == mongo.ErrNoDocuments {
				log.Println(errrr)
				return nil, "813", 0, errrr
			}
			fmt.Println(errrr)
			return nil, "814", 0, errrr
		}
		cur, errr = collection.Find(context.TODO(), filter, &opts)
	}
	if count == 0 {
		log.Println("No Documents found")
		return nil, "815", 0, errors.New("no documents found")
	}
	if errr != nil {
		if errr == mongo.ErrNoDocuments {
			log.Println(errr)
			return nil, "813", 0, errr
		}
		log.Println(errr)
		return nil, "814", 0, errr
	}
	for cur.Next(context.TODO()) {
		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
			return nil, "816", 0, err
		}
		uid := elem["user_id"].(string)
		ID, err := primitive.ObjectIDFromHex(uid)
		if err != nil {
			return nil, "804", 0, err
		}
		collection = tr.DbClient.Database(tr.DatabaseName).Collection("users")
		filter = map[string]interface{}{"_id": ID}
		r := collection.FindOne(context.Background(), filter)
		result := make(map[string]interface{})
		er := r.Decode(&result)
		if er != nil {
			return nil, "805", 0, er
		}
		elem["user_details"] = result
		delete(elem, "user_id")
		results = append(results, elem)
	}
	cur.Close(context.TODO())
	return results, "817", int(count), nil
}

func (tr *TemplateRepository) Update(id string, data map[string]interface{}) (map[string]interface{}, string, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("employees")
	result := make(map[string]interface{})
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, "804", err
	}
	delete(data, "_id")
	delete(data, "user_id")
	err = collection.FindOneAndUpdate(context.Background(), map[string]interface{}{"_id": _id}, map[string]interface{}{"$set": data}).Decode(&result)
	if err != nil {
		return result, "809", err
	}
	return result, "", nil
}

func (tr *TemplateRepository) Delete(id string) (string, error) {
	collection := tr.DbClient.Database(tr.DatabaseName).Collection("employees")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "804", err
	}
	fmt.Println(oid)
	result := make(map[string]interface{})
	res := collection.FindOneAndDelete(context.Background(), bson.M{"_id": oid})
	err = res.Decode(&result)
	if err != nil {
		return "811", err
	}
	return "", nil
}
