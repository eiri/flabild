package flabild

import (
	"errors"
	"plugin"
	"reflect"
	"testing"

	"github.com/mroth/weightedrand/v2"
)

type mockPlugin struct {
	lookupFn func(string) (plugin.Symbol, error)
}

func (m *mockPlugin) Lookup(s string) (plugin.Symbol, error) {
	return m.lookupFn(s)
}

func TestNewChooser(t *testing.T) {
	validMap := PairMap{
		NewPair('_', '_'): {
			weightedrand.NewChoice('a', 1),
			weightedrand.NewChoice('|', 1),
		},
	}

	tests := []struct {
		name      string
		plugin    flabildPlugin
		expectErr bool
		expectMap PairMap
	}{
		{
			name: "lookup failure",
			plugin: &mockPlugin{
				lookupFn: func(string) (plugin.Symbol, error) {
					return nil, errors.New("lookup failed")
				},
			},
			expectErr: true,
		},
		{
			name: "invalid API signature",
			plugin: &mockPlugin{
				lookupFn: func(string) (plugin.Symbol, error) {
					return 42, nil
				},
			},
			expectErr: true,
		},
		{
			name: "successful chooser creation",
			plugin: &mockPlugin{
				lookupFn: func(string) (plugin.Symbol, error) {
					return func() PairMap { return validMap }, nil
				},
			},
			expectMap: validMap,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewChooser(tt.plugin)

			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(c.choices, tt.expectMap) {
				t.Fatalf("choices mismatch")
			}
		})
	}
}

func TestChooserWord(t *testing.T) {
	tests := []struct {
		name      string
		choices   PairMap
		expect    string
		expectErr bool
	}{
		{
			name: "single character word",
			choices: PairMap{
				NewPair('_', '_'): {
					weightedrand.NewChoice('a', 1),
				},
				NewPair('_', 'a'): {
					weightedrand.NewChoice('|', 1),
				},
			},
			expect: "a",
		},
		{
			name: "empty word",
			choices: PairMap{
				NewPair('_', '_'): {
					weightedrand.NewChoice('|', 1),
				},
			},
			expect: "",
		},
		{
			name: "multi character word",
			choices: PairMap{
				NewPair('_', '_'): {
					weightedrand.NewChoice('a', 1),
				},
				NewPair('_', 'a'): {
					weightedrand.NewChoice('b', 1),
				},
				NewPair('a', 'b'): {
					weightedrand.NewChoice('|', 1),
				},
			},
			expect: "ab",
		},
		{
			name: "missing pair causes error",
			choices: PairMap{
				NewPair('_', '_'): {
					weightedrand.NewChoice('a', 1),
				},
			},
			expectErr: true,
		},
		{
			name: "invalid weighted chooser",
			choices: PairMap{
				NewPair('_', '_'): {},
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chooser{choices: tt.choices}
			word, err := c.Word()

			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if word != tt.expect {
				t.Fatalf("expected %q, got %q", tt.expect, word)
			}
		})
	}
}
