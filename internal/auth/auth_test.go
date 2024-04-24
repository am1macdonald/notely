package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input http.Header
		want  string
	}
	tests := []test{
		{input: http.Header{
			"Authorization": []string{"ApiKey test"},
		}, want: "test"},
		{input: http.Header{
			"Authorization": []string{"ApiKey"},
		}, want: ""},
		{input: http.Header{
			"Authorization": []string{"Bearer test"},
		}, want: ""},
		{input: http.Header{
			"Authorization": []string{},
		}, want: ""},
		{input: http.Header{}, want: "test"},
	}
	for _, tc := range tests {
		got, _ := GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

}
