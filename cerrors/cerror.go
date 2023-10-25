package cerrors

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// 예기치 않은 오류가 발생하면 ErrUnknown이 반환됩니다.
	ErrUnknown = Error("err_unknown: unknown error occurred")

	// 매개 변수 또는 요청 본문이 올바르지 않으면 ErrRInvalidRequest가 반환됩니다.
	ErrInvalidRequest = Error("err_invalid_request: invalid request received")

	// 매개 변수가 유효성 검사를 통과하지 못할 경우 ErrrValidation이 반환됩니다.
	ErrValidation = Error("err_validation: failed validation")

	// 요청한 리소스를 찾을 수 없을 때 ErrNotFound가 반환됩니다.
	ErrNotFound = Error("err_not_found: not found")

	ErrRequestFailed = Error("err_reqeust_failed: request to another server failed.")
)
