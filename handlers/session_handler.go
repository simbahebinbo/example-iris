package handlers

import (
	"example-iris/session"
	"github.com/kataras/iris/v12"
	"net/http"
)

// Login is where session is created
func Login(ctx iris.Context) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if (username == "user1" && password == "pass1") || (username == "user2" && password == "pass2") {
		session := session.Sess.Start(ctx)
		session.Set("authenticated", true)

		ctx.WriteString("session created....")
		ctx.StatusCode(http.StatusOK)
		return
	}

	ctx.StatusCode(http.StatusUnauthorized)
	ctx.Values().Set("error", "invalid username/password.")
}

// Logout is where session is deleted
func Logout(ctx iris.Context) {
	session := session.Sess.Start(ctx)
	session.Set("authenticated", false)

	ctx.StatusCode(http.StatusOK)
	ctx.WriteString("session is successfully deleted.")
}
