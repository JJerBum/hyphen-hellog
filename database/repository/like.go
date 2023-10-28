package repository

import (
	"context"
	"hyphen-hellog/cerrors"
	"hyphen-hellog/ent"
	"hyphen-hellog/ent/author"
	"hyphen-hellog/ent/like"
	"hyphen-hellog/ent/post"
)

func (d *DBType) CreateLikeX(ctx context.Context, authorID int, PostID int) *ent.Like {
	temp, err := d.Like.Create().
		SetAuthorID(authorID).
		SetPostID(PostID).
		Save(ctx)

	if err != nil {
		panic(cerrors.CreateErr{
			Err: err.Error(),
		})
	}

	return temp
}

func (d *DBType) IsLikedByAuthorIDX(ctx context.Context, authorID int) bool {
	temp, err := d.Like.Query().QueryAuthor().Where(author.ID(authorID)).Exist(ctx)
	if err != nil {
		panic(cerrors.SelectErr{
			Err: err.Error(),
		})
	}

	return temp
}

func (d *DBType) UpdateLike(ctx context.Context, authorID int, PostID int) (*ent.Like, error) {
	return d.Like.Create().
		SetAuthorID(authorID).
		SetPostID(PostID).
		Save(ctx)
}

func (d *DBType) DeleteLikeX(ctx context.Context, authorID int, postID int) {

	_, err := d.Like.Delete().
		Where(
			like.HasAuthorWith(author.ID(authorID)),
			like.HasPostWith(post.ID(postID)),
		).
		Exec(ctx)

	if err != nil {
		panic(cerrors.DeleteErr{
			Err: err.Error(),
		})
	}
}
