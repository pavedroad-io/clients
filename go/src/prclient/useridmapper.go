/*
user ID mapper translates 3rd party sides credentials into internal user id.

HTTP verbs are translated into the following function calls:

Verbs to    keys                       Functions
----------  -------------------------- ---------
POST                                   Create
GET         /credential                Get
GET         /prUserIdMappersLIST       List resources
PUT         /credential                Replace
DELETE      /credential                Delete
PATCH       /credential                Edit

*/
package prclient

import (
	"context"
	"errors"
	"fmt"
)

// UserIdMappersService handles communication with the token related
// methods of the PavedRoad API.
type UserIdMappersService service

// prUserIdMapper data structure for token storage
type UserIdMapper struct {
  APIVersion string `json:"apiVersion"`
  ObjVersion string `json:"objVersion"`
  Kind       string `json:"kind"`
  Credential string `json:"login"`
  UserUUID   string `json:"userUUID"`
  LoginCount int    `json:"loginCount"`
  Created    string `json:"created,ignoreempty"`
  Updated    string `json:"updated,ignoreempty"`
  Active     string `json:"active"`
}

func (u UserIdMapper) String() string {
	return Stringify(u)
}

// Create a token
// PavedRoad API endpoint /prUserIdMappers/
func (s *UserIdMappersService) Create(ctx context.Context, newUserIdMapper UserIdMapper) (*UserIdMapper, *Response, error) {
	var u = fmt.Sprintf("%s/", mapperResource)

	req, err := s.client.NewRequest("POST", u, newUserIdMapper)
	if err != nil {
		return nil, nil, err
	}

  rUserIdMapper := &UserIdMapper{}
	resp, err := s.client.Do(ctx, req, rUserIdMapper)
	if err != nil {
		return nil, resp, err
	}

	return rUserIdMapper, resp, nil
}

// Get fetches a token using based on a Ucredential.
// PavedRoad API endpoint /prUserIdMappers/credential.
func (s *UserIdMappersService) Get(ctx context.Context, cred string) (*UserIdMapper, *Response, error) {
	var u string
	if cred != "" {
		u = fmt.Sprintf("%s/%v", mapperResource, cred)
	} else {
		return nil, nil, errors.New("Ucredential is required")
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(UserIdMapper)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// Delete a token using a Ucredential.
// PavedRoad API endpoint /prUserIdMappers/cred.
func (s *UserIdMappersService) Delete(ctx context.Context, cred string) (*Response, error) {
	var u string

	if cred != "" {
		u = fmt.Sprintf("%s/%v", mapperResource, cred)
	} else {
		return nil, errors.New("Ucredential is required")
	}
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Edit a token.
// PavedRoad API docs: https://developer.pavedroad.io/v1/token/#update-token
func (s *UserIdMappersService) Edit(ctx context.Context, token *UserIdMapper, cred string) (*UserIdMapper, *Response, error) {
	var u string
	if cred != "" {
		u = fmt.Sprintf("%s/%v", mapperResource, cred)
	} else {
		return nil, nil, errors.New("Ucredential is required")
	}

	req, err := s.client.NewRequest("PATCH", u, token)
	if err != nil {
		return nil, nil, err
	}

	tResp := new(UserIdMapper)
	resp, err := s.client.Do(ctx, req, tResp)
	if err != nil {
		return nil, resp, err
	}

	return tResp, resp, nil
}

// Replace a token.
// PavedRoad API docs: https://developer.pavedroad.io/v1/token/#replace-token
func (s *UserIdMappersService) Replace(ctx context.Context, token *UserIdMapper, cred string) (*UserIdMapper, *Response, error) {
	var u string
	if cred != "" {
		u = fmt.Sprintf("%s/%v", mapperResource, cred)
	} else {
		return nil, nil, errors.New("Ucredential is required")
	}

	req, err := s.client.NewRequest("PUT", u, token)
	if err != nil {
		return nil, nil, err
	}

	tResp := new(UserIdMapper)
	resp, err := s.client.Do(ctx, req, tResp)
	if err != nil {
		return nil, resp, err
	}

	return tResp, resp, nil
}


// UserIdMapperListOptions specifies optional parameters to the UserIdMappersService.ListAll
// method.
type UserIdMapperListOptions struct {
       // ID of the last token seen
       Since int64 `url:"since,omitempty"`

       // Note: Pagination is powered exclusively by the Since parameter,
       // ListOptions.Page has no effect.
       // ListOptions.PerPage controls an undocumented PavedRoad API parameter.
       ListOptions
}

// ListAll lists all PavedRoad token.
func (s *UserIdMappersService) List(ctx context.Context, opt *UserIdMapperListOptions) (*Response, error) {
	var u = fmt.Sprintf("%s/", mapperResourceList)

  // convert to our method of paging

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var token []*UserIdMapper
	resp, err := s.client.Do(ctx, req, &token)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
