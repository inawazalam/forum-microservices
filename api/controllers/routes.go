package controllers

import (
	"github.com/inawazalam/forum-microservices/api/middlewares"
)

//
func (s *Server) initializeRoutes() {
	s.Router.Use(middlewares.AccessControlMiddleware)

	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/users/validate-coupon", middlewares.SetMiddlewareJSON(s.ValidateCoupon)).Methods("GET")
	s.Router.HandleFunc("/users/get-post", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetPostByID))).Methods("GET")

	// Post Route
	s.Router.HandleFunc("/users/get-recent-post", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetPost))).Methods("GET")

	//s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUser))).Methods("GET")

	s.Router.HandleFunc("/users/add-new-post", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.AddNewPost))).Methods("POST")

	s.Router.HandleFunc("/users/comment", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.Comment))).Methods("POST")

	//	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	//Coupon Route
	s.Router.HandleFunc("/users/new-coupon", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.AddNewCoupon))).Methods("POST")

	//

}
