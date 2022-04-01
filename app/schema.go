package app

type Session struct {
}

type CreateGuestSchema struct {
	Name *string `json:"name" xml:"name" form:"name"`
	Doc  *string `json:"doc" xml:"doc" form:"doc"`
}

type SearchGuestsSchema struct {
	PageToken *string `query:"page_token"`
	PageSize  *int    `query:"page_size"`
}
