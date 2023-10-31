package data_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"user/internal/biz"
	"user/internal/data"
)

var _ = ginkgo.Describe("User", func() {
	var ro biz.UserRepo
	var uD *biz.User

	ginkgo.BeforeEach(func() {
		ro = data.NewUserRepo(Db, nil)

		uD = &biz.User{
			ID:       1,
			Mobile:   "13803881388",
			Password: "admin123456",
			NickName: "aliliin",
			Role:     1,
			Birthday: 693629981,
		}
	})

	ginkgo.It("CreateUser", func() {
		u, err := ro.CreateUser(ctx, uD)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(u.Mobile).Should(gomega.Equal("13803881388"))
	})
})
