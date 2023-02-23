package db

import (
	"context"
	"errors"
	"handson/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
	Age  int32              `bson:"age,omitempty"`
}

type UserDB interface {
	UpsertOne(ctx context.Context, user *User) error
}

type userDB struct {
	*mongo.Client
}

func NewUserDB(ctx context.Context, config *config.Config) (*userDB, error) {
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(config.Mongo.Conn).
			SetAuth(options.Credential{
				AuthSource: config.Mongo.Database,
				Username:   config.Mongo.User,
				Password:   config.Mongo.Password,
			}))
	if err != nil {
		return nil, errors.New("invalid mongodb options")
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, errors.New("cannot connect to mongodb instance")
	}
	return &userDB{client}, nil
}

func (u *userDB) UpsertOne(ctx context.Context, user *User) error {
	return nil
}
