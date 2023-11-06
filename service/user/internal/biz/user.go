package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID        int64
	Mobile    string
	Password  string
	NickName  string
	Birthday  *time.Time
	Gender    string
	Role      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

//go:generate mockgen -destination=../mocks/mrepo/user.go -package=mrepo . UserRepo
type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	ListUser(context.Context, int, int) ([]*User, int64, error)
	UserByMobile(context.Context, string) (*User, error)
	GetUserById(context.Context, int64) (*User, error)
	UpdateUser(context.Context, *User) (bool, error)
	CheckPassword(context.Context, string, string) (bool, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUsecase) List(ctx context.Context, page, pageSize int) ([]*User, int64, error) {
	return uc.repo.ListUser(ctx, page, pageSize)
}

func (uc *UserUsecase) UserByMobile(ctx context.Context, mobile string) (*User, error) {
	return uc.repo.UserByMobile(ctx, mobile)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, u *User) (bool, error) {
	return uc.repo.UpdateUser(ctx, u)
}

func (uc *UserUsecase) CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error) {
	return uc.repo.CheckPassword(ctx, password, encryptedPassword)
}

func (uc *UserUsecase) UserById(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUserById(ctx, id)
}
