package golang

import (
	"testing"

	"github.com/sqlc-dev/sqlc/internal/codegen/golang/opts"
)

func TestStructName(t *testing.T) {
	options := &opts.Options{
		Rename: map[string]string{
			"foo": "foorenamed",
			"bar": "barrenamed",
		},
		InitialismsMap: map[string]struct{}{
			"id":  {},
			"url": {},
		},
	}

	tests := []struct {
		name string
		want string
	}{
		// Simple
		{"a", "A"},
		{"word", "Word"},
		{"foo_bar", "FooBar"},
		{"foo_bar_baz", "FooBarBaz"},

		// Initialisms
		{"id", "ID"},
		{"url", "URL"},
		{"ids", "IDs"},
		{"urls", "URLs"},
		{"id_token", "IDToken"},
		{"foo_id", "FooID"},
		{"foo_ids", "FooIDs"},
		{"some_url_thing", "SomeURLThing"},

		// Words containing initialisms
		{"kid", "Kid"},
		{"curl", "Curl"},
		{"identity", "Identity"},
		{"identities", "Identities"},
		{"the_kid", "TheKid"},
		{"identity_crisis", "IdentityCrisis"},
		{"curly", "Curly"},

		// Rename
		{"foo", "foorenamed"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StructName(tt.name, options)
			if tt.want != got {
				t.Errorf("want StructName %q got %q", tt.want, got)
			}
		})
	}
}
