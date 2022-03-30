package service

import (
	"context"
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

func (s *Service) CreateGuest(ctx context.Context, req *entity.CreateGuestRequest) (*entity.HTTPResponse, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var httpResponse *entity.HTTPResponse

	_, err := s.Rest.R().
		SetBody(req).
		SetResult(&httpResponse).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Post(*s.Baseurl + "/guests")
	if err != nil {
		log.Println(err)
	}

	return httpResponse, httpError
}
