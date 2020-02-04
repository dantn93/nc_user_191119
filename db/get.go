package db

import (
	"context"
	"time"

	"github.com/golang191119/nc_user/model"
	"github.com/golang191119/nc_user/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginUser(req model.LoginReq) (*model.LoginResp, error) {
	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"email": req.Email, "password": utils.MD5(req.Password)} //map[string]interface{}
	err := Client.Database(DbName).Collection(ColName).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	token := utils.GenerateToken(user.ID, user.Phone, user.Email)
	resp := model.LoginResp{&user, token}
	return &resp, err
}

func FindUserByID(ID int) (*model.User, error) {
	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": ID} //map[string]interface{}
	err := Client.Database(DbName).Collection(ColName).FindOne(ctx, filter).Decode(&user)
	return &user, err
}
