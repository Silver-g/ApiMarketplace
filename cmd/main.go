package main

import (
	"ApiMarketplace/internal/config"
	"ApiMarketplace/internal/handlers"
	"ApiMarketplace/internal/handlers/adshandler"
	"ApiMarketplace/internal/handlers/userhandler"
	"ApiMarketplace/internal/service/adsservice"
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
	adsRepo := postgres.NewAdsPostgres(db)
	adsService := adsservice.NewAdsService(adsRepo)
	getAdsListHandlet := adshandler.NewHandlerGetAdsList(adsService)
	adsCreateHandler := adshandler.NewHandlerCreateAds(adsService)
	userService := userservice.NewUserService(userRepo)
	userRegisterHandler := userhandler.NewHandlerRegister(userService)

	userLoginHandler := userhandler.NewLoginHandler(userService)

	var router handlers.RouteInfo
	router.CreateAdsHandler = adsCreateHandler
	router.GetAdsListHandler = getAdsListHandlet
	handlerRouter := &router
	//
	http.HandleFunc("/", ServerHandler)
	http.HandleFunc("/register", userRegisterHandler.RegisterUserHandler)
	http.HandleFunc("/login", userLoginHandler.LoginUserHandler)
	http.Handle("/ads", handlerRouter)

	fmt.Println("Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске")
	}
}
