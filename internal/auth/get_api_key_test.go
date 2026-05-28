package auth

import (
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input map[string][]string
		want  string
	}

	tests := []test{
		{
			input: map[string][]string{
				"Authorization": {"ApiKey 12345"},
			},
			want: "12345",
		},
		{
			input: map[string][]string{
				"Authorization": {"ApiKey 67890"},
			},
			want: "67890",
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if got != tc.want {
			t.Fatalf("got %s, want %s", got, tc.want)
		}
		if err != nil {
			t.Fatalf("error: %v", err)
		}
	}
}
