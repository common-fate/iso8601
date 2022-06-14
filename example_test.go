package iso8601_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/common-fate/iso8601"
)

func ExampleTime() {
	type MyStruct struct {
		CreatedAt iso8601.Time `json:"createdAt"`
	}

	s := MyStruct{
		CreatedAt: iso8601.New(time.Date(2000, 01, 01, 10, 0, 0, 0, time.UTC)),
	}

	str, _ := json.Marshal(s)

	fmt.Printf("%s", str)
	// Output: {"createdAt":"2000-01-01T10:00:00Z"}
}
