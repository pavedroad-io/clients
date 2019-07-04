package main

import (
  "context"
  "prclient"
  "fmt"
)

func main(){

  client := prclient.NewClient(nil)
  ctx := context.Background()
  token, response, err := client.Token.Get(ctx, "12312312")
  fmt.Println(token, response, err)
}
