package handlers

import (
	"hyphen-hellog/database"
	"hyphen-hellog/ent"
	"hyphen-hellog/model/request"
	"hyphen-hellog/model/response"
	"hyphen-hellog/verifier"

	"github.com/gofiber/fiber/v2"
)

func CreateComment(c *fiber.Ctx) error {
	clientRequest := new(request.CreateComment).Parse(c)
	verifier.Validate(c)

	database.Get().CreateCommentX(c.Context(),
		&ent.Comment{
			Content: clientRequest.Content,
		},
		clientRequest.ParentID,
		clientRequest.PostID,
		c.Locals("user").(*ent.Author).ID)

	return c.Status(fiber.StatusOK).JSON(response.Genreal{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func GetComment(c *fiber.Ctx) error {
	clientRequest := new(request.GetComments).Parse(c)
	verifier.Validate(c)

	r := response.GetComments{}
	post := database.Get().GetPostX(c.Context(), clientRequest.PostID)
	commentParents := database.Get().GetCommentParentXByPost(c.Context(), post)

	// 한 포스트의 상위 댓글 loop
	for _, commentParent := range commentParents {

		// 상위 댓글의 하위 댓글들 추출
		commentChilds := database.Get().GetCommentChildrenXByComment(c.Context(), commentParent)

		// 하위 댓글을 담을 변수
		newCommentChild := []response.CommentOfComment{}

		// 하위 댓글을 loop
		for _, commentChild := range commentChilds {
			// 값을 저장
			newCommentChild = append(newCommentChild, response.CommentOfComment{
				Comment: commentChild,
				Author:  database.Get().GetAuthorXByCommentID(c.Context(), commentChild.ID),
			})
		}

		// Comments 저장
		r.Comments = append(r.Comments, response.Comment{
			Comment:          commentParent,
			Author:           database.Get().GetAuthorXByCommentID(c.Context(), commentParent.ID),
			CommentOfComment: newCommentChild,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Genreal{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    r,
	})
}

func UpdateComment(c *fiber.Ctx) error {
	clientRequest := new(request.UpdateComment).Parse(c)

	// 이 사람이 접근해도 되는 사람인가?
	if database.Get().GetAuthorXByCommentID(c.Context(), clientRequest.CommentID).ID != c.Locals("user").(*ent.Author).ID {
		panic("잘못된 접근 입니다. 이 사용자는 이 댓글의 주인이 아닙니다.")
	}

	database.Get().UpdateCommentX(c.Context(), &ent.Comment{
		ID:      clientRequest.CommentID,
		Content: clientRequest.Content,
	})

	return c.Status(fiber.StatusOK).JSON(response.Genreal{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func DeleteComment(c *fiber.Ctx) error {
	clientRequest := new(request.DeleteComment).Parse(c)

	// 이 사람이 접근해도 되는 사람인가?
	if database.Get().GetAuthorXByCommentID(c.Context(), clientRequest.CommentID).ID != c.Locals("user").(*ent.Author).ID {
		panic("잘못된 접근 입니다. 이 사용자는 이 댓글의 주인이 아닙니다.")
	}

	database.Get().DeleteCommentX(c.Context(), clientRequest.CommentID)

	return c.Status(fiber.StatusOK).JSON(response.Genreal{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}
