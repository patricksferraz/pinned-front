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

type CreateEmployeeSchema struct {
	Name *string `json:"name" xml:"name" form:"name"`
}

type SearchEmployeesSchema struct {
	PageToken *string `query:"page_token"`
	PageSize  *int    `query:"page_size"`
}

type CreatePlaceSchema struct {
	Name *string `json:"name" xml:"name" form:"name"`
}

type SearchPlacesSchema struct {
	PageToken *string `query:"page_token"`
	PageSize  *int    `query:"page_size"`
}

type CreateMenuSchema struct {
	Name *string `json:"name" xml:"name" form:"name"`
}

type SearchMenusSchema struct {
	PageToken *string `query:"page_token"`
	PageSize  *int    `query:"page_size"`
}
