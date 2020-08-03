package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Coupon sr
type Coupon struct {
	CouponCode string `json:"couponcode"`
	Amount     string `json:"amount"`
	CreatedAt  time.Time
}

//
func (c *Coupon) Prepare() {
	c.CouponCode = html.EscapeString(strings.TrimSpace(c.CouponCode))
	c.Amount = html.EscapeString(strings.TrimSpace(c.Amount))
	c.CreatedAt = time.Now()

}

//
func (c *Coupon) Validate() error {

	if c.CouponCode == "" {
		return errors.New("Required Coupon Code")
	}
	if c.Amount == "" {
		return errors.New("Required Coupon Amount")
	}

	return nil
}

//
func SaveCoupon(client *mongo.Client, coupon Coupon) (Coupon, error) {
	var err error
	//con := controllers.InitializeMong(os.Getenv("MONGO_DB_DRIVER"), os.Getenv("MONGO_DB_NAME"))

	//conn := mongodb.GetMongoConnection("localhost", "traceable")
	// Get a handle for your collection
	collection := client.Database("traceable").Collection("coupon")

	// Some dummy data to add to the Database
	// := Coupon{"Ruan", 32, "Cape Town"}ruan
	/*james := Coupon{"James", 32, "Nairobi"}
	frankie := Coupon{"Frankie", 31, "Nairobi"}*/

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), coupon)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return coupon, err
}

//
func ValidateCouponCode(client *mongo.Client, code string) (Coupon, error) {
	var err error //var filter bson.M
	var result Coupon
	var temp interface{}
	poo := map[string]string{}
	err = json.Unmarshal([]byte(code), &poo)
	if err != nil {
		temp = code
	} else {
		temp = poo
	}
	collection := client.Database("traceable").Collection("coupons")

	filters := bson.D{{"couponcode", temp}}
	err = collection.FindOne(context.TODO(), filters).Decode(&result)
	//}
	if err != nil {
		return result, err
	}
	return result, err
}
