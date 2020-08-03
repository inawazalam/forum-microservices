package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inawazalam/forum-microservices/api/models"
	"github.com/inawazalam/forum-microservices/api/responses"
)

//
func (s *Server) AddNewPost(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	post := models.Post{}

	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//post.Prepare()
	fmt.Println(post)
	savePost, er := models.SavePost(s.Client, post)
	if er != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	fmt.Println("Added Coupon: ", savePost)

	responses.JSON(w, http.StatusOK, "Post Added in database")

}

//
func (s *Server) GetPostByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	//var autherID uint64
	GetPost, er := models.GetPostByID(s.Client, vars["postID"])
	if er != nil {
		responses.ERROR(w, http.StatusExpectationFailed, er)
	}

	responses.JSON(w, http.StatusOK, GetPost)

}

//GetPost Vulnerabilities
func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	//post := models.Post{}

	posts, err := models.FindAllPost(s.Client)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

//
func (s *Server) Comment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	comment := models.CommentRequest{}

	err = json.Unmarshal(body, &comment)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if vars["postID"] == "" {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	comment.ID = vars["postID"]
	postData, er := models.CommentOnPost(s.Client, comment)
	if er != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, er)
		return
	}
	responses.JSON(w, http.StatusOK, postData)
}
