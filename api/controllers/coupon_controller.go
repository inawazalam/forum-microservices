package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/inawazalam/forum-microservices/api/models"
	"github.com/inawazalam/forum-microservices/api/responses"
)

//
func (s *Server) AddNewCoupon(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	coupon := models.Coupon{}
	err = json.Unmarshal(body, &coupon)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	coupon.Prepare()
	fmt.Println(coupon)
	savedCoupon, er := models.SaveCoupon(s.Client, coupon)
	if er != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	fmt.Println("Added Coupon: ", savedCoupon)

	responses.JSON(w, http.StatusOK, "Coupon Added in database")

}

//ValidateCoupon no sql injection
func (s *Server) ValidateCoupon(w http.ResponseWriter, r *http.Request) {
	var param string
	param = r.URL.Query().Get("coupon_code")
	if param == "" {
		//responses.ERROR(w, http.StatusBadRequest)
		responses.JSON(w, http.StatusBadRequest, "Invalid Param")
	}
	coupon, err := models.ValidateCouponCode(s.Client, param)
	if err != nil {
		responses.JSON(w, http.StatusExpectationFailed, "Sorry we didn't find any data")
	}
	responses.JSON(w, http.StatusOK, coupon)
}
