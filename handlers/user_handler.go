package handlers

import (
	"github.com/kataras/iris"
)

// GetSessionID is the handler function for the get method
func GetSessionID(ctx iris.Context) {
	ctx.Writef("Hello world")
}
