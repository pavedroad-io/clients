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

// blankUserIdMapper is an initialized object with defaults
var blankUserIdMapperJSON = `{
  "apiVersion" :"1",
  "objVersion": "",
  "kind" : "",
  "login": "",
  "userUUID": "",
  "loginCount": 0,
  "created": "",
  "updated": "",
  "active": ""
}`

var blanUserIdMapperObject = UserIdMapper{
    APIVersion: "core.pavedroad.io/v1alpha1",
    Kind:       "prUserIdMapper",
    ObjVersion: "v1beta1",
    Credential: "",
    UserUUID:   "",
    LoginCount: 0,
    Created:    "2002-10-02T15:00:00.05Z",
    Updated:    "2002-10-02T15:00:00.05Z",
    Active:     "true"}

func TestUserIdMapper_marshall(t *testing.T) {
	u := &UserIdMapper{APIVersion: "1"}
	testJSONMarshal(t, u, blankUserIdMapperJSON)
}

func TestUserIdMappersService_Get_specifiedUserIdMapper(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, blankUserIdMapperJSON)
	})

	mapping, _, err := client.UserIdMapper.Get(context.Background(), "1234")
	if err != nil {
		t.Errorf("UserIdMappers.Get returned error: %v", err)
	}

	want := &UserIdMapper{APIVersion: "1"}
	if !cmp.Equal(mapping, want) {
		t.Errorf("UserIdMappers.Get returned %+v, want %+v", mapping, want)
	}
}

func TestUserIdMappersService_Get_invalidUserIdMapper(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, _, err := client.UserIdMapper.Get(context.Background(), "%")
	testURLParseError(t, err)
}

func TestUserIdMappersService_Delete_specifiedUserIdMapper(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    testMethod(t, r, "DELETE")
  })

  _, err := client.UserIdMapper.Delete(context.Background(), "1")

  if err != nil {
    t.Errorf("UserIdMappers delete returned error: %v", err)
  }
}

func TestUserIdMapperService_Replace(t *testing.T) {
  client, mux, _, teardown := setup()
  defer teardown()

  input := &UserIdMapper{APIVersion: "1"}

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    v := &UserIdMapper{APIVersion: "1"}
    json.NewDecoder(r.Body).Decode(v)
    testMethod(t, r, "PUT")
		fmt.Fprint(w, blankUserIdMapperJSON)
  })

  token, _, err := client.UserIdMapper.Replace(context.Background(), input, "foo")
  if err != nil {
    t.Errorf("UserIdMapper.Replace returned error: %v", err)
  }

	want := &UserIdMapper{APIVersion: "1"}
	if !cmp.Equal(token, want) {
		t.Errorf("UserIdMappers.Replace returned %+v, want %+v", token, want)
	}
}

func TestUserIdMappersService_List_UserIdMappers(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
  var opt *UserIdMapperListOptions

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
    fmt.Fprint(w, "")
	})

  //var u = fmt.Sprintf("%s/", tokenResourceList)
  _, err := client.UserIdMapper.List(context.Background(), opt )
	if err != nil {
		t.Errorf("UserIdMappers.List returned error: %v", err)
	}
}
