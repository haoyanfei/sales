package middleware

import "github.com/astaxie/beego/context"

type Auth struct {
	isPermission bool
}

func (a *Auth) Check() bool {
	return a.isPermission
}

func (a *Auth) Write(ctx *context.Context) {
	ctx.SetCookie("uid", "xxxxx")
}
