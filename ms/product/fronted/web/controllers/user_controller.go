package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
	"github.com/linzhenlong/my-go-code/ms/product/services"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// UserController .
type UserController struct {
	Ctx         iris.Context
	UserService services.IUserService
	Session     *sessions.Session
}

// GetRegister 用户注册.
func (u *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

// GetLogin .
func (u *UserController) GetLogin() mvc.View {
	redirectURL := u.Ctx.URLParam("url")
	return mvc.View{
		Name: "user/login.html",
		Data: iris.Map{
			"redirectURL": redirectURL,
		},
	}
}

// PostRegister .
func (u *UserController) PostRegister() {
	var (
		nickName = u.Ctx.FormValue("nickName")
		userName = u.Ctx.FormValue("userName")
		password = u.Ctx.FormValue("password")
	)
	log.Printf("%v\n", nickName)
	log.Printf("%v\n", len(nickName))
	log.Printf("%v\n", len(strings.TrimSpace(nickName)))
	u.Ctx.Application().Logger().Debug(len(strings.TrimSpace(nickName)))
	// ozzo-validation
	if len(strings.TrimSpace(nickName)) == 0 {
		u.Ctx.HTML(errHTML("nickname不能空", "/user/register"))
		return
	}
	if len(strings.TrimSpace(userName)) == 0 {
		u.Ctx.HTML(errHTML("username 不能为空", "/user/register"))
		return
	}
	if len(strings.TrimSpace(password)) == 0 {
		u.Ctx.HTML(errHTML("pasword 不能为空", "/user/register"))
		return
	}

	// 判断用户名是否被占用.
	existsUser, err := u.UserService.GetUserByName(userName)
	u.Ctx.Application().Logger().Debug(err)
	if err != nil {
		u.Ctx.HTML(errHTML(err.Error(), "/user/register"))
		return
	}
	if len(strings.TrimSpace(existsUser.UserName)) > 0 {
		u.Ctx.HTML(errHTML(userName+"已存在！！", "/user/register"))
		return
	}

	user := datamodels.User{
		UserName: userName,
		NickName: nickName,
		Password: password,
	}

	_, err = u.UserService.AddUser(&user)
	if err != nil {
		u.Ctx.StatusCode(iris.StatusInternalServerError)
	}
	u.Ctx.HTML(errHTML("注册成功", "/user/login"))
	return // Redirect 后面需要一个return
}

// PostLogin .
func (u *UserController) PostLogin() {
	var (
		userName = u.Ctx.FormValue("userName")
		password = u.Ctx.FormValue("password")
		url      = u.Ctx.FormValue("url")
	)
	user, isOK := u.UserService.IsPwdSuccess(userName, password)
	if !isOK {
		u.Ctx.HTML(errHTML("用户名或密码错误", "/user/login"))
		return
	}
	u.Ctx.SetCookie(&http.Cookie{
		Name:  "uid",
		Value: strconv.FormatInt(user.ID, 10),
		Path:  "/",
	})
	u.Session.Set("uid", strconv.FormatInt(user.ID, 10))
	u.Ctx.Application().Logger().Debug("url===>", url)
	if url != "" {
		u.Ctx.Redirect(url)
	}
	return
}

// errHtml.
func errHTML(msg, href string) string {
	errhtml := "<script>alert('" + msg + "');location.href='" + href + "';</script>"
	return errhtml
}
