package main

import (
	"fmt"
	"gopinger/lib/pinger"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
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
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func printHelper() {
	fmt.Println("Server running on port 8080")
	fmt.Println("There are two routes:")
	fmt.Println("/ping/<ip>")
	fmt.Println("/pingrange/<start>/<end>")
	fmt.Println("Examples:")
	fmt.Println("http://localhost:8080/ping/192.168.1.1")
	fmt.Println("http://localhost:8080/pingrange/192.168.1.1/192.168.1.254")
}

func Ping(w rest.ResponseWriter, r *rest.Request) {
	//Get the IP from the request
	ip := r.PathParam("ip")
	log.Println(ip)
	//Create a new instance of the ping class
	result := pinger.ScanSingle(ip)
	//Return the result
	log.Println(result)
	w.WriteJson(result)
}

func PingRange(w rest.ResponseWriter, r *rest.Request) {
	//Get the IP from the request
	start := r.PathParam("start")
	end := r.PathParam("end")
	//Create a new instance of the ping class
	result := pinger.ScanRange(start, end)

	//Return the result
	fmt.Println(result)
	w.WriteJson(result)
}
