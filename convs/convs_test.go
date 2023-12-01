package convs

import (
	"testing"
)

func TestStringToType(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    any
		wantErr bool
	}{
		{
			name:  "convert string to int",
			input: "123",
			want:  int(123),
		},
		{
			name:  "convert string to int32",
			input: "12345",
			want:  int32(12345),
		},
		{
			name:  "convert string to int64",
			input: "123456789012345",
			want:  int64(123456789012345),
		},
		{
			name:  "convert string to float32",
			input: "123.45",
			want:  float32(123.45),
		},
		{
			name:  "convert string to float64",
			input: "1234567890.123456",
			want:  float64(1234567890.123456),
		},
		{
			name:  "convert string to bool (true)",
			input: "true",
			want:  true,
		},
		{
			name:  "convert string to bool (false)",
			input: "false",
			want:  false,
		},
		{
			name:    "invalid input for int",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got any
			var err error

			switch v := tt.want.(type) {
			case int:
				converter := NewConverter[int]()
				got, err = converter.StringToType(tt.input)
			case int32:
				converter := NewConverter[int32]()
				got, err = converter.StringToType(tt.input)
			case int64:
				converter := NewConverter[int64]()
				got, err = converter.StringToType(tt.input)
			case float32:
				converter := NewConverter[float32]()
				got, err = converter.StringToType(tt.input)
			case float64:
				converter := NewConverter[float64]()
				got, err = converter.StringToType(tt.input)
			case bool:
				converter := NewConverter[bool]()
				got, err = converter.StringToType(tt.input)
			case string:
				converter := NewConverter[string]()
				got, err = converter.StringToType(tt.input)
			default:
				t.Fatalf("Unsupported type: %T", v)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("StringToType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringToType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:  "convert int to string",
			input: int(123),
			want:  "123",
		},
		{
			name:  "convert int32 to string",
			input: int32(12345),
			want:  "12345",
		},
		{
			name:  "convert int64 to string",
			input: int64(123456789012345),
			want:  "123456789012345",
		},
		{
			name:  "convert float32 to string",
			input: float32(123.45),
			want:  "123.45",
		},
		{
			name:  "convert float64 to string",
			input: float64(1234567890.123456),
			want:  "1234567890.123456",
		},
		{
			name:  "convert bool (true) to string",
			input: true,
			want:  "true",
		},
		{
			name:  "convert bool (false) to string",
			input: false,
			want:  "false",
		},
		{
			name:    "invalid input for int",
			input:   "abc",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter := NewConverter[any]()
			got, err := converter.ToString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ToString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
