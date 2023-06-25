package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/passport/cmd/passport/ddd/register"
	"github.com/lishimeng/passport/cmd/passport/ddd/signin"
	"github.com/lishimeng/passport/cmd/passport/ddd/signup"
	"github.com/lishimeng/passport/cmd/passport/ddd/user"
)

func Route(root *iris.Application) {
	p := root.Party("/api")
	signin.Route(p.Party("/sign_in"))
	signup.Route(p.Party("/sign_up"))
	register.Route(p.Party("/register"))
	user.Route(p.Party("/user"))
}
