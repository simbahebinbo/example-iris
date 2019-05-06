package middlewares

import (
	"github.com/arjunajithtp/example-iris/session"
	"net/http"

	"github.com/kataras/iris"
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
