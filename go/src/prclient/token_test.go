package prclient

import (
_	"context"
_	"encoding/json"
_	"fmt"
_	"net/http"
_	"reflect"
	"testing"
)

//create a token, set default values
func NewToken () (t* Token) {
  Token := Token{}

  Token.APIVersion = "core.pavedroad.io/v1alpha1"
  Token.Kind = "PrToken"
  Token.Metadata.Name = "testoken"
  Token.Metadata.UID = ""
  Token.Metadata.Site = "github"
  Token.Metadata.EndPoint = "https://api.github.com"
  Token.Metadata.Token = "#####################"
  Token.Metadata.Scope =  append(Token.Metadata.Scope, "user")
  Token.Metadata.Scope =  append(Token.Metadata.Scope, "repo")
  Token.Created = ""
  Token.Updated = ""
  Token.Active = true

  return &Token
}

 func TestToken_marshall(t *testing.T) {
  blankToken := `{"apiVersion":"","kind":"","metadata":{"name":"","namespace":"","uid":"","site":"","endPoint":"","token":"","scope":null},"created":"","updated":"","active":false}`
 	testJSONMarshal(t, &Token{}, blankToken)
 	u :=  NewToken()
 
  want := `{
	"apiVersion": "core.pavedroad.io/v1alpha1",
	"kind": "PrToken",
	"metadata": {
		"name": "testoken",
		"namespace": "",
		"uid": "",
		"site": "github",
		"endPoint": "https://api.github.com",
		"token": "#####################",
		"scope": ["user", "repo"]
	},
	"created": "",
	"updated": "",
	"active": true
  }`

 	testJSONMarshal(t, u, want)
 }
 
// func TestUsersService_Get_authenticatedUser(t *testing.T) {
// 	client, mux, _, teardown := setup()
// 	defer teardown()
// 
// 	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		fmt.Fprint(w, `{"id":1}`)
// 	})
// 
// 	user, _, err := client.Users.Get(context.Background(), "")
// 	if err != nil {
// 		t.Errorf("Users.Get returned error: %v", err)
// 	}
// 
// 	want := &User{ID: Int64(1)}
// 	if !reflect.DeepEqual(user, want) {
// 		t.Errorf("Users.Get returned %+v, want %+v", user, want)
// 	}
// }
// 
// func TestUsersService_Get_specifiedUser(t *testing.T) {
// 	client, mux, _, teardown := setup()
// 	defer teardown()
// 
// 	mux.HandleFunc("/users/u", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		fmt.Fprint(w, `{"id":1}`)
// 	})
// 
// 	user, _, err := client.Users.Get(context.Background(), "u")
// 	if err != nil {
// 		t.Errorf("Users.Get returned error: %v", err)
// 	}
// 
// 	want := &User{ID: Int64(1)}
// 	if !reflect.DeepEqual(user, want) {
// 		t.Errorf("Users.Get returned %+v, want %+v", user, want)
// 	}
// }
// 
// func TestUsersService_Get_invalidUser(t *testing.T) {
// 	client, _, _, teardown := setup()
// 	defer teardown()
// 
// 	_, _, err := client.Users.Get(context.Background(), "%")
// 	testURLParseError(t, err)
// }
// 
// func TestUsersService_GetByID(t *testing.T) {
// 	client, mux, _, teardown := setup()
// 	defer teardown()
// 
// 	mux.HandleFunc("/user/1", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		fmt.Fprint(w, `{"id":1}`)
// 	})
// 
// 	user, _, err := client.Users.GetByID(context.Background(), 1)
// 	if err != nil {
// 		t.Fatalf("Users.GetByID returned error: %v", err)
// 	}
// 
// 	want := &User{ID: Int64(1)}
// 	if !reflect.DeepEqual(user, want) {
// 		t.Errorf("Users.GetByID returned %+v, want %+v", user, want)
// 	}
// }
// 
// func TestUsersService_Edit(t *testing.T) {
// 	client, mux, _, teardown := setup()
// 	defer teardown()
// 
// 	input := &User{Name: String("n")}
// 
// 	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
// 		v := new(User)
// 		json.NewDecoder(r.Body).Decode(v)
// 
// 		testMethod(t, r, "PATCH")
// 		if !reflect.DeepEqual(v, input) {
// 			t.Errorf("Request body = %+v, want %+v", v, input)
// 		}
// 
// 		fmt.Fprint(w, `{"id":1}`)
// 	})
// 
// 	user, _, err := client.Users.Edit(context.Background(), input)
// 	if err != nil {
// 		t.Errorf("Users.Edit returned error: %v", err)
// 	}
// 
// 	want := &User{ID: Int64(1)}
// 	if !reflect.DeepEqual(user, want) {
// 		t.Errorf("Users.Edit returned %+v, want %+v", user, want)
// 	}
// }
// 
// func TestUsersService_GetHovercard(t *testing.T) {
// 	client, mux, _, teardown := setup()
// 	defer teardown()
// 
// 	mux.HandleFunc("/users/u/hovercard", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		testHeader(t, r, "Accept", mediaTypeHovercardPreview)
// 		testFormValues(t, r, values{"subject_type": "repository", "subject_id": "20180408"})
// 		fmt.Fprint(w, `{"contexts": [{"message":"Owns this repository", "octicon": "repo"}]}`)
// 	})
// 
// 	opt := &HovercardOptions{SubjectType: "repository", SubjectID: "20180408"}
// 	hovercard, _, err := client.Users.GetHovercard(context.Background(), "u", opt)
// 	if err != nil {
// 		t.Errorf("Users.GetHovercard returned error: %v", err)
// 	}
// 
// 	want := &Hovercard{Contexts: []*UserContext{{Message: String("Owns this repository"), Octicon: String("repo")}}}
// 	if !reflect.DeepEqual(hovercard, want) {
// 		t.Errorf("Users.GetHovercard returned %+v, want %+v", hovercard, want)
// 	}
// }
// 
// func TestUsersService_ListAll(t *testing.T) {
// 	client, mux, _, teardown := setup()
// 	defer teardown()
// 
// 	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		testFormValues(t, r, values{"since": "1", "page": "2"})
// 		fmt.Fprint(w, `[{"id":2}]`)
// 	})
// 
// 	opt := &UserListOptions{1, ListOptions{Page: 2}}
// 	users, _, err := client.Users.ListAll(context.Background(), opt)
// 	if err != nil {
// 		t.Errorf("Users.Get returned error: %v", err)
// 	}
// 
// 	want := []*User{{ID: Int64(2)}}
// 	if !reflect.DeepEqual(users, want) {
// 		t.Errorf("Users.ListAll returned %+v, want %+v", users, want)
// 	}
// }
// 
// func TestUsersService_ListInvitations(t *testing.T) {
// 	client, mux, _, teardown := setup()
// 	defer teardown()
// 
// 	mux.HandleFunc("/user/repository_invitations", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		fmt.Fprintf(w, `[{"id":1}, {"id":2}]`)
// 	})
// 
// 	got, _, err := client.Users.ListInvitations(context.Background(), nil)
// 	if err != nil {
// 		t.Errorf("Users.ListInvitations returned error: %v", err)
// 	}
// 
// 	want := []*RepositoryInvitation{{ID: Int64(1)}, {ID: Int64(2)}}
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("Users.ListInvitations = %+v, want %+v", got, want)
// 	}
// }
// 
// func TestUsersService_ListInvitations_withOptions(t *testing.T) {
// 	client, mux, _, teardown := setup()
// 	defer teardown()
// 
// 	mux.HandleFunc("/user/repository_invitations", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		testFormValues(t, r, values{
// 			"page": "2",
// 		})
// 		fmt.Fprintf(w, `[{"id":1}, {"id":2}]`)
// 	})
// 
// 	_, _, err := client.Users.ListInvitations(context.Background(), &ListOptions{Page: 2})
// 	if err != nil {
// 		t.Errorf("Users.ListInvitations returned error: %v", err)
// 	}
// }
// func TestUsersService_AcceptInvitation(t *testing.T) {
// 	client, mux, _, teardown := setup()
// 	defer teardown()
// 
// 	mux.HandleFunc("/user/repository_invitations/1", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "PATCH")
// 		w.WriteHeader(http.StatusNoContent)
// 	})
// 
// 	if _, err := client.Users.AcceptInvitation(context.Background(), 1); err != nil {
// 		t.Errorf("Users.AcceptInvitation returned error: %v", err)
// 	}
// }
// 
// func TestUsersService_DeclineInvitation(t *testing.T) {
// 	client, mux, _, teardown := setup()
// 	defer teardown()
// 
// 	mux.HandleFunc("/user/repository_invitations/1", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "DELETE")
// 		w.WriteHeader(http.StatusNoContent)
// 	})
// 
// 	if _, err := client.Users.DeclineInvitation(context.Background(), 1); err != nil {
// 		t.Errorf("Users.DeclineInvitation returned error: %v", err)
// 	}
// }
