package main

import (
	_ "github.com/nikolajbotalov/effective_mobile_test/docs"
	"github.com/nikolajbotalov/effective_mobile_test/internal/app"
)

// @title Subscriptions Service
// @version 1.0
// @description Сервис для агрегации данных об онлайн-подписках пользователей
// @host localhost:8080
// @BasePath /api/v1
func main() {
	application, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	application.Server.Run()
}
