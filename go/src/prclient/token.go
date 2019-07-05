/*
Token implements access to prToken microservice which stores information about token to access 3rd party sites.

HTTP verbs are translated into the following function calls:

Verbs to Functions
------   ---------
POST     Create
GET      Get
GET/     List resources
PUT      Replace
DELETE   Delete
PATCH    Edit
*/
package prclient

import (
	"context"
	"errors"
	"fmt"
)

// TokensService handles communication with the token related
// methods of the PavedRoad API.
type TokensService service

// Token data structure for token storage
type Token struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Created    string   `json:"created,ignoreempty"`
	Updated    string   `json:"updated"`
	Active     bool     `json:"active"`
}

// Metadata stored for a token
type Metadata struct {
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	UID       string   `json:"uid"`
	Site      string   `json:"site"`
	EndPoint  string   `json:"endPoint"`
	Token     string   `json:"token"`
	Scope     []string `json:"scope"`
}

func (u Token) String() string {
	return Stringify(u)
}

// Create a token
// PavedRoad API endpoint /prTokens/
func (s *TokensService) Create(ctx context.Context, newToken Token) (*Token, *Response, error) {
	var u = fmt.Sprintf("%s/", tokenResource)

	req, err := s.client.NewRequest("POST", u, newToken)
	if err != nil {
		return nil, nil, err
	}

  rToken := &Token{}
	resp, err := s.client.Do(ctx, req, rToken)
	if err != nil {
		return nil, resp, err
	}

	return rToken, resp, nil
}

// Get fetches a token using based on a UUID.
// PavedRoad API endpoint /prTokens/uuid.
func (s *TokensService) Get(ctx context.Context, uuid string) (*Token, *Response, error) {
	var u string
	if uuid != "" {
		u = fmt.Sprintf("%s/%v", tokenResource, uuid)
	} else {
		return nil, nil, errors.New("UUID is required")
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(Token)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// Delete a token using a UUID.
// PavedRoad API endpoint /prTokens/uuid.
func (s *TokensService) Delete(ctx context.Context, uuid string) (*Response, error) {
	var u string

	if uuid != "" {
		u = fmt.Sprintf("%s/%v", tokenResource, uuid)
	} else {
		return nil, errors.New("UUID is required")
	}
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Edit a token.
// PavedRoad API docs: https://developer.pavedroad.io/v1/token/#update-token
func (s *TokensService) Edit(ctx context.Context, token *Token, uuid string) (*Token, *Response, error) {
	var u string
	if uuid != "" {
		u = fmt.Sprintf("%s/%v", tokenResource, uuid)
	} else {
		return nil, nil, errors.New("UUID is required")
	}

	req, err := s.client.NewRequest("PATCH", u, token)
	if err != nil {
		return nil, nil, err
	}

	tResp := new(Token)
	resp, err := s.client.Do(ctx, req, tResp)
	if err != nil {
		return nil, resp, err
	}

	return tResp, resp, nil
}

// Replace a token.
// PavedRoad API docs: https://developer.pavedroad.io/v1/token/#replace-token
func (s *TokensService) Replace(ctx context.Context, token *Token, uuid string) (*Token, *Response, error) {
	var u string
	if uuid != "" {
		u = fmt.Sprintf("%s/%v", tokenResource, uuid)
	} else {
		return nil, nil, errors.New("UUID is required")
	}

	req, err := s.client.NewRequest("PUT", u, token)
	if err != nil {
		return nil, nil, err
	}

	tResp := new(Token)
	resp, err := s.client.Do(ctx, req, tResp)
	if err != nil {
		return nil, resp, err
	}

	return tResp, resp, nil
}


// TokenListOptions specifies optional parameters to the TokensService.ListAll
// method.
type TokenListOptions struct {
       // ID of the last token seen
       Since int64 `url:"since,omitempty"`

       // Note: Pagination is powered exclusively by the Since parameter,
       // ListOptions.Page has no effect.
       // ListOptions.PerPage controls an undocumented PavedRoad API parameter.
       ListOptions
}

// ListAll lists all PavedRoad token.
func (s *TokensService) List(ctx context.Context, opt *TokenListOptions) (*Response, error) {
	var u = fmt.Sprintf("%s/", tokenResourceList)

  // convert to our method of paging

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var token []*Token
	resp, err := s.client.Do(ctx, req, &token)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
