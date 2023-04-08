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
	conn := conf.NewDBConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()
	i := interactor.NewInteractor(conn)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)
	if err := e.Start(fmt.Sprintf(":%d", conf.Current.Server.Port)); err != nil {
		log.Fatal(fmt.Sprintf("Failed to start: %v", err))
	}
}
