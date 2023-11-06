package biz

import (
	"context"
	"errors"
	"shop/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrPasswordInvalid     = errors.New("password invalid")
	ErrUsernameInvalid     = errors.New("username invalid")
	ErrCaptchaInvalid      = errors.New("verification code error")
	ErrMobileInvalid       = errors.New("mobile invalid")
	ErrUserNotFound        = errors.New("user not found")
	ErrLoginFailed         = errors.New("login failed")
	ErrGenerateTokenFailed = errors.New("generate token failed")
	ErrAuthFailed          = errors.New("authentication failed")
)

type User struct {
	ID        int64
	Mobile    string
	Password  string
	NickName  string
	Birthday  int64
	Gender    string
	Role      int
	CreatedAt time.Time
}

type UserRepo interface {
	CreateUser(c context.Context, u *User) (*User, error)
	UserByMobile(ctx context.Context, mobile string) (*User, error)
	UserById(ctx context.Context, Id int64) (*User, error)
	CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
}

type UserUsecase struct {
	uRepo      UserRepo
	log        *log.Helper
	signingKey string
}

func NewUserUsecase(repo UserRepo, logger log.Logger, conf *conf.Auth) *UserUsecase {
	helper := log.NewHelper(log.With(logger, "module", "usecase/shop"))
	return &UserUsecase{
		uRepo:      repo,
		log:        helper,
		signingKey: conf.JwtKey,
	}
}

func (u *UserUsecase) CreateUser(c context.Context, user *User) (*User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserUsecase) UserByMobile(ctx context.Context, mobile string) (*User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserUsecase) UserById(ctx context.Context, Id int64) (*User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserUsecase) CheckPassword(ctx context.Context, password string, encryptedPassword string) (bool, error) {
	panic("not implemented") // TODO: Implement
}
