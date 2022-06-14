package iso8601

import (
	"encoding/json"
	"testing"
	"time"
)

func TestMarshalJSON(t *testing.T) {
	type testcase struct {
		name string
		give time.Time
		want string
	}

	testcases := []testcase{
		{"ok", time.Date(2022, 01, 01, 10, 0, 0, 0, time.UTC), `"2022-01-01T10:00:00Z"`},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(tc.give)
			if err != nil {
				t.Fatal(err)
			}
			if tc.want != string(got) {
				t.Fatalf("want: %s, got: %s", tc.want, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	type testcase struct {
		name string
		give Time
		want string
	}

	testcases := []testcase{
		{"ok", New(time.Date(2022, 01, 01, 10, 0, 0, 0, time.UTC)), `2022-01-01T10:00:00Z`},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.give.String()
			if tc.want != string(got) {
				t.Fatalf("want: %s, got: %s", tc.want, got)
			}
		})
	}
}
