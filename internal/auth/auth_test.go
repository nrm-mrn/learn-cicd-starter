package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type Result struct {
		value string
		err   error
	}
	tests := map[string]struct {
		input http.Header
		want  Result
	}{
		"first_h": {
			input: http.Header{"Authorization": []string{"ApiKey 43214234"}},
			want:  Result{"43214234", nil},
		},
		"second_h": {
			input: http.Header{"Authorization": []string{"43214234"}},
			want:  Result{"", errors.New("malformed authorization   header")},
		},
		"third_h": {
			input: http.Header{"content/type": []string{"43214234"}},
			want:  Result{"", ErrNoAuthHeaderIncluded},
		},
	}

	for name, tc := range tests {
		val, err := GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.want.value, val) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want.value, val)
		}
		if !reflect.DeepEqual(tc.want.err, err) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want.value, val)
		}
	}
}
