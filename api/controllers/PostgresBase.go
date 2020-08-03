package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Server k
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
	Client *mongo.Client
}

//

//
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			fmt.Println("This is the error:", err)
		} else {
			fmt.Println("We are connected to the %s database", Dbdriver)
		}
	}
	/*if Dbdriver == "mongo" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
	}*/
	//server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration

	server.Router = mux.NewRouter()
	/*mux = http.NewServeMux()
	mux.HandleFunc("/plm/cors",Cors)
	http.ListenAndServe(":8081", mux)*/

	server.initializeRoutes()
	//corsOpts.Handler(server.Router)
}

//var mongoConnection *mogo.Connection = nil

//

//
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8087")
	fmt.Println(http.ListenAndServe(addr, server.Router))
}
