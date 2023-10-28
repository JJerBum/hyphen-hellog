package user

import (
	"encoding/json"
	"hyphen-hellog/cerrors"
	"hyphen-hellog/model"
	"io"
	"net/http"
)

var serverURL = "http://101.101.217.155:8081/api"

// Get()함수는 User token 매개변수를 이용해 마이크로서비스에게 /api/user/info로 요청하여 응답값을 반환하는 함수입니다.
func Get(token string) *model.InGetUserInfo {

	userInfoModel := new(model.InGetUserInfo)

	// 요청 헤더에 토큰 값을 설정합니다.
	req, err := http.NewRequest("GET", serverURL+"/user/info", nil)
	if err != nil {
		panic(cerrors.RequestFailedErr{
			Err: err.Error(),
		})
	}
	req.Header.Set("Authorization", token)

	// 요청을 보냅니다.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(cerrors.RequestFailedErr{
			Err: err.Error(),
		})
	}

	// 응답을 수신합니다.
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(cerrors.RequestFailedErr{
			Err: err.Error(),
		})
	}

	err = json.Unmarshal(body, userInfoModel)
	if err != nil {
		panic(cerrors.RequestFailedErr{
			Err: err.Error(),
		})
	}

	if userInfoModel.Code != 200 {
		panic(cerrors.RequestFailedErr{
			Err: "Response HTTP 1.1 Status 200이 아닙니다.",
		})
	}

	return userInfoModel
}

// Validate() 함수는 매개변수 token을 이용하여
func Validate(token string) (*model.InGetUserValidate, error) {
	userValidateModel := new(model.InGetUserValidate)

	// 요청 헤더에 토큰 값을 설정합니다.
	req, err := http.NewRequest("POST", serverURL+"/token/validate", nil)
	if err != nil {
		return nil, cerrors.RequestFailedErr{
			Err: err.Error(),
		}
	}

	req.Header.Set("Authorization", token)

	// 요청을 보냅니다.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, cerrors.RequestFailedErr{
			Err: err.Error(),
		}
	}

	// 응답을 수신합니다.
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, cerrors.RequestFailedErr{
			Err: err.Error(),
		}
	}

	if len(body) == 0 {
		return nil, cerrors.RequestFailedErr{
			Err: "response body len is 0",
		}
	}

	err = json.Unmarshal(body, userValidateModel)
	if err != nil {
		return nil, cerrors.RequestFailedErr{
			Err: err.Error(),
		}
	}

	if userValidateModel.Code != 200 {
		panic(cerrors.UnauthorizedErr{
			Err: "Response HTTP 1.1 Status 200이 아닙니다.",
		})
	}

	return userValidateModel, nil
}
