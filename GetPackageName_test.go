package system

import "testing"

func TestGetPackageName(t *testing.T) {
	tests := []struct {
		name     string
		wantName string
	}{
		{
			name:     "Test",
			wantName: "system",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotName := GetPackageName(); gotName != tt.wantName {
				t.Errorf("GetPackageName() = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}
