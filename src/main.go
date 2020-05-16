package main

import (
	"fmt"
	"log"
	"os"

	"controllers"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r := router.New()
	r.GET("/domain", controllers.GetDomainsQueries)
	r.GET("/domain/{name}", controllers.GetDomainParameters)
	fmt.Println("Your application is running here:  http://localhost" + PORT)
	log.Fatal(fasthttp.ListenAndServe(PORT, r.Handler))
}
