package data_test

import (
	"time"
	"user/internal/biz"
	"user/internal/data"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("User", func() {
	var ro biz.UserRepo
	var uD *biz.User

	ginkgo.BeforeEach(func() {
		ro = data.NewUserRepo(Db, nil)

		t := time.Now()
		uD = &biz.User{
			ID:       1,
			Mobile:   "13803881388",
			Password: "admin123456",
			NickName: "aliliin",
			Role:     1,
			Birthday: &t,
		}
	})

	ginkgo.It("CreateUser", func() {
		u, err := ro.CreateUser(ctx, uD)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(u.Mobile).Should(gomega.Equal("13803881388"))
	})

	ginkgo.It("ListUser", func() {
		user, total, err := ro.ListUser(ctx, 1, 10)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(user).ShouldNot(gomega.BeEmpty())
		gomega.Ω(total).Should(gomega.Equal(int64(1)))
		gomega.Ω(len(user)).Should(gomega.Equal(1))
		gomega.Ω(user[0].Mobile).Should(gomega.Equal("13803881388"))
	})

	ginkgo.It("UpdateUser", func() {
		birthday := time.Now()
		uD.NickName = "gyl"
		uD.Birthday = &birthday
		uD.Gender = "female"
		user, err := ro.UpdateUser(ctx, uD)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(user).Should(gomega.BeTrue())
	})

	ginkgo.It("CheckPassword", func() {
		p1 := "admin"
		encryptedPassword := "$pbkdf2-sha512$5p7doUNIS9I5mvhA$b18171ff58b04c02ed70ea4f39bda036029c107294bce83301a02fb53a1bcae0"
		password, err := ro.CheckPassword(ctx, p1, encryptedPassword)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(password).Should(gomega.BeTrue())

		encryptedPassword1 := "$pbkdf2-sha512$5p7doUNIS9I5mvhA$b18171ff58b04c02ed70ea4f39"
		password1, err := ro.CheckPassword(ctx, p1, encryptedPassword1)
		if err != nil {
			return
		}
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(password1).Should(gomega.BeFalse())
	})
})
