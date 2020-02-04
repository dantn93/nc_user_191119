package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/golang191119/nc_user/model"
	"github.com/golang191119/nc_user/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetNextSequenceValue(sequenceName string) (id int, err error) {
	collection := Client.Database(DbName).Collection("counters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	doc := make(map[string]interface{})
	collection.FindOneAndUpdate(ctx,
		bson.M{"id": sequenceName},
		bson.M{"$inc": bson.M{"sequence_value": 1}},
	).Decode(&doc)

	if doc["sequence_value"] == nil {
		_, err := collection.InsertOne(ctx, bson.M{"id": sequenceName, "sequence_value": 0})
		if err != nil {
			return 0, err
		}
		return 0, nil
	}

	counters := model.Counters{}
	byteCounter, err := json.Marshal(doc)
	if err != nil {
		return 0, err
	}
	json.Unmarshal(byteCounter, &counters)
	return counters.SequenceValue + 1, nil
}

func AddUser(user *model.User) (interface{}, error) {
	user.Password = utils.MD5(user.Password)
	collection := Client.Database(DbName).Collection(ColName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id, err := GetNextSequenceValue("user_id")
	if err != nil {
		return nil, err
	}
	user.ID = id
	res, err := collection.InsertOne(ctx, user)
	return res, err
}

func UpdateUser(req *model.User) (interface{}, error) {
	user, err := FindUserByID(req.ID)
	if err != nil {
		return nil, err
	}
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Password != "" {
		user.Password = utils.MD5(req.Password)
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	collection := Client.Database(DbName).Collection(ColName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": user.ID}
	update := bson.M{"$set": user}
	res, err := collection.UpdateOne(ctx, filter, update)
	return res, err
}
