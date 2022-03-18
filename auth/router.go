package auth

import (
	"main/common/df"
	"main/common/kv"
	"main/common/mw"
	"main/common/orm"
	"main/common/rt"
	"main/common/tl"
	"net/http"
	"net/url"
)

func signIn(r *http.Request, msg *mw.Message, extra mw.Extra, next mw.Next) {
	bodyJson := (extra["body"]).(kv.Element)
	acc := bodyJson.GetProperty("acc").GetString()
	pwd := bodyJson.GetProperty("pwd").GetString()

	if acc == "" || pwd == "" {
		msg.Code = df.HTTP_STATUS_PARAM_ERROR
		msg.Msg = "用户名和密码都不能为空"
		return
	}

	orm := orm.NewOrm("tb_admin")
	where := kv.NewObject()
	where.Set("acc", kv.NewString(acc))
	where.Set("pwd", kv.NewString(tl.ToSha512String(pwd)))
	item, err := orm.First(where, nil)
	if err != nil {
		msg.Code = df.HTTP_STATUS_SERVER_ERROR
		msg.Msg = err.Error()
		return
	}
	if item == nil {
		msg.Code = df.HTTP_STATUS_SERVER_ERROR
		msg.Msg = "用户名或密码错误"
		return
	}

	token := url.QueryEscape(tl.Encrypt(tl.Encrypt(acc)))
	u := newUser(item)
	u.Token = token
	users[acc] = u
	msg.Msg = `{"token":"` + token + `"}`
	msg.Code = 200
}

func signOut(r *http.Request, msg *mw.Message, extra mw.Extra, next mw.Next) {

}

func check(r *http.Request, msg *mw.Message, extra mw.Extra, next mw.Next) {
	msg.Code = 200
	msg.Msg = `已登录`
}

func init() {
	rt.RegisterRouter("/api/auth/check", rt.NewFactory(func() rt.Router {
		rt := rt.NewRouter()
		rt.UseMiddleware(authMiddleware)
		rt.UseMiddleware(mw.NewMiddleware(check))
		return rt
	}))

	rt.RegisterRouter("/api/auth/signin", rt.NewFactory(func() rt.Router {
		rt := rt.NewRouter()
		rt.UseMiddleware(mw.JsonBodyMiddleware)
		rt.UseMiddleware(mw.NewMiddleware(signIn))
		return rt
	}))
	rt.RegisterRouter("/api/auth/signout", rt.NewFactory(func() rt.Router {
		rt := rt.NewRouter()
		rt.UseMiddleware(authMiddleware)
		rt.UseMiddleware(mw.NewMiddleware(signOut))
		return rt
	}))
}
