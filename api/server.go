package api

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/inawazalam/forum-microservices/api/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

//
func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
}

//
func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values", os.Getenv("DB_DRIVER"))
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	//var mongoConnection *mogo.Connection = server.InitializeMongo(os.Getenv("MONGO_DB_DRIVER"), os.Getenv("MONGO_DB_USER"), os.Getenv("MONGO_DB_PASSWORD"), os.Getenv("MONGO_DB_PORT"), os.Getenv("MONGO_DB_HOST"), os.Getenv("MONGO_DB_NAME"))
	//server.InitializeMongo(os.Getenv("MONGO_DB_DRIVER"), os.Getenv("MONGO_DB_USER"), os.Getenv("MONGO_DB_PASSWORD"), os.Getenv("MONGO_DB_PORT"), os.Getenv("MONGO_DB_HOST"), os.Getenv("MONGO_DB_NAME"))

	//seed.Load(server.DB)
	//seed.LoadMongo(mongoConnection)
	//server.InitializeMong(os.Getenv("MONGO_DB_DRIVER"), os.Getenv("MONGO_DB_NAME"))

	server.Run(":8088")

}
