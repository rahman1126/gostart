package hmac

import "testing"

func TestHex(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hex(); got != tt.want {
				t.Errorf("Hex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkHex(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Hex()
	}
}

func BenchmarkBase64(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Base64()
	}
}
