package server

import (
	"context"
	"handson/user/db"
	userpb "handson/user/proto"
	"log"
	"net"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type userDBSpy struct {
}

func (u *userDBSpy) UpsertOne(ctx context.Context, user *db.User) error {
	return nil
}

var client userpb.UserServiceClient

func TestMain(m *testing.M) {
	lis := bufconn.Listen(1024 * 1024)
	closer := serve(lis, &userDBSpy{})
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
	t.Run("Get", func(t *testing.T) {

		res, err := client.Get(context.Background(), &userpb.GetRequest{
			Id: "1",
		})
		if err != nil {
			t.Errorf("ERROR: %s", err)
		}
		if res.Name != "taro" {
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
