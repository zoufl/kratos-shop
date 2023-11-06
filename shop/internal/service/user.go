package service

import (
	"context"

	pb "shop/api/shop/v1"
	"shop/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
)

type ShopService struct {
	pb.UnimplementedShopServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewShopService(uc *biz.UserUsecase, logger log.Logger) *ShopService {
	return &ShopService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/shop")),
	}
}

func (s *ShopService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *ShopService) Login(ctx context.Context, req *pb.LoginReq) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *ShopService) Captcha(ctx context.Context, req *empty.Empty) (*pb.CaptchaReply, error) {
	return &pb.CaptchaReply{}, nil
}
func (s *ShopService) Detail(ctx context.Context, req *empty.Empty) (*pb.UserDetailResponse, error) {
	return &pb.UserDetailResponse{}, nil
}
