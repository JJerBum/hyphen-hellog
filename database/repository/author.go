package repository

import (
	"context"
	"hyphen-hellog/cerrors"
	"hyphen-hellog/ent"
	"hyphen-hellog/ent/author"
	"hyphen-hellog/ent/comment"
	"hyphen-hellog/ent/post"
)

// PostAuthor 함수는 ctx, author를 매개변수로 받아 값을 저장하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) CreateAuthorX(ctx context.Context, author *ent.Author) *ent.Author {
	temp, err := d.Author.Create().
		SetAuthorID(author.AuthorID).
		Save(ctx)

	if err != nil {
		panic(cerrors.CreateErr{
			Err: err.Error(),
		})
	}

	return temp

}

// GetAuthor 함수는 ctx, ID를 매개변수로 받아 값을 조회하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) GetAuthorX(ctx context.Context, ID int) *ent.Author {
	temp, err := d.Author.
		Get(ctx, ID)

	if err != nil {
		panic(cerrors.SelectErr{
			Err: err.Error(),
		})
	}

	return temp
}

func (d *DBType) GetAuthorByAuthorIDX(ctx context.Context, authorID int) *ent.Author {
	temp, err := d.Author.
		Query().
		Where(author.AuthorID(authorID)).
		Only(ctx)

	if err != nil {
		panic(cerrors.SelectErr{
			Err: err.Error(),
		})
	}

	return temp
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

func (d *DBType) GetAuthorByPostIDX(ctx context.Context, postID int) *ent.Author {
	temp, err := d.Author.
		Query().
		Where(author.HasPostsWith(post.ID(postID))).
		Only(ctx)

	if err != nil {
		panic(cerrors.SelectErr{
			Err: err.Error(),
		})
	}

	return temp
}

func (d *DBType) GetAuthorByPostID(ctx context.Context, postID int) (*ent.Author, error) {
	return d.Author.
		Query().
		Where(author.HasPostsWith(post.ID(postID))).
		Only(ctx)

}

func (d *DBType) GetAuthorByCommentIDX(ctx context.Context, commentID int) *ent.Author {
	temp, err := d.Author.
		Query().
		Where(author.HasCommentsWith(comment.ID(commentID))).
		Only(ctx)

	if err != nil {
		panic(cerrors.SelectErr{
			Err: err.Error(),
		})
	}

	return temp

}

// UpdateAuthor 함수는 ctx, ID, author를 매개변수로 받아 값을 갱신하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) UpdateAuthorX(ctx context.Context, author *ent.Author) *ent.Author {
	temp, err := d.Author.UpdateOneID(author.ID).
		SetAuthorID(author.AuthorID).
		Save(ctx)

	if err != nil {
		panic(cerrors.UpdateErr{
			Err: err.Error(),
		})
	}

	return temp
}

// DeleteAuthor 함수는 ctx, ID를 매개변수로 받아 값을 삭제하는 함수 입니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) DeleteAuthorX(ctx context.Context, ID int) {
	err := d.Author.DeleteOneID(ID).
		Exec(ctx)

	if err != nil {
		panic(cerrors.DeleteErr{
			Err: err.Error(),
		})
	}
}
