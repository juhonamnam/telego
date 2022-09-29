package controller

import (
	"github.com/juhonamnam/telego/src/main/controller/message"
	"github.com/juhonamnam/telego/src/telego"
)

func UpdateHandler(ctx *telego.Context) {
	if ctx.Update.Message != nil {
		message.HandleMessage(ctx)
	}
}
