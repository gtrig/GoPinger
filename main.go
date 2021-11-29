package main

import (
	"fmt"
	"gopinger/lib/config"
	"gopinger/lib/pinger"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	config.Init()
	//Create a new instance of the API
	api := rest.NewApi()

	//Create a new router
	router, err := rest.MakeRouter(
		rest.Get("/ping/#ip", Ping),
		rest.Get("/pingrange/#start/#end", PingRange),
	)
	if err != nil {
		log.Fatal(err)
	}
	//Add the router to the API
	api.SetApp(router)
	printHelper()
	//Start the API
	log.Fatal(http.ListenAndServe(":"+config.GetString("GOPINGER_API_PORT"), api.MakeHandler()))
}

func printHelper() {
	fmt.Printf("Server running on port %s\n", config.GetString("GOPINGER_API_PORT"))
	fmt.Printf("There are two routes:\n")
	fmt.Printf("/ping/<ip>\n")
	fmt.Printf("/pingrange/<start>/<end>\n")
	fmt.Printf("Examples:\n")
	fmt.Printf("http://localhost:%s/ping/192.168.1.1\n", config.GetString("GOPINGER_API_PORT"))
	fmt.Printf("http://localhost:%s/pingrange/192.168.1.1/192.168.1.254\n", config.GetString("GOPINGER_API_PORT"))
}

func Ping(w rest.ResponseWriter, r *rest.Request) {
	//Get the IP from the request
	ip := r.PathParam("ip")
	log.Println(ip)
	//Create a new instance of the ping class
	result := pinger.ScanSingle(ip)
	//Return the result
	w.WriteJson(result)
}

func PingRange(w rest.ResponseWriter, r *rest.Request) {
	//Get the IP from the request
	start := r.PathParam("start")
	end := r.PathParam("end")
	//Create a new instance of the ping class
	result := pinger.ScanRange(start, end)

	//Return the result
	w.WriteJson(result)
}
