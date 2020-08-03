package seed

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/inawazalam/forum-microservices/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var coupons = []models.Coupon{
	models.Coupon{
		CouponCode: "TRAC075",
		Amount:     "75",
		CreatedAt:  time.Now(),
	},
	models.Coupon{
		CouponCode: "TRAC065",
		Amount:     "65",
		CreatedAt:  time.Now(),
	},
	models.Coupon{
		CouponCode: "TRAC125",
		Amount:     "125",
		CreatedAt:  time.Now(),
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

//
func LoadMongoData(mongoClient *mongo.Client) {
	var couponResult interface{}
	var postResult interface{}
	collection := mongoClient.Database(os.Getenv("MONGO_DB_NAME")).Collection("coupons")
	// get a MongoDB document using the FindOne() method
	err := collection.FindOne(context.TODO(), bson.D{}).Decode(&couponResult)
	if err != nil {
		for i, _ := range coupons {
			couponData, err := collection.InsertOne(context.TODO(), coupons[i])
			fmt.Println(couponData, err)
		}
	}

	er := collection.FindOne(context.TODO(), bson.D{}).Decode(&postResult)
	if er != nil {
		for j, _ := range posts {
			postData, err := collection.InsertOne(context.TODO(), posts[j])
			fmt.Println(postData, err)
		}
	}
}
