package response

import "hyphen-hellog/ent"

type GetPost struct {
	*ent.Post   `json:"post"`
	*ent.Author `json:"author"`
}
