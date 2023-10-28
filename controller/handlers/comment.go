package handlers

import (
	"hyphen-hellog/cerrors"
	"hyphen-hellog/database"
	"hyphen-hellog/ent"
	"hyphen-hellog/model"

	"github.com/gofiber/fiber/v2"
)

func CreateComment(c *fiber.Ctx) error {
	clientRequest := new(model.InCreateComment).ParseX(c)

	database.Get().CreateCommentX(c.Context(),
		&ent.Comment{
			Content: clientRequest.Content,
		},
		clientRequest.ParentID,
		clientRequest.PostID,
		c.Locals("user").(*ent.Author).ID)

	return c.Status(fiber.StatusOK).JSON(model.General{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func GetComment(c *fiber.Ctx) error {
	clientRequest := new(model.InGetComments).ParseX(c)

	r := model.OutGetComments{}
	post := database.Get().GetPostX(c.Context(), clientRequest.PostID)
	commentParents := database.Get().GetCommentParentByPostX(c.Context(), post)

	// 한 포스트의 상위 댓글 loop
	for _, commentParent := range commentParents {

		// 상위 댓글의 하위 댓글들 추출
		commentChilds := database.Get().GetCommentChildrenByCommentX(c.Context(), commentParent)

		// 하위 댓글을 담을 변수
		newCommentChild := []model.CommentOfComment{}

		// 하위 댓글을 loop
		for _, commentChild := range commentChilds {
			// 값을 저장
			newCommentChild = append(newCommentChild, model.CommentOfComment{
				Comment: commentChild,
				Author:  database.Get().GetAuthorByCommentIDX(c.Context(), commentChild.ID),
			})
		}

		// Comments 저장
		r.Comments = append(r.Comments, model.Comment{
			Comment:          commentParent,
			Author:           database.Get().GetAuthorByCommentIDX(c.Context(), commentParent.ID),
			CommentOfComment: newCommentChild,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.General{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    r,
	})
}

func UpdateComment(c *fiber.Ctx) error {
	clientRequest := new(model.InUpdateComment).ParseX(c)

	// 이 사람이 접근해도 되는 사람인가?
	if database.Get().GetAuthorByCommentIDX(c.Context(), clientRequest.CommentID).ID != c.Locals("user").(*ent.Author).ID {
		panic(cerrors.WrongApproachErr{
			Err: "해당 사용자는 이 댓글의 주인이 아니므로 수정 하지 못합니다.",
		})
	}

	database.Get().UpdateCommentX(c.Context(), &ent.Comment{
		ID:      clientRequest.CommentID,
		Content: clientRequest.Content,
	})

	return c.Status(fiber.StatusOK).JSON(model.General{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func DeleteComment(c *fiber.Ctx) error {
	clientRequest := new(model.InDeleteComment).ParseX(c)

	// 이 사람이 접근해도 되는 사람인가?
	if database.Get().GetAuthorByCommentIDX(c.Context(), clientRequest.CommentID).ID != c.Locals("user").(*ent.Author).ID {
		panic(cerrors.WrongApproachErr{
			Err: "해당 사용자는 이 댓글의 주인이 아니므로 삭제 하지 못합니다.",
		})
	}

	database.Get().DeleteCommentX(c.Context(), clientRequest.CommentID)

	return c.Status(fiber.StatusOK).JSON(model.General{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}
