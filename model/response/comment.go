package response

import "hyphen-hellog/ent"

type GetComments struct {
	Comments []Comment
}

type Comment struct {
	*ent.Comment     `json:"comment"`
	*ent.Author      `json:"author"`
	CommentOfComment []CommentOfComment `json:"comment_of_comment"`
}

type CommentOfComment struct {
	*ent.Comment `json:"comment"`
	*ent.Author  `json:"author"`
}
