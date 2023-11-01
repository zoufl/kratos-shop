package biz_test

import (
	"github.com/golang/mock/gomock"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"user/internal/biz"
	"user/internal/mocks/mrepo"
)

var _ = ginkgo.Describe("UserUsecase", func() {
	var userCase *biz.UserUsecase
	var mUserRepo *mrepo.MockUserRepo

	ginkgo.BeforeEach(func() {
		mUserRepo = mrepo.NewMockUserRepo(ctl)
		userCase = biz.NewUserUsecase(mUserRepo, nil)
	})

	ginkgo.It("Create", func() {
		info := &biz.User{
			ID:       1,
			Mobile:   "13803881388",
			Password: "admin123456",
			NickName: "aliliin",
			Role:     1,
			Birthday: 693629981,
		}

		mUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(info, nil)
		l, err := userCase.Create(ctx, info)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(err).ToNot(gomega.HaveOccurred())
		gomega.Ω(l.ID).To(gomega.Equal(int64(1)))
		gomega.Ω(l.Mobile).To(gomega.Equal("13803881388"))
	})
})
