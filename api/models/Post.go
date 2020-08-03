package models

import (
	"context"
	"errors"
	"fmt"
	"html"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/lithammer/shortuuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//
type Post struct {
	ID        string     `gorm:"primary_key;auto_increment" json:"id"`
	Title     string     `gorm:"size:255;not null;unique" json:"title"`
	Content   string     `gorm:"size:255;not null;" json:"content"`
	Author    User       `json:"author"`
	Comments  []Comments `json:"comments"`
	AuthorID  uint64     `sql:"type:int REFERENCES users(id)" json:"authorid"`
	CreatedAt time.Time
}

//
func (p *Post) Prepare() {
	p.ID = shortuuid.New()
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.AuthorID = autherID
	p.Comments = []Comments{}
	p.CreatedAt = time.Now()
}

//
func (p *Post) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	if p.AuthorID < 1 {
		return errors.New("Required Author")
	}
	return nil
}

//
func Prepare() User {
	var u User
	u.Nickname = nickname
	u.Email = userEmail
	u.VehicleID = vehicleID
	u.CreatedAt = time.Now()
	u.Picurl = picurl
	return u
}

//
func SavePost(client *mongo.Client, post Post) (Post, error) {
	var err error

	post.ID = shortuuid.New()
	post.AuthorID = autherID
	post.CreatedAt = time.Now()
	post.Author = Prepare()

	collection := client.Database("traceable").Collection("post")
	insertResult, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return post, err
}

//
func GetPostByID(client *mongo.Client, ID string) (Post, error) {
	//var err error
	var post Post

	//filter := bson.D{{"name", "Ash"}}
	collection := client.Database("traceable").Collection("post")
	filter := bson.D{{"id", ID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&post)

	return post, err

}

//
func FindAllPost(client *mongo.Client) ([]interface{}, error) {
	var err error
	post := []Post{}

	options := options.Find()
	options.SetSort(bson.D{{"_id", -1}})
	options.SetLimit(10)
	collection := client.Database("traceable").Collection("post")
	cur, err := collection.Find(context.Background(), bson.D{}, options)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cur)
	objectType := reflect.TypeOf(post).Elem()
	var list = make([]interface{}, 0)
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		result := reflect.New(objectType).Interface()
		err := cur.Decode(result)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		list = append(list, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return list, err
}
