package prclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	_ "reflect"
	"testing"
)

// blankToken is an initialized object with defaults
var blankTokenJSON = `{"apiVersion":"1","kind":"","metadata":{"name":"","namespace":"","uid":"","site":"","endPoint":"","token":"","scope":null},"created":"","updated":"","active":false}`

var blankTokenObject = `{APIVersion:"", Kind:"", Metadata:prclient.Metadata{Name:"", Namespace:"", UID:"", Site:"", EndPoint:"", Token:""}, Created:"", Updated:"", Active:false}`

// fakeToken is created by NewToken which provides sample data
var fakeTokenObject = `{APIVersion:"core.pavedroad.io/v1alpha1", Kind:"PrToken", Metadata:prclient.Metadata{Name:"testoken", Namespace:"", UID:"", Site:"github", EndPoint:"https://api.github.com", Token:"#####################", Scope:["user" "repo"]}, Created:"", Updated:"", Active:true}`
var fakeTokenJSON = `{"apiVersion":"core.pavedroad.io/v1alpha1", "kind":"prToken", "metadata":{"name":"testoken", "namespace":"", "uid":"", "site":"github", "endPoint":"https://api.github.com", "token":"#####################", "scope":["user", "repo"]}, "created":"", "updated":"", "active":true}`

//create a token, set default values
func NewToken() (t *Token) {
	Token := Token{}

	Token.APIVersion = "core.pavedroad.io/v1alpha1"
	Token.Kind = "PrToken"
	Token.Metadata.Name = "testoken"
	Token.Metadata.UID = ""
	Token.Metadata.Site = "github"
	Token.Metadata.EndPoint = "https://api.github.com"
	Token.Metadata.Token = "#####################"
	Token.Metadata.Scope = append(Token.Metadata.Scope, "user")
	Token.Metadata.Scope = append(Token.Metadata.Scope, "repo")
	Token.Created = ""
	Token.Updated = ""
	Token.Active = true

	return &Token
}

func TestToken_marshall(t *testing.T) {
	u := &Token{APIVersion: "1"}
	testJSONMarshal(t, u, blankTokenJSON)
}

func TestTokensService_Get_specifiedToken(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, blankTokenJSON)
	})

	token, _, err := client.Token.Get(context.Background(), "1234")
	if err != nil {
		t.Errorf("Tokens.Get returned error: %v", err)
	}

	want := &Token{APIVersion: "1"}
	if !cmp.Equal(token, want) {
		t.Errorf("Tokens.Get returned %+v, want %+v", token, want)
	}
}

func TestTokensService_Get_invalidToken(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, _, err := client.Token.Get(context.Background(), "%")
	testURLParseError(t, err)
}

func TestTokensService_Delete_specifiedToken(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    testMethod(t, r, "DELETE")
  })

  _, err := client.Token.Delete(context.Background(), "1")

  if err != nil {
    t.Errorf("UserIdMapper delete returned error: %v", err)
  }
}

func TestTokenService_Replace(t *testing.T) {
  client, mux, _, teardown := setup()
  defer teardown()

  input := &Token{APIVersion: "1"}

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    v := &Token{APIVersion: "1"}
    json.NewDecoder(r.Body).Decode(v)
    testMethod(t, r, "PUT")
		fmt.Fprint(w, blankTokenJSON)
  })

  token, _, err := client.Token.Replace(context.Background(), input, "foo")
  if err != nil {
    t.Errorf("Token.Replace returned error: %v", err)
  }

	want := &Token{APIVersion: "1"}
	if !cmp.Equal(token, want) {
		t.Errorf("Tokens.Replace returned %+v, want %+v", token, want)
	}
}

func TestTokensService_List_Tokens(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
  var opt *TokenListOptions

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
    fmt.Fprint(w, "")
	})

  //var u = fmt.Sprintf("%s/", tokenResourceList)
  _, err := client.Token.List(context.Background(), opt )
	if err != nil {
		t.Errorf("Tokens.List returned error: %v", err)
	}
}


