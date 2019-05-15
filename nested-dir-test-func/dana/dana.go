package dana

import (
	"encoding/json"

	"github.com/nuclio/nuclio-sdk-go"
)

func Sum(event nuclio.Event) int {
	numbersMap := make(map[string]int)
	json.Unmarshal(event.GetBody(), &numbersMap)

	sum := 0
	for _, v := range numbersMap {
		sum += v
	}
	return sum

}
