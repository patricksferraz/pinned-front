package entity

import "time"

type Base struct {
	ID        *string    `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type HTTPError struct {
	Msg *string `json:"msg"`
}

type HTTPResponse struct {
	ID  *string `json:"id"`
	Msg *string `json:"msg"`
}

type CreateGuestRequest struct {
	Name *string `json:"name"`
	Doc  *string `json:"doc"`
}

type SearchGuestsRequest struct {
	PageToken *string `json:"page_token"`
	PageSize  *int    `json:"page_size"`
}

type SearchGuestsResponse struct {
	Guests        []*Guest `json:"guests"`
	NextPageToken *string  `json:"next_page_token"`
}

type Guest struct {
	Base `json:",inline"`
	Name *string `json:"name"`
	Doc  *string `json:"doc"`
}
