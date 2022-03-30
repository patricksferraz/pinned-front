package entity

type Session struct {
}

type HTTPError struct {
	Msg *string `json:"msg"`
}

type HTTPResponse struct {
	ID  *string `json:"id"`
	Msg *string `json:"msg"`
}

type CreateGuestRequest struct {
	Name *string `json:"name" xml:"name" form:"name"`
	Doc  *string `json:"doc" xml:"doc" form:"doc"`
}
