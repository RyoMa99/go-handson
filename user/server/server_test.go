package server

import (
	"context"
	"handson/user/userpb"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestServer(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		lis := bufconn.Listen(1024 * 1024)
		closer := serve(lis)
		defer closer()

		conn, err := grpc.Dial("localhost:50051", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return lis.Dial() }), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		client := userpb.NewUserServiceClient(conn)

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
}
