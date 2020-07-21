package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"log"
	"strings"
	"time"

	"github.com/inawazalam/forum-microservices/api/mongodb"
	"go.mongodb.org/mongo-driver/bson"
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
func SaveCoupon(coupon Coupon) (Coupon, error) {
	var err error
	//con := controllers.InitializeMong(os.Getenv("MONGO_DB_DRIVER"), os.Getenv("MONGO_DB_NAME"))
	if client == nil {
		client = mongodb.InitializeMongo()
	}
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
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return coupon, err
}

//
func ValidateCouponCode(code string) (Coupon, error) {
	var err error //var filter bson.M
	var result Coupon
	var temp interface{}
	if client == nil {
		client = mongodb.InitializeMongo()
	}
	poo := map[string]string{}
	err = json.Unmarshal([]byte(code), &poo)
	if err != nil {
		temp = code
	} else {
		temp = poo
	}
	collection := client.Database("traceable").Collection("coupon")
	//injection := strings.ContainsAny(code, ",{$")
	/*injection := strings.ContainsAny(code, ",{$")
	if injection {
		filter := bson.M{"couponcode": bson.M{`$gt`: ""}}
		err = collection.FindOne(context.TODO(), filter).Decode(&result)
	} else {*/
	filters := bson.D{{"couponcode", temp}}
	err = collection.FindOne(context.TODO(), filters).Decode(&result)
	//}
	if err != nil {
		return result, err
	}
	return result, err
}
