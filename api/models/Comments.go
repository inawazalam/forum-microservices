package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//
type Comments struct {
	Content   string `gorm:"size:255;not null;" json:"content"`
	CreatedAt time.Time
	Author    User `json:"author"`
}

//
func CommentOnPost(client *mongo.Client, postComment CommentRequest) (Post, error) {
	var err error
	var preData Post
	var updatePost Post
	var comments Comments
	//Comments.Author = Prepare()
	preData, err = GetPostByID(client, postComment.ID)
	comments.Content = postComment.Content
	comments.Author = Prepare()
	comments.CreatedAt = time.Now()
	updatePost = preData
	updatePost.Comments = append(updatePost.Comments, comments)

	fmt.Println(updatePost)
	update := bson.D{{"$set", bson.D{{"comments", updatePost.Comments}}}}

	collection := client.Database("traceable").Collection("post")

	insertResult, err := collection.UpdateOne(context.TODO(), preData, update)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(insertResult)

	return updatePost, err
}
