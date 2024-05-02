package auth

import "testing"

func TestEncodePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"vienas",
			"pass",
			"ce4ed2c457e6b5c1d1cae33c6207f3009c9f50de",
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomHex(20)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodePassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}
