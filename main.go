package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"taskManagement/mongoClient"
	"taskManagement/server"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	err := mongoClient.SetUpMongo()
	if err != nil {
		panic("Failed to initialize database")
	}
	log.Print("database connection successful")

	srv := server.SetUpRoutes()
	go func(srv *server.Server) {
		if err := srv.Run(); err != nil {
			log.Fatal("unable to start http server")
		}
	}(srv)
	log.Print("server connection successful")

	<-done

	log.Print("shutting down server")
	dbCloseErr := mongoClient.ShutdownMongo()
	if dbCloseErr != nil {
		log.Fatal("failed to gracefully close db")
		return
	}

}
