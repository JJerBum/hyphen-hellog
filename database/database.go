package database

import (
	"context"
	"hyphen-hellog/cerrors/exception"
	"hyphen-hellog/ent"
	"hyphen-hellog/ent/author"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

type databaseType ent.Client

var instance *databaseType = nil
var once sync.Once

func load(c config) {
	username := c.Get("DATASOURCE_USERNAME")
	password := c.Get("DATASOURCE_PASSWORD")
	host := c.Get("DATASOURCE_HOST")
	port := c.Get("DATASOURCE_PORT")
	dbName := c.Get("DATASOURCE_DB_NAME")
	maxPoolIdle, err := strconv.Atoi(c.Get("DATASOURCE_POOL_IDLE_CONN"))
	maxPoolOpen, err := strconv.Atoi(c.Get("DATASOURCE_POOL_MAX_CONN"))
	maxPollLifeTime, err := strconv.Atoi(c.Get("DATASOURCE_POOL_LIFE_TIME"))
	exception.Sniff(err)

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
	drv, err := sql.Open("mysql", dsn)
	exception.Sniff(err)

	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(maxPoolIdle)
	db.SetMaxOpenConns(maxPoolOpen)
	db.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)
	client := ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	var temp = databaseType(*client)
	instance = &temp
}

func New() *databaseType {
	if instance == nil {
		once.Do(func() {
			load(newConfig())
		})
	}

	return instance

}

// repository

// PostAuthor 함수는 ctx, author를 매개변수로 받아 값을 저장하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *databaseType) CreateAuthor(ctx context.Context, author *ent.Author) *ent.Author {
	return d.Author.Create().
		SetAuthorID(author.AuthorID).
		SaveX(ctx)
}

// GetAuthor 함수는 ctx, ID를 매개변수로 받아 값을 조회하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *databaseType) GetAuthorX(ctx context.Context, ID int) *ent.Author {
	return d.Author.
		GetX(ctx, ID)
}

func (d *databaseType) GetAuthorXByAuthorID(ctx context.Context, authorID int) *ent.Author {
	return d.Author.
		Query().
		Where(author.AuthorID(authorID)).
		OnlyX(ctx)
}

// GetAuthor 함수는 ctx, ID를 매개변수로 받아 값을 조회하는 함수 입니다.
func (d *databaseType) GetAuthor(ctx context.Context, ID int) (*ent.Author, error) {
	return d.Author.Get(ctx, ID)
}

// UpdateAuthor 함수는 ctx, ID, author를 매개변수로 받아 값을 갱신하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *databaseType) UpdateAuthor(ctx context.Context, author *ent.Author) *ent.Author {
	return d.Author.UpdateOneID(author.ID).
		SetAuthorID(author.AuthorID).
		SaveX(ctx)
}

// DeleteAuthor 함수는 ctx, ID를 매개변수로 받아 값을 삭제하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *databaseType) DeleteAuthor(ctx context.Context, ID int) {
	d.Author.DeleteOneID(ID).
		ExecX(ctx)
}

// CreateAuthor 함수는 ctx, post, authorID를 매개변수로 받아 데이터베이스 값을 저장하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *databaseType) CreatePost(ctx context.Context, post *ent.Post, authorID int) *ent.Post {
	return d.Post.Create().
		SetTitle(post.Title).
		SetContent(post.Content).
		SetPreviewImage(post.PreviewImage).
		SetIsPrivate(post.IsPrivate).
		SetAuthorID(authorID).
		SaveX(ctx)
}

// GetPost 함수는 ctx, ID를 매개변수로 받아 값을 조회하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *databaseType) GetPost(ctx context.Context, ID int) *ent.Post {
	return d.Post.
		GetX(ctx, ID)
}

// CreateAuthor 함수는 ctx, post, authorID를 매개변수로 받아 데이터베이스 값을 갱신하 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *databaseType) UpdatePost(ctx context.Context, post *ent.Post, authorID int) *ent.Post {
	return d.Post.UpdateOneID(post.ID).
		SetTitle(post.Title).
		SetContent(post.Content).
		SetPreviewImage(post.PreviewImage).
		SetIsPrivate(post.IsPrivate).
		SetAuthorID(authorID).
		SaveX(ctx)
}

func (d *databaseType) DeletePost(ctx context.Context, ID int) {
	d.Post.DeleteOneID(ID).
		ExecX(ctx)
}

// CreateComment 함수는 ctx, comment, parentCommentID, postID, authorID를 매개변수로 받아 데이터 값을 저장합니다.
// parentIDComment 값을 -1 이하의 값으로 전달 할 경우 null의 값으로 저장됩니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *databaseType) CreateComment(ctx context.Context, comment *ent.Comment, parentCommentID int, postID int, authorID int) *ent.Comment {
	temp := d.Comment.Create().
		SetContent(comment.Content).
		SetPostID(postID).
		SetAuthorID(authorID)

	if parentCommentID > 0 {
		temp = temp.SetParentID(parentCommentID)
	}

	return temp.SaveX(ctx)
}

func (d *databaseType) GetComment(ctx context.Context, ID int) *ent.Comment {
	return d.Comment.
		GetX(ctx, ID)
}

// CreateComment 함수는 ctx, comment, parentCommentID, postID, authorID를 매개변수로 받아 데이터 값을 갱신합니다.
// parentIDComment 값을 -1 이하의 값으로 전달 할 경우 null의 값으로 저장됩니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *databaseType) UpdateComment(ctx context.Context, comment *ent.Comment, parentCommentID int, postID int, authorID int) *ent.Comment {
	temp := d.Comment.UpdateOneID(comment.ID).
		SetContent(comment.Content).
		SetPostID(postID).
		SetAuthorID(authorID)

	if parentCommentID > 0 {
		temp = temp.SetParentID(parentCommentID)
	}

	return temp.SaveX(ctx)
}

// DeleteComment 함수는 ctx, ID를 매개변수로 받아 값을 삭제합니다.
func (d *databaseType) DeleteComment(ctx context.Context, ID int) {
	d.Comment.DeleteOneID(ID).
		ExecX(ctx)
}

// func (d *databaseType) CreateLike()
// func (d *databaseType) GetLike()
// func (d *databaseType) UpdateLike()
// func (d *databaseType) DeleteLike()
