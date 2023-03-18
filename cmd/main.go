package main

import(
	"log"
	"net/url"
	"net/http"

	"cinemaxxi_layout/infrastructure"
	"cinemaxxi_layout/gateway"

	"github.com/go-chi/chi"
)

func main() {
	routes := chi.NewRouter()

	r := gateway.InjectRoutes()
	r.RegisterRoutes(routes)

	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", routes))
}

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := infrastructure.Environment{}
	env.SetEnvironment()
	env.InitCinemaConfigInMemory()
}