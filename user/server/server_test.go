package server

import (
	"context"
	"fmt"
	"handson/user/db"
	userpb "handson/user/proto"
	"log"
	"net"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type userDBSpy struct {
	users []*db.User
}

func (u *userDBSpy) UpsertOne(ctx context.Context, user *db.User) error {
	return nil
}

func (u *userDBSpy) FindOne(ctx context.Context, name string) (*db.User, error) {
	for _, user := range u.users {
		if user.Name == name {
			return user, nil
		}
	}
	return nil, fmt.Errorf("%s isn't registered.", name)
}

var client userpb.UserServiceClient

var user = &db.User{
	Id:   primitive.NewObjectID(),
	Name: "jiro",
	Age:  21,
}

func TestMain(m *testing.M) {
	lis := bufconn.Listen(1024 * 1024)
	closer := serve(lis, &userDBSpy{
		users: []*db.User{user},
	})
	defer closer()

	conn, err := grpc.Dial("localhost:50051", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return lis.Dial() }), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close()

	client = userpb.NewUserServiceClient(conn)

	code := m.Run()

	os.Exit(code)
}

func TestServer(t *testing.T) {
	t.Run("find user", func(t *testing.T) {
		res, err := client.FindUser(context.Background(), &userpb.FindUSerRequest{
			Name: user.Name,
		})

		if err != nil {
			t.Errorf("ERROR: %s", err)
		}

		if res.Name != user.Name {
			t.Errorf("want taro,but get %s", res.Name)
		}
	})

	t.Run("create user", func(t *testing.T) {
		_, err := client.CreateUser(context.Background(), &userpb.CreateUserRequest{
			Name: "taro",
			Age:  23,
		})

		if err != nil {
			t.Errorf("ERROR: %s", err)
		}
	})
}
