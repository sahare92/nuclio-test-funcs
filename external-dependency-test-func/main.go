package main

import (
	"encoding/json"
	"fmt"
	"github.com/nuclio/nuclio-sdk-go"
	"os"
)

func Handler(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	numbersMap := make(map[string]int)
	json.Unmarshal(event.GetBody(), &numbersMap)
	context.Logger.DebugWith("numbersMap is", "numbersMap", numbersMap)

	sum := 0
	for _, v := range numbersMap {
		sum += v
	}

	if naipiVar := os.Getenv("NAIPI_VAR"); naipiVar != "naipi_var" {
		return nuclio.Response{
			StatusCode:  400,
			ContentType: "application/text",
			Body:        []byte("Missing env NAIPI_VAR"),
		}, nil
	}

	return nuclio.Response{
		StatusCode:  200,
		ContentType: "application/text",
		Body:        []byte(fmt.Sprintf("%d", sum)),
	}, nil
}