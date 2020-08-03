package controllers

import (
	"github.com/inawazalam/forum-microservices/api/middlewares"
)

//
func (s *Server) initializeRoutes() {

	s.Router.Use(middlewares.AccessControlMiddleware)

	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/api/v2/coupon/validate-coupon", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.ValidateCoupon, s.DB))).Methods("GET", "OPTIONS")
	// Post Route
	s.Router.HandleFunc("/api/v2/community/posts/recent", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetPost, s.DB))).Methods("GET", "OPTIONS")

	s.Router.HandleFunc("/api/v2/community/posts/{postID}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetPostByID, s.DB))).Methods("GET", "OPTIONS")

	//s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUser))).Methods("GET")

	s.Router.HandleFunc("/api/v2/community/posts", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.AddNewPost, s.DB))).Methods("POST", "OPTIONS")

	s.Router.HandleFunc("/api/v2/commuinty/posts/{postID}/comment", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.Comment, s.DB))).Methods("POST", "OPTIONS")

	//	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	//Coupon Route
	s.Router.HandleFunc("/users/new-coupon", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.AddNewCoupon, s.DB))).Methods("POST", "OPTIONS")

}
