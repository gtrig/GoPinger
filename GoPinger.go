// //An application that will send a ping an ip or a range of ips, and expose the results via a REST API

// package gopinger

// import (
// 	"log"
// 	"net/http"

// 	"github.com/ant0ine/go-json-rest/rest"
// )

// func main() {
// 	//Create a new instance of the API
// 	api := rest.NewApi()
// 	//Create a new router
// 	router, err := rest.MakeRouter(
// 		rest.Get("/ping/:ip", Ping),
// 		rest.Get("/ping/:ip/:count", PingCount),
// 		rest.Get("/ping/:ip/:count/:timeout", PingCountTimeout),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//Add the router to the API
// 	api.SetApp(router)
// 	//Start the API
// 	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
// }

// func Ping(w rest.ResponseWriter, r *rest.Request) {
// 	//Get the IP from the request
// 	ip := r.PathParam("ip")
// 	//Create a new instance of the ping class
// 	p := ping.New()
// 	//Ping the IP
// 	result := p.Ping(ip)
// 	//Return the result
// 	w.WriteJson(result)
// }

// func PingCount(w rest.ResponseWriter, r *rest.Request) {
// 	//Get the IP from the request
// 	ip := r.PathParam("ip")
// 	//Get the count from the request
// 	count := r.PathParam("count")
// 	//Create a new instance of the ping class
// 	p := ping.New()
// 	//Ping the IP
// 	result := p.PingCount(ip, count)
// 	//Return the result
// 	w.WriteJson(result)
// }

// func PingCountTimeout(w rest.ResponseWriter, r *rest.Request) {
// 	//Get the IP from the request
// 	ip := r.PathParam("ip")
// 	//Get the count from the request
// 	count := r.PathParam("count")
// 	//Get the timeout from the request
// 	timeout := r.PathParam("timeout")
// 	//Create a new instance of the ping class
// 	p := ping.New()
// 	//Ping the IP
// 	result := p.PingCountTimeout(ip, count, timeout)
// 	//Return the result
// 	w.WriteJson(result)
// }
