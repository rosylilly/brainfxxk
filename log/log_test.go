package log

import "testing"

func TestLogParse(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Do Log Parse"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogParse()
		})
	}
}
