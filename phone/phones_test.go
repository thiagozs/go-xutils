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

func TestIsValid(t *testing.T) {
	tests := []struct {
		name    string
		phone   string
		want    bool
		wantErr bool
	}{
		{
			name:    "Valid mobile phone number with country code",
			phone:   "+55 11 98765-4321",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Valid mobile phone number without country code",
			phone:   "11 98765-4321",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Valid mobile phone number with country code parantheses",
			phone:   "+55 (11) 98765-4321",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Valid mobile phone number without country code parantheses and dash",
			phone:   "(11) 98765-4321",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Invalid phone number",
			phone:   "00000000",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Invalid phone number with country code",
			phone:   "+55 00 000000000",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Invalid phone number with country code and parantheses",
			phone:   "+55 (00) 000000000",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Invalid phone number with country code and parantheses and dash",
			phone:   "+55 (00) 00000-0000",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Valid a residential phone number with country code",
			phone:   "+55 19 3611-4444",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Valid a residential phone number with country code and parantheses and dash",
			phone:   "+55 (19) 3611-4444",
			want:    true,
			wantErr: false,
		},
	}

	p := New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := p.IsValid(tt.phone, "BR")
			if got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		name  string
		limit int
	}{
		{
			name:  "Generate 10 random phone numbers",
			limit: 10,
		},
		{
			name:  "Generate 100 random phone numbers",
			limit: 100,
		},
		{
			name:  "Generate 1000 random phone numbers",
			limit: 1000,
		},
		{
			name:  "Generate 10000 random phone numbers",
			limit: 10000,
		},
	}

	p := New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := p.Generate(tt.limit)
			if len(got) != tt.limit {
				t.Errorf("Generate() = %v, want %v", len(got), tt.limit)
			}
		})
	}
}

func TestGenMobile(t *testing.T) {
	tests := []struct {
		name    string
		country string
		limit   int
	}{
		{
			name:    "Generate 10 random mobile phone numbers",
			country: "BR",
			limit:   10,
		},
		{
			name:    "Generate 100 random mobile phone numbers",
			country: "BR",
			limit:   100,
		},
		{
			name:    "Generate 1000 random mobile phone numbers",
			country: "BR",
			limit:   1000,
		},
		{
			name:    "Generate 10000 random mobile phone numbers",
			country: "BR",
			limit:   10000,
		},
	}

	p := New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := p.GenMobile(tt.country, tt.limit)
			if len(got) != tt.limit {
				t.Errorf("GenMobile() = %v, want %v", len(got), tt.limit)
			}
		})
	}
}

func TestGenLandline(t *testing.T) {
	tests := []struct {
		name    string
		country string
		limit   int
	}{
		{
			name:    "Generate 10 random landline phone numbers",
			country: "BR",
			limit:   10,
		},
		{
			name:    "Generate 100 random landline phone numbers",
			country: "BR",
			limit:   100,
		},
		{
			name:    "Generate 1000 random landline phone numbers",
			country: "BR",
			limit:   1000,
		},
		{
			name:    "Generate 10000 random landline phone numbers",
			country: "BR",
			limit:   10000,
		},
	}

	p := New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := p.GenLandline(tt.country, tt.limit)
			if len(got) != tt.limit {
				t.Errorf("GenLandline() = %v, want %v", len(got), tt.limit)
			}
		})
	}
}

func TestGenMobileWithMask(t *testing.T) {
	tests := []struct {
		name  string
		limit int
	}{
		{
			name:  "Generate 10 random mobile phone numbers with mask",
			limit: 10,
		},
		{
			name:  "Generate 100 random mobile phone numbers with mask",
			limit: 100,
		},
		{
			name:  "Generate 1000 random mobile phone numbers with mask",
			limit: 1000,
		},
		{
			name:  "Generate 10000 random mobile phone numbers with mask",
			limit: 10000,
		},
	}

	p := New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := p.GenMobileWithMask(tt.limit)
			if len(got) != tt.limit {
				t.Errorf("GenMobileWithMask() = %v, want %v", len(got), tt.limit)
			}
		})
	}
}

func BenchmarkNormalize(b *testing.B) {
	p := New()

	for i := 0; i < b.N; i++ {
		p.Normalize("+55 11 98765-4321", "BR")
	}
}

func BenchmarkIsValid(b *testing.B) {
	p := New()

	for i := 0; i < b.N; i++ {
		p.IsValid("+55 11 98765-4321", "BR")
	}
}

func BenchmarkGenerate(b *testing.B) {
	p := New()

	for i := 0; i < b.N; i++ {
		p.Generate(1000)
	}
}
