package repository

import (
	"context"
	"hyphen-hellog/cerrors"
	"hyphen-hellog/ent"
	"hyphen-hellog/ent/comment"
	"hyphen-hellog/ent/post"
)

func (d *DBType) GetCommentsByPostIDX(ctx context.Context, postID int) []*ent.Comment {
	temp, err := d.Comment.
		Query().
		Where(comment.HasPostWith(post.ID(postID))).
		All(ctx)

	if err != nil {
		panic(cerrors.SelectErr{
			Err: err.Error(),
		})
	}

	return temp
}

func (d *DBType) GetCommentChildrenByCommentX(ctx context.Context, comment *ent.Comment) []*ent.Comment {
	temp, err := comment.QueryChildrens().All(ctx)
	if err != nil {
		panic(cerrors.SelectErr{
			Err: err.Error(),
		})
	}

	return temp
}

func (d *DBType) GetCommentParentByPostX(ctx context.Context, post *ent.Post) []*ent.Comment {
	temp, err := post.QueryComments().Where(comment.Not(comment.HasParent())).All(ctx)
	if err != nil {
		panic(cerrors.SelectErr{
			Err: err.Error(),
		})
	}

	return temp
}

// CreateComment 함수는 ctx, comment, parentCommentID, postID, authorID를 매개변수로 받아 데이터 값을 저장합니다.
// parentIDComment 값을 -1 이하의 값으로 전달 할 경우 null의 값으로 저장됩니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) CreateCommentX(ctx context.Context, comment *ent.Comment, parentCommentID int, postID int, authorID int) *ent.Comment {
	entity := d.Comment.Create().
		SetContent(comment.Content).
		SetPostID(postID).
		SetAuthorID(authorID)

	if parentCommentID > 0 {
		entity.SetParentID(parentCommentID)
	}

	temp, err := entity.Save(ctx)
	if err != nil {
		panic(cerrors.CreateErr{
			Err: err.Error(),
		})
	}

	return temp
}

func (d *DBType) GetCommentX(ctx context.Context, ID int) *ent.Comment {
	temp, err := d.Comment.
		Get(ctx, ID)

	if err != nil {
		panic(cerrors.SelectErr{
			Err: err.Error(),
		})
	}

	return temp.Edges.Parent
}

// CreateComment 함수는 ctx, comment, parentCommentID, postID, authorID를 매개변수로 받아 데이터 값을 갱신합니다.
// parentIDComment 값을 -1 이하의 값으로 전달 할 경우 null의 값으로 저장됩니다.
// 에러가 발생하면 패닉이 발생됩니다.
func (d *DBType) UpdateCommentX(ctx context.Context, comment *ent.Comment) *ent.Comment {
	temp, err := d.Comment.UpdateOneID(comment.ID).
		SetContent(comment.Content).
		Save(ctx)

	if err != nil {
		panic(cerrors.UpdateErr{
			Err: err.Error(),
		})
	}

	return temp

}

// DeleteComment 함수는 ctx, ID를 매개변수로 받아 값을 삭제합니다.
func (d *DBType) DeleteCommentX(ctx context.Context, ID int) {
	err := d.Comment.DeleteOneID(ID).
		Exec(ctx)

	if err != nil {
		panic(cerrors.DeleteErr{
			Err: err.Error(),
		})
	}
}
