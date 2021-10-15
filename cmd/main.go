package main

import (
	"fmt"
	"git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/generated/openapi"
	_ "git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/generated/swaggerui"
	"git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/internal/core/middleware"
	"git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/internal/handler"
	"git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/pkg/properties"
	"github.com/go-chi/chi/v5"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
)

var (
	serviceProps            = properties.GetInstance()
	router       chi.Router = nil
)

func main() {
	initRouter()
	setupServer()
	setupSwaggerUi()
	startServer()
}

func initRouter() {
	router = chi.NewRouter()
}

func setupSwaggerUi() {
	swaggerUI, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	docsPath := serviceProps.ContextPath + "/docs/"
	router.Handle(docsPath+"*", http.StripPrefix(docsPath, http.FileServer(swaggerUI)))
}

func setupServer() {
	h := handler.NewHandler()

	router.Use(middleware.HttpLogger)
	router.Use(middleware.Recovery)

	openapi.HandlerFromMuxWithBaseURL(h, router, serviceProps.ContextPath)
}

func startServer() {
	log.Println("Starting service...")

	if router == nil {
		panic("Router not set")
	}

	addr := fmt.Sprintf("%s:%s", serviceProps.Host, serviceProps.Port)

	done := make(chan bool)
	go func() {
		err := http.ListenAndServe(addr, router)
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("Started service on http://%s%s", addr, serviceProps.ContextPath)
	<-done
}
