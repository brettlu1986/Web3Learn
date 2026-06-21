package walllet

import "testing"

func TestValidateAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		wantErr bool
	}{
		{
			name:    "valid address",
			address: "0x1234567890123456789012345678901234567890",
			wantErr: false,
		},
		{
			name:    "empty address",
			address: "",
			wantErr: true,
		},
		{
			name:    "missing 0x",
			address: "1234567890123456789012345678901234567890",
			wantErr: true,
		},
		{
			name:    "wrong length",
			address: "0x123",
			wantErr: true,
		},
		{
			name:    "invalid hex char",
			address: "0x123456789012345678901234567890123456789Z",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAddress(tt.address)

			if tt.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("expected nil, got %v", err)
			}
		})
	}
}
