package response

import "hyphen-hellog/ent"

type GetPost struct {
	*ent.Post   `json:"post"`
	IsLiked     bool `json:"is_liked"`
	*ent.Author `json:"author"`
}
