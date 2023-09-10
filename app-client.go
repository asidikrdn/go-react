package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type ClientService struct {
	Path   string
	Mode   string
	Port   string
	server *http.Server
}

func (cs *ClientService) installReactDependency() {
	// install dependency
	install := exec.Command("npm", "install")
	install.Dir = cs.Path
	log.Println("Installing frontend dependencies...")
	install.Start()
	if err := install.Wait(); err != nil {
		log.Fatalf("Error Installing frontend dependencies: %v", err)
	} else if install.ProcessState.Success() {
		log.Println("Frontend dependencies successfully installed !")
	}
}

func (cs *ClientService) reactBuild() {
	// install dependency
	cs.installReactDependency()

	// build frontend
	build := exec.Command("npm", "run", "build")
	build.Dir = cs.Path
	log.Println("Building frontend...")
	// build.Stdout = os.Stdout
	// build.Stderr = os.Stderr
	build.Start()

	if err := build.Wait(); err != nil {
		log.Fatalf("Error building frontend: %v", err)
	} else if build.ProcessState.Success() {
		log.Println("Build frontend success !")
	}
}

func (cs *ClientService) RunDevelopment() {
	// install dependency
	cs.installReactDependency()

	// serving project
	serve := exec.Command("npm", "start")
	serve.Dir = cs.Path
	serve.Stdout = os.Stdout
	serve.Stderr = os.Stderr
	err := serve.Start()
	if err != nil {
		log.Fatalf("Error starting 'npm start': %v", err)
	}

	log.Println("Frontend running in development mode...")

	err = serve.Wait()
	if err != nil {
		log.Fatalf("'npm start' process exited with error: %v", err)
	}
}

func (cs *ClientService) RunProduction() {
	// build react
	cs.reactBuild()

	// gin Mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	// create new router
	router := gin.Default()

	// file server endpoint
	router.Static("/", cs.Path+"/build/")

	// set fallback route to redirect all request to react
	router.NoRoute(func(c *gin.Context) {
		c.File(cs.Path + "/build/")
	})

	// running services with gin
	// fmt.Println("CLIENT running on http://localhost:" + cs.Port)
	// router.Run(":" + cs.Port)

	// create server
	cs.server = &http.Server{
		Addr:    ":" + cs.Port,
		Handler: router,
	}

	// run server
	fmt.Println("CLIENT running on http://localhost" + cs.server.Addr)
	cs.server.ListenAndServe()
}

// shutdown metode untuk ClientService
func (cs *ClientService) Shutdown(ctx context.Context) {
	log.Println("Shutting down client service gracefully...")

	// start gracefully shutdown
	if err := cs.server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Client service gracefully shut down !")
}
