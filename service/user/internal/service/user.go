package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"

	pb "user/api/user/v1"
	"user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserInfo) (*pb.UserInfoResponse, error) {
	user, err := s.uc.Create(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		NickName: req.NickName,
	})
	if err != nil {
		return nil, err
	}

	userInfoResp := pb.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Birthday: uint64(user.Birthday.Unix()),
	}

	return &userInfoResp, nil
}

func (s *UserService) GetUserList(ctx context.Context, req *pb.PageInfo) (*pb.UserListResponse, error) {
	list, total, err := s.uc.List(ctx, int(req.Pn), int(req.PSize))
	if err != nil {
		return nil, err
	}

	resp := &pb.UserListResponse{}
	resp.Total = int32(total)

	for _, user := range list {
		userInfo := UserResponse(user)

		if user.Birthday != nil {
			userInfo.Birthday = uint64(user.Birthday.Unix())
		}

		resp.Data = append(resp.Data, &userInfo)
	}

	return resp, nil
}

func UserResponse(user *biz.User) pb.UserInfoResponse {
	userInfoRsp := pb.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoRsp.Birthday = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}

func (s *UserService) GetUserByMobile(ctx context.Context, req *pb.MobileRequest) (*pb.UserInfoResponse, error) {
	user, err := s.uc.UserByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	resp := UserResponse(user)

	return &resp, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserInfo) (*emptypb.Empty, error) {
	birthDay := time.Unix(int64(req.Birthday), 0)
	user, err := s.uc.UpdateUser(ctx, &biz.User{
		ID:       req.Id,
		Gender:   req.Gender,
		Birthday: &birthDay,
		NickName: req.NickName,
	})
	if err != nil {
		return nil, err
	}

	if user == false {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *UserService) CheckPassword(ctx context.Context, req *pb.PasswordCheckInfo) (*pb.CheckResponse, error) {
	check, err := s.uc.CheckPassword(ctx, req.Password, req.EncryptedPassword)
	if err != nil {
		return nil, err
	}

	return &pb.CheckResponse{Success: check}, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.IdRequest) (*pb.UserInfoResponse, error) {
	user, err := s.uc.UserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp := UserResponse(user)

	return &resp, nil
}
