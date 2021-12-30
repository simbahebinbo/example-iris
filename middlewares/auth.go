package middlewares

import (
	"example-iris/session"
	"github.com/kataras/iris/v12"
	"net/http"
)

// CustomAuth is our custom authentication method
func CustomAuth(ctx iris.Context) {

	if auth, _ := session.Sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.WriteString("error: invalid session.")
		return
	}

	ctx.Next()
}
