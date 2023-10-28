package repository

import (
	"context"
	"hyphen-hellog/cerrors"
	"hyphen-hellog/ent"
)

// CreateAuthor 함수는 ctx, post, authorID를 매개변수로 받아 데이터베이스 값을 저장하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) CreatePostX(ctx context.Context, post *ent.Post, authorID int) *ent.Post {
	temp, err := d.Post.Create().
		SetTitle(post.Title).
		SetContent(post.Content).
		SetPreviewImage(post.PreviewImage).
		SetIsPrivate(post.IsPrivate).
		SetAuthorID(authorID).
		Save(ctx)

	if err != nil {
		panic(cerrors.CreateErr{
			Err: err.Error(),
		})
	}

	return temp
}

// GetPost 함수는 ctx, ID를 매개변수로 받아 값을 조회하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) GetPostX(ctx context.Context, ID int) *ent.Post {
	temp, err := d.Post.
		Get(ctx, ID)

	if err != nil {
		panic(cerrors.SelectErr{
			Err: err.Error(),
		})
	}

	return temp
}

// func (d *DBType) GetAuthor(ctx context.Context) *ent.Author {
// 	return d.Post.Query().QueryAuthor().Where()
// }

// CreateAuthor 함수는 ctx, post, authorID를 매개변수로 받아 데이터베이스 값을 갱신하 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) UpdatePostX(ctx context.Context, post *ent.Post, authorID int) *ent.Post {
	temp, err := d.Post.UpdateOneID(post.ID).
		SetTitle(post.Title).
		SetContent(post.Content).
		SetPreviewImage(post.PreviewImage).
		SetIsPrivate(post.IsPrivate).
		SetAuthorID(authorID).
		Save(ctx)

	if err != nil {
		panic(cerrors.UpdateErr{
			Err: err.Error(),
		})
	}

	return temp
}

func (d *DBType) DeletePostX(ctx context.Context, ID int) {
	err := d.Post.DeleteOneID(ID).
		Exec(ctx)

	if err != nil {
		panic(cerrors.DeleteErr{
			Err: err.Error(),
		})
	}
}
