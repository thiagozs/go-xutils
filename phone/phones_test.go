package phone

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		name    string
		phone   string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid mobile phone number with country code",
			phone:   "+55 11 98765-4321",
			want:    "11987654321",
			wantErr: false,
		},
		{
			name:    "Valid mobile phone number without country code",
			phone:   "11 98765-4321",
			want:    "11987654321",
			wantErr: false,
		},
		{
			name:    "Valid mobile phone number with country code parantheses",
			phone:   "+55 (11) 98765-4321",
			want:    "11987654321",
			wantErr: false,
		},
		{
			name:    "Valid mobile phone number without country code parantheses and dash",
			phone:   "(11) 98765-4321",
			want:    "11987654321",
			wantErr: false,
		},
		{
			name:    "Invalid phone number",
			phone:   "00000000",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid phone number with country code",
			phone:   "+55 00 000000000",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid phone number with country code and parantheses",
			phone:   "+55 (00) 000000000",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid phone number with country code and parantheses and dash",
			phone:   "+55 (00) 00000-0000",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Valid a residential phone number with country code",
			phone:   "+55 19 3611-4444",
			want:    "1936114444",
			wantErr: false,
		},
		{
			name:    "Valid a residential phone number with country code and parantheses and dash",
			phone:   "+55 (19) 3611-4444",
			want:    "1936114444",
			wantErr: false,
		},
	}

	p := New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Normalize(tt.phone, "BR")
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeBrazilianPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NormalizeBrazilianPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
