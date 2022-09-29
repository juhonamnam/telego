package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/juhonamnam/telego/src/main/controller"
	"github.com/juhonamnam/telego/src/telego"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: Cannot read .env file")
		panic(err.Error())
	}
	app := telego.Default(os.Getenv("TELEGRAM_BOT_API_KEY"))
	app.SetUpdateHandler(controller.UpdateHandler)
	app.Start()
}
