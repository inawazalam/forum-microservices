package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
	savePost, er := models.SavePost(post)
	if er != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	fmt.Println("Added Coupon: ", savePost)

	responses.JSON(w, http.StatusOK, "Post Added in database")

}

//
func (s *Server) GetPostByID(w http.ResponseWriter, r *http.Request) {

	auther := r.URL.Query().Get("autherid")
	if auther == "" {
		//responses.ERROR(w, http.StatusBadRequest)
		responses.JSON(w, http.StatusBadRequest, "Invalid Param")
	}
	autherID, err := strconv.ParseUint(auther, 0, 16)
	if err != nil {
		responses.ERROR(w, http.StatusExpectationFailed, err)
	}
	//var autherID uint64
	GetPost, er := models.GetPostByID(autherID)
	if er != nil {
		responses.ERROR(w, http.StatusExpectationFailed, er)
	}

	responses.JSON(w, http.StatusOK, GetPost)

}

//GetPost Vulnerabilities
func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	//post := models.Post{}

	posts, err := models.FindAllPost()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

//
func (s *Server) Comment(w http.ResponseWriter, r *http.Request) {

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
	//comment, er := models.CommentOnPost(post)
}
