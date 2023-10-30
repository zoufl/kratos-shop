package test

import (
	"context"
	"fmt"
	"testing"
	v1 "user/api/user/v1"

	"google.golang.org/grpc"
)

var userClient v1.UserClient
var conn *grpc.ClientConn

// Init 初始化 grpc 链接 注意这里链接的 端口
func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		panic("grpc link err" + err.Error())
	}
	userClient = v1.NewUserClient(conn)
}

func TestCreateUser(t *testing.T) {
	Init()
	defer conn.Close()

	rsp, err := userClient.CreateUser(context.Background(), &v1.CreateUserInfo{
		Mobile:   fmt.Sprintf("1388888888%d", 1),
		Password: "admin123",
		NickName: fmt.Sprintf("YWWW%d", 1),
	})
	if err != nil {
		panic("grpc 创建用户失败" + err.Error())
	}
	t.Log(rsp.Id)
}
