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

func (s *Service) GetGuest(ctx context.Context, guestID *string) (*entity.Guest, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var response *entity.Guest

	_, err := s.Rest.R().
		SetPathParam("guest_id", *guestID).
		SetResult(&response).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Get(*s.Baseurl + "/guests/{guest_id}")
	if err != nil {
		log.Println(err)
	}

	return response, httpError
}

func (s *Service) CreateEmployee(ctx context.Context, name *string) (*entity.HTTPResponse, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var response *entity.HTTPResponse

	_, err := s.Rest.R().
		SetBody(&entity.CreateEmployeeRequest{
			Name: name,
		}).
		SetResult(&response).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Post(*s.Baseurl + "/employees")
	if err != nil {
		log.Println(err)
	}

	return response, httpError
}

func (s *Service) SearchEmployees(ctx context.Context, pageToken *string, pageSize *int) (*entity.SearchEmployeesResponse, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var response *entity.SearchEmployeesResponse

	_, err := s.Rest.R().
		SetQueryParams(map[string]string{
			"page_token": *pageToken,
			"page_size":  fmt.Sprint(*pageSize),
		}).
		SetResult(&response).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Get(*s.Baseurl + "/employees")
	if err != nil {
		log.Println(err)
	}

	return response, httpError
}

func (s *Service) GetEmployee(ctx context.Context, employeeID *string) (*entity.Employee, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var response *entity.Employee

	_, err := s.Rest.R().
		SetPathParam("employee_id", *employeeID).
		SetResult(&response).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Get(*s.Baseurl + "/employees/{employee_id}")
	if err != nil {
		log.Println(err)
	}

	return response, httpError
}

func (s *Service) CreatePlace(ctx context.Context, name *string) (*entity.HTTPResponse, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var response *entity.HTTPResponse

	_, err := s.Rest.R().
		SetBody(&entity.CreatePlaceRequest{
			Name: name,
		}).
		SetResult(&response).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Post(*s.Baseurl + "/places")
	if err != nil {
		log.Println(err)
	}

	return response, httpError
}

func (s *Service) SearchPlaces(ctx context.Context, pageToken *string, pageSize *int) (*entity.SearchPlacesResponse, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var response *entity.SearchPlacesResponse

	_, err := s.Rest.R().
		SetQueryParams(map[string]string{
			"page_token": *pageToken,
			"page_size":  fmt.Sprint(*pageSize),
		}).
		SetResult(&response).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Get(*s.Baseurl + "/places")
	if err != nil {
		log.Println(err)
	}

	return response, httpError
}

func (s *Service) GetPlace(ctx context.Context, placeID *string) (*entity.Place, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var response *entity.Place

	_, err := s.Rest.R().
		SetPathParam("place_id", *placeID).
		SetResult(&response).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Get(*s.Baseurl + "/places/{place_id}")
	if err != nil {
		log.Println(err)
	}

	return response, httpError
}

func (s *Service) CreateMenu(ctx context.Context, name *string) (*entity.HTTPResponse, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var response *entity.HTTPResponse

	_, err := s.Rest.R().
		SetBody(&entity.CreateMenuRequest{
			Name: name,
		}).
		SetResult(&response).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Post(*s.Baseurl + "/menus")
	if err != nil {
		log.Println(err)
	}

	return response, httpError
}

func (s *Service) SearchMenus(ctx context.Context, pageToken *string, pageSize *int) (*entity.SearchMenusResponse, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var response *entity.SearchMenusResponse

	_, err := s.Rest.R().
		SetQueryParams(map[string]string{
			"page_token": *pageToken,
			"page_size":  fmt.Sprint(*pageSize),
		}).
		SetResult(&response).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Get(*s.Baseurl + "/menus")
	if err != nil {
		log.Println(err)
	}

	return response, httpError
}

func (s *Service) GetMenu(ctx context.Context, menuID *string) (*entity.Menu, *entity.HTTPError) {
	var httpError *entity.HTTPError
	var response *entity.Menu

	_, err := s.Rest.R().
		SetPathParam("menu_id", *menuID).
		SetResult(&response).
		SetError(&httpError).
		SetHeader("Accept", "application/json").
		Get(*s.Baseurl + "/menus/{menu_id}")
	if err != nil {
		log.Println(err)
	}

	return response, httpError
}
