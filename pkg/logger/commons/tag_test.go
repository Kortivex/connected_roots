package commons

import "testing"

func TestNewTag(t *testing.T) {
	if got := NewTag("platform/testing"); got != Tag("platform/testing") {
		t.Errorf("NewTag() = %v, want %v", got, Tag("platform/testing"))
	}

	newTagComposed := NewTag("platform")
	if got := newTagComposed.Add("testing"); got != Tag("platform/testing") {
		t.Errorf("NewTag() = %v, want %v", got, Tag("platform/testing"))
	}
}
