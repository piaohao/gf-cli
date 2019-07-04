package ctl_web

import (
	"github.com/dgrijalva/jwt-go"
	svcAdmin "github.com/gogf/gf-cli/app/service/admin"
	"github.com/gogf/gf-cli/util"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/util/gconv"
	"time"
)

type MyController struct {
	gmvc.Controller
}

func (c *MyController) Login() {
	util.Html(c.Request, "/web/login.html")
}

func (c *MyController) Auth() {
	r := c.Request
	username := r.GetPostString("username")
	password := r.GetPostString("password")
	user := svcAdmin.SysUserService.GetByLoginInfo(username, password)
	if user == nil {
		util.WriteErrorByDefaultCode(c.Request, "用户名密码错误")
	} else {
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(12)).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   gconv.String(user["id"]),
		}
		tokenStr, err := token.SignedString([]byte("gf-cli"))
		if err != nil {
			util.WriteDefaultError(c.Request)
		}
		util.WriteSuccess(c.Request, g.Map{"token": tokenStr})
	}
}

func (c *MyController) GetUserId() {
	proxyId := c.Request.GetParam("userId").Int()
	user := svcAdmin.SysUserService.Get(proxyId)
	if user == nil {
		util.WriteErrorByDefaultCode(c.Request, "用户不存在")
	}
	util.WriteSuccess(c.Request, proxyId)
}
