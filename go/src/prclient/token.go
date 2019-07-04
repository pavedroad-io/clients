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

// Get fetches a token using based on a users UUID.
// PavedRoad API endpoint /prTokens/:uuid.
func (s *TokensService) Get(ctx context.Context, uuid string) (*Token, *Response, error) {
	var u string
	if uuid != "" {
		u = fmt.Sprintf("%s/%v", prTokenResourceType, uuid)
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

// Edit the authenticated user.
//
// PavedRoad API docs: https://developer.github.com/v3/users/#update-the-authenticated-user
func (s *TokensService) Edit(ctx context.Context, user *Token) (*Token, *Response, error) {
	u := "user"
	req, err := s.client.NewRequest("PATCH", u, user)
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

// HovercardOptions specifies optional parameters to the TokensService.GetHovercard
// method.
type HovercardOptions struct {
	// SubjectType specifies the additional information to be received about the hovercard.
	// Possible values are: organization, repository, issue, pull_request. (Required when using subject_id.)
	SubjectType string `url:"subject_type"`

	// SubjectID specifies the ID for the SubjectType. (Required when using subject_type.)
	SubjectID string `url:"subject_id"`
}

// Hovercard represents hovercard information about a user.
type Hovercard struct {
	Contexts []*TokenContext `json:"contexts,omitempty"`
}

// TokenContext represents the contextual information about user.
type TokenContext struct {
	Message *string `json:"message,omitempty"`
	Octicon *string `json:"octicon,omitempty"`
}

// GetHovercard fetches contextual information about user. It requires authentication
// via Basic Auth or via OAuth with the repo scope.
//
// PavedRoad API docs: https://developer.github.com/v3/users/#get-contextual-information-about-a-user
func (s *TokensService) GetHovercard(ctx context.Context, user string, opt *HovercardOptions) (*Hovercard, *Response, error) {
	u := fmt.Sprintf("users/%v/hovercard", user)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	hc := new(Hovercard)
	resp, err := s.client.Do(ctx, req, hc)
	if err != nil {
		return nil, resp, err
	}

	return hc, resp, nil
}

// TokenListOptions specifies optional parameters to the TokensService.ListAll
// method.
type TokenListOptions struct {
	// ID of the last user seen
	Since int64 `url:"since,omitempty"`

	// Note: Pagination is powered exclusively by the Since parameter,
	// ListOptions.Page has no effect.
	// ListOptions.PerPage controls an undocumented PavedRoad API parameter.
	ListOptions
}

// ListAll lists all PavedRoad users.
//
// To paginate through all users, populate 'Since' with the ID of the last user.
//
// PavedRoad API docs: https://developer.github.com/v3/users/#get-all-users
func (s *TokensService) ListAll(ctx context.Context, opt *TokenListOptions) ([]*Token, *Response, error) {
	u, err := addOptions("users", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*Token
	resp, err := s.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

// AcceptInvitation accepts the currently-open repository invitation for the
// authenticated user.
//
// PavedRoad API docs: https://developer.github.com/v3/repos/invitations/#accept-a-repository-invitation
func (s *TokensService) AcceptInvitation(ctx context.Context, invitationID int64) (*Response, error) {
	u := fmt.Sprintf("user/repository_invitations/%v", invitationID)
	req, err := s.client.NewRequest("PATCH", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeclineInvitation declines the currently-open repository invitation for the
// authenticated user.
//
// PavedRoad API docs: https://developer.github.com/v3/repos/invitations/#decline-a-repository-invitation
func (s *TokensService) DeclineInvitation(ctx context.Context, invitationID int64) (*Response, error) {
	u := fmt.Sprintf("user/repository_invitations/%v", invitationID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
