package logger

import (
	"testing"
)

func TestGetLevel(t *testing.T) {
	if got := GetLevel("debug"); got != -1 {
		t.Errorf("GetLevel() = %v, want %v", got, -1)
	}
	if got := GetLevel("info"); got != 0 {
		t.Errorf("GetLevel() = %v, want %v", got, 0)
	}
	if got := GetLevel("warn"); got != 1 {
		t.Errorf("GetLevel() = %v, want %v", got, 1)
	}
	if got := GetLevel("error"); got != 2 {
		t.Errorf("GetLevel() = %v, want %v", got, 2)
	}
	if got := GetLevel("silent"); got != -2 {
		t.Errorf("GetLevel() = %v, want %v", got, -2)
	}
}
