package model

type InGetUserInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		CreatedDateTime  string `json:"createdDateTime"`
		ModifiedDateTime string `json:"modifiedDateTime"`
		ID               int    `json:"id"`
		Email            string `json:"email"`
		Name             string `json:"name"`
		Image            any    `json:"image"`
		SocialID         any    `json:"socialId"`
		SocialType       any    `json:"socialType"`
		UserStatus       string `json:"userStatus"`
		UserRole         string `json:"userRole"`
	} `json:"data"`
}

type InGetUserValidate struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    int    `json:"data"`
}
