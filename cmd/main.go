package main

import (
	"ApiMarketplace/internal/config"
	"ApiMarketplace/internal/handlers/userhandler"
	"ApiMarketplace/internal/service/userservice"

	"ApiMarketplace/internal/store/db"
	"ApiMarketplace/internal/store/postgres"
	"fmt"
	"log"
	"net/http"
)

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Нулевой ендпоинт")
}

func main() {
	var err error
	err = config.InitConfig(".env")
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Не удалось открыть соединение с базой данных")
	}
	userRepo := postgres.NewUserPostgres(db)
	userService := userservice.NewUserService(userRepo)
	userRegisterHandler := userhandler.NewHandlerRegister(userService)

	userLoginHandler := userhandler.NewLoginHandler(userService)
	//
	http.HandleFunc("/", ServerHandler)
	http.HandleFunc("/register", userRegisterHandler.RegisterUserHandler)
	http.HandleFunc("/login", userLoginHandler.LoginUserHandler)
	fmt.Println("Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске")
	}
}
