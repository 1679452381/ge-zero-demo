package errorx

const DefaultCode = 1101

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"Msg"`
}
type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"Msg"`
}

func NewCodeError(code int, msg string) *CodeError {
	return &CodeError{
		code,
		msg,
	}
}

func NewDefaultCodeError(msg string) *CodeError {
	return &CodeError{
		DefaultCode,
		msg,
	}
}

func (e CodeError) Error() string {
	return e.Msg
}

func (e CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
