package user

import (
	"encoding/json"
	"hyphen-hellog/cerrors"
	"hyphen-hellog/cerrors/exception"
	"hyphen-hellog/model/response"
	"io"
	"net/http"
)

var serverURL = "http://101.101.217.155:8081/api"

// Get()함수는 User token 매개변수를 이용해 마이크로서비스에게 /api/user/info로 요청하여 응답값을 반환하는 함수입니다.
func Get(token string) *response.GetUserInfo {

	userInfoModel := new(response.GetUserInfo)

	// 요청 헤더에 토큰 값을 설정합니다.
	req, err := http.NewRequest("GET", serverURL+"/user/info", nil)
	exception.Sniff(err)
	req.Header.Set("Authorization", token)

	// 요청을 보냅니다.
	client := &http.Client{}
	resp, err := client.Do(req)
	exception.Sniff(err)

	// 응답을 수신합니다.
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	exception.Sniff(err)

	err = json.Unmarshal(body, userInfoModel)
	exception.Sniff(err)

	if userInfoModel.Code != 200 {
		panic(cerrors.ErrInvalidRequest)
	}

	return userInfoModel
}

// Validate() 함수는 매개변수 token을 이용하여
func Validate(token string) *response.GetUserValidate {
	userValidateModel := new(response.GetUserValidate)

	// 요청 헤더에 토큰 값을 설정합니다.
	req, err := http.NewRequest("POST", serverURL+"/token/validate", nil)
	exception.Sniff(err)
	req.Header.Set("Authorization", token)

	// 요청을 보냅니다.
	client := &http.Client{}
	resp, err := client.Do(req)
	exception.Sniff(err)

	// 응답을 수신합니다.
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	exception.Sniff(err)

	err = json.Unmarshal(body, userValidateModel)
	exception.Sniff(err)

	if userValidateModel.Code != 200 {
		panic(cerrors.ErrInvalidRequest)
	}

	return userValidateModel
}
