// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

type UserRegisterQequest struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type IdRequest struct {
	Id string `json:"id" path:"id"`
}

type UserResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type LoginQequest struct {
	Id string `json:"id"`
}

type LoginReply struct {
	StatusCode int
	Message    string
	Token      TokenInfo
}
type TokenInfo struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}
