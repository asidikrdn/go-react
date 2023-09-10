package main

import (
	"context"
	"fmt"
	"go-react/server/routes"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ServerService struct {
	Path   string
	Port   string
	server *http.Server
}

func (ss *ServerService) RunServer(ginMode string) {

	// gin Mode
	gin.SetMode(ginMode)

	// create new router
	router := gin.Default()

	// call logger middleware before route to any routes
	// router.Use(middleware.Logger())

	// set up CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Replace with your allowed origins
	config.AllowMethods = []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Authorization"}

	// add cors middleware on all route
	router.Use(cors.New(config))

	// call routerinit with pathprefix
	routes.RouterInit(router.Group("/api/v1"))

	// file server endpoint
	router.Static("/static", ss.Path+"/uploads")

	// running services with gin
	// fmt.Println("SERVER running on http://localhost:" + ss.Port)
	// router.Run(":" + ss.Port)

	// create server
	ss.server = &http.Server{
		Addr:    ":" + ss.Port,
		Handler: router,
	}

	// run server
	fmt.Println("SERVER running on http://localhost" + ss.server.Addr)
	ss.server.ListenAndServe()
}

// shutdown metode untuk ServerService
func (ss *ServerService) Shutdown(ctx context.Context) {
	log.Println("Shutting down server service gracefully...")

	// start gracefully shutdown
	if err := ss.server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server service gracefully shut down !")
}
