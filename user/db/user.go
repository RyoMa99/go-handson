package db

import (
	"context"
	"errors"
	"handson/config"

	"go.mongodb.org/mongo-driver/bson"
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
	FindOne(ctx context.Context, name string) (*User, error)
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
	return &userDB{
		collection: client.Database(config.Mongo.Database).Collection(config.Mongo.Collection),
	}, nil
}

type userDB struct {
	collection *mongo.Collection
}

func (u *userDB) UpsertOne(ctx context.Context, user *User) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": user.Id}
	update := bson.M{
		"$set": bson.M{
			"age":  user.Age,
			"name": user.Name,
		},
	}

	_, err := u.collection.UpdateOne(ctx, filter, update, opts)

	return err
}

func (u *userDB) FindOne(ctx context.Context, name string) (*User, error) {
	var user *User
	err := u.collection.FindOne(ctx, bson.M{"name": name}).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
