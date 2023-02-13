package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"lms/routes"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	r := routes.Router()
	fmt.Println("Server dijalankan pada port" + os.Getenv("PORT"))
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	origins := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED"), os.Getenv("ORIGIN_ALLOWED2")})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(origins, headers, methods)(r)))
}
