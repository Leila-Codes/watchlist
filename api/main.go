package main

import (
	"flag"
	"fmt"
	"net/http"

	"lunasphere.co.uk/watchlist/api/controllers"
	"lunasphere.co.uk/watchlist/api/services"
)

func main() {
	cfgPathPtr := flag.String("config", "./config.json", "Path to JSON configuration file.")
	portPtr := flag.Int("port", 8080, "Port to run HTTP server on.")

	flag.Parse()

	err := services.Connect(*cfgPathPtr)
	if err != nil {
		panic(err)
	}

	Server(*portPtr)
}

func Server(listenPort int) {
	controllers.TitlesAPI()

	err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
	if err != nil {
		panic(err)
	}
}
