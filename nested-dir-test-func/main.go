package main

import (
	"fmt"

	"github.com/nuclio/nuclio-sdk-go"
	"github.com/nuclio/handler/dana"
)

func Handler(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	return nuclio.Response{
		StatusCode:  200,
		ContentType: "application/text",
		Body:        []byte(fmt.Sprintf("%d", dana.Sum(event))),
	}, nil
}
