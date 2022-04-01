package service

import (
	"context"
	"fmt"
	"log"

	"github.com/c-4u/pinned-front/domain/entity"
	"github.com/go-resty/resty/v2"
)

type Service struct {
	Rest    *resty.Client
	Baseurl *string
}

func NewService(rest *resty.Client, baseurl *string) *Service {
	return &Service{
		Rest:    rest,
		Baseurl: baseurl,
	}
}

func (s *Service) CreateGuest(ctx context.Context, name, doc *string) (*entity.HTTPResponse, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var response *entity.HTTPResponse

	_, err := s.Rest.R().
		SetBody(&entity.CreateGuestRequest{
			Name: name,
			Doc:  doc,
		}).
		SetResult(&response).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Post(*s.Baseurl + "/guests")
	if err != nil {
		log.Println(err)
	}

	return response, httpError
}

func (s *Service) SearchGuests(ctx context.Context, pageToken *string, pageSize *int) (*entity.SearchGuestsResponse, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var response *entity.SearchGuestsResponse

	_, err := s.Rest.R().
		SetQueryParams(map[string]string{
			"page_token": *pageToken,
			"page_size":  fmt.Sprint(*pageSize),
		}).
		SetResult(&response).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Get(*s.Baseurl + "/guests")
	if err != nil {
		log.Println(err)
	}

	return response, httpError
}
