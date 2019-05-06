package session

import (
	"github.com/kataras/iris/sessions"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	// Sess is our session
	Sess = sessions.New(sessions.Config{Cookie: cookieNameForSessionID, AllowReclaim: true})
)
