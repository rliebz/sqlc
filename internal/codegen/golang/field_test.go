package golang

import "testing"

func Test_toLowerCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"a", "a"},
		{"A", "a"},
		{"id", "id"},
		{"ID", "id"},
		{"abc", "abc"},
		{"ABC", "abc"},
		{"FirstName", "firstName"},
		{"firstName", "firstName"},
		{"FooID", "fooID"},
		{"IDToken", "idToken"},
		{"HTTPToken", "httpToken"},
		{"XmlHTTPRequest", "xmlHTTPRequest"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toLowerCase(tt.input)
			if tt.want != got {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}
}
