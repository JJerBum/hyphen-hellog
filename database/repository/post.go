package repository

import (
	"context"
	"hyphen-hellog/ent"
	"hyphen-hellog/ent/author"
	"hyphen-hellog/ent/comment"
	"hyphen-hellog/ent/like"
	"hyphen-hellog/ent/post"
)

// PostAuthor 함수는 ctx, author를 매개변수로 받아 값을 저장하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) CreateAuthorX(ctx context.Context, author *ent.Author) *ent.Author {
	return d.Author.Create().
		SetAuthorID(author.AuthorID).
		SaveX(ctx)
}

// GetAuthor 함수는 ctx, ID를 매개변수로 받아 값을 조회하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) GetAuthorX(ctx context.Context, ID int) *ent.Author {
	return d.Author.
		GetX(ctx, ID)
}

func (d *DBType) GetAuthorXByAuthorID(ctx context.Context, authorID int) *ent.Author {
	return d.Author.
		Query().
		Where(author.AuthorID(authorID)).
		OnlyX(ctx)
}

func (d *DBType) GetAuthorByAuthorID(ctx context.Context, authorID int) (*ent.Author, error) {
	return d.Author.
		Query().
		Where(author.AuthorID(authorID)).
		Only(ctx)
}

// GetAuthor 함수는 ctx, ID를 매개변수로 받아 값을 조회하는 함수 입니다.
func (d *DBType) GetAuthor(ctx context.Context, ID int) (*ent.Author, error) {
	return d.Author.Get(ctx, ID)
}

func (d *DBType) GetAuthorXByPostID(ctx context.Context, postID int) *ent.Author {
	return d.Author.
		Query().
		Where(author.HasPostsWith(post.ID(postID))).
		OnlyX(ctx)
}

func (d *DBType) GetAuthorXByCommentID(ctx context.Context, commentID int) *ent.Author {
	return d.Author.
		Query().
		Where(author.HasCommentsWith(comment.ID(commentID))).
		OnlyX(ctx)
}

func (d *DBType) GetCommentsXByPostID(ctx context.Context, postID int) []*ent.Comment {
	return d.Comment.
		Query().
		Where(comment.HasPostWith(post.ID(postID))).
		AllX(ctx)
}

func (d *DBType) GetCommentChildrenXByComment(ctx context.Context, comment *ent.Comment) []*ent.Comment {
	return comment.QueryChildrens().AllX(ctx)
}

func (d *DBType) GetCommentParentXByPost(ctx context.Context, post *ent.Post) []*ent.Comment {
	return post.QueryComments().Where(comment.Not(comment.HasParent())).AllX(ctx)
}

// UpdateAuthor 함수는 ctx, ID, author를 매개변수로 받아 값을 갱신하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) UpdateAuthorX(ctx context.Context, author *ent.Author) *ent.Author {
	return d.Author.UpdateOneID(author.ID).
		SetAuthorID(author.AuthorID).
		SaveX(ctx)
}

// DeleteAuthor 함수는 ctx, ID를 매개변수로 받아 값을 삭제하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) DeleteAuthorX(ctx context.Context, ID int) {
	d.Author.DeleteOneID(ID).
		ExecX(ctx)
}

// CreateAuthor 함수는 ctx, post, authorID를 매개변수로 받아 데이터베이스 값을 저장하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) CreatePostX(ctx context.Context, post *ent.Post, authorID int) *ent.Post {
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
func (d *DBType) GetPostX(ctx context.Context, ID int) *ent.Post {
	return d.Post.
		GetX(ctx, ID)
}

// func (d *DBType) GetAuthor(ctx context.Context) *ent.Author {
// 	return d.Post.Query().QueryAuthor().Where()
// }

// CreateAuthor 함수는 ctx, post, authorID를 매개변수로 받아 데이터베이스 값을 갱신하 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) UpdatePostX(ctx context.Context, post *ent.Post, authorID int) *ent.Post {
	return d.Post.UpdateOneID(post.ID).
		SetTitle(post.Title).
		SetContent(post.Content).
		SetPreviewImage(post.PreviewImage).
		SetIsPrivate(post.IsPrivate).
		SetAuthorID(authorID).
		SaveX(ctx)
}

func (d *DBType) DeletePostX(ctx context.Context, ID int) {
	d.Post.DeleteOneID(ID).
		ExecX(ctx)
}

// CreateComment 함수는 ctx, comment, parentCommentID, postID, authorID를 매개변수로 받아 데이터 값을 저장합니다.
// parentIDComment 값을 -1 이하의 값으로 전달 할 경우 null의 값으로 저장됩니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) CreateCommentX(ctx context.Context, comment *ent.Comment, parentCommentID int, postID int, authorID int) *ent.Comment {
	temp := d.Comment.Create().
		SetContent(comment.Content).
		SetPostID(postID).
		SetAuthorID(authorID)

	if parentCommentID > 0 {
		temp = temp.SetParentID(parentCommentID)
	}

	return temp.SaveX(ctx)
}

func (d *DBType) GetCommentX(ctx context.Context, ID int) *ent.Comment {
	return d.Comment.
		GetX(ctx, ID)
}

// CreateComment 함수는 ctx, comment, parentCommentID, postID, authorID를 매개변수로 받아 데이터 값을 갱신합니다.
// parentIDComment 값을 -1 이하의 값으로 전달 할 경우 null의 값으로 저장됩니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) UpdateCommentX(ctx context.Context, comment *ent.Comment) *ent.Comment {
	return d.Comment.UpdateOneID(comment.ID).
		SetContent(comment.Content).
		SaveX(ctx)

}

// DeleteComment 함수는 ctx, ID를 매개변수로 받아 값을 삭제합니다.
func (d *DBType) DeleteCommentX(ctx context.Context, ID int) {
	d.Comment.DeleteOneID(ID).
		ExecX(ctx)
}

func (d *DBType) CreateLikeX(ctx context.Context, authorID int, PostID int) *ent.Like {
	return d.Like.Create().
		SetAuthorID(authorID).
		SetPostID(PostID).
		SaveX(ctx)
}

func (d *DBType) IsLikedXByAuthorID(ctx context.Context, authorID int) bool {
	return d.Like.Query().QueryAuthor().Where(author.ID(authorID)).ExistX(ctx)
}

func (d *DBType) UpdateLike(ctx context.Context, authorID int, PostID int) (*ent.Like, error) {
	return d.Like.Create().
		SetAuthorID(authorID).
		SetPostID(PostID).
		Save(ctx)
}

func (d *DBType) DeleteLikeX(ctx context.Context, authorID int, postID int) {

	d.Like.Delete().
		Where(
			like.HasAuthorWith(author.ID(authorID)),
			like.HasPostWith(post.ID(postID)),
		).
		ExecX(ctx)
}
