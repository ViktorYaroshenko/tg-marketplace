package main

import "testing"

func TestParseAddCommand(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantName string
		wantQty  int
		wantErr  bool
	}{
		{
			name:    "zero_quantity",
			input:   "/add Яблоки 0",
			wantErr: true,
		},
		{
			name:    "invalid_quantity",
			input:   "/add Яблоки abc",
			wantErr: true,
		},
		{
			name:    "missing_product_name",
			input:   "/add 10",
			wantErr: true,
		},
		{
			name:    "missing_arguments",
			input:   "/add",
			wantErr: true,
		},
		{
			name:    "empty_command",
			input:   "",
			wantErr: true,
		},
		{
			name:    "wrong_command",
			input:   "/remove Яблоки 10",
			wantErr: true,
		},
		{
			name:    "quantity_overflow",
			input:   "/add Яблоки 999999999999999999999",
			wantErr: true,
		},
		{
			name:     "extra_spaces",
			input:    "  /add   Зеленые   яблоки   10  ",
			wantName: "Зеленые яблоки",
			wantQty:  10,
			wantErr:  false,
		},
		{
			name:     "valid command",
			input:    "/add Яблоки 10",
			wantName: "Яблоки",
			wantQty:  10,
			wantErr:  false,
		},
		{
			name:     "multiword product name",
			input:    "/add Зелёные яблоки 15",
			wantName: "Зелёные яблоки",
			wantQty:  15,
			wantErr:  false,
		},
		{
			name:    "negative quantity",
			input:   "/add Яблоки -5",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, qty, err := parseAddCommand(tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}

			if tt.wantErr {
				return
			}

			if name != tt.wantName {
				t.Errorf("expected name %q, got %q", tt.wantName, name)
			}

			if qty != tt.wantQty {
				t.Errorf("expected quantity %d, got %d", tt.wantQty, qty)
			}
		})
	}
}
