package controller

import "github.com/juhonamnam/telego/src/telego"

func UpdateHandler(ctx *telego.UpdateContext) {
	if ctx.Update.Message != nil {
		ctx.SendMessage(ctx.Update.Message.From.Id, ctx.Update.Message.From.FirstName)
	}
}
