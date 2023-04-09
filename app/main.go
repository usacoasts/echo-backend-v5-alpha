package main

import (
	"app/conf"
	"app/interactor"
	"app/presenter/http/middleware"
	"app/presenter/http/router"
	"flag"
	"fmt"
	"github.com/labstack/echo/v5"
	"log"
)

func main() {
	flag.Parse()
	conf.NewConfig()

	// Echo instance
	e := echo.New()
	db := conf.NewDBConnection()

	mysqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mysqlDB.Close(); err != nil {
			log.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()
	i := interactor.NewInteractor(db)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)
	if err := e.Start(fmt.Sprintf(":%d", conf.Current.Server.Port)); err != nil {
		log.Fatal(fmt.Sprintf("Failed to start: %v", err))
	}
}
