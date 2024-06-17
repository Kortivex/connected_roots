package commons

import (
	"errors"
	"reflect"
	"testing"
)

func TestErrorS(t *testing.T) {
	errorS := NewErrorS(400, "error", map[string]any{"k": "v"}, errors.New("new error"))

	want := &ErrorS{
		Status:  400,
		Message: "error",
		Details: map[string]any{"k": "v"},
		err:     errors.New("new error"),
	}
	if got := errorS; !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDefaultErrorS(t *testing.T) {
	errorS := NewDefaultErrorS(errors.New("new error"))

	want := &ErrorS{
		Status:  0o0,
		Message: "",
		Details: nil,
		err:     errors.New("new error"),
	}
	if got := errorS; !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestCastingErrorS(t *testing.T) {
	errorS := NewErrorS(400, "error", map[string]any{"k": "v"}, errors.New("new error"))

	err := errorS.(error)

	want := &ErrorS{
		Status:  400,
		Message: "error",
		Details: map[string]any{"k": "v"},
		err:     errors.New("new error"),
	}


	//println(fmt.Sprintf("%+v", want))

	if ok := reflect.DeepEqual(err, want); !ok {
		t.Errorf("got = %v, want %v", err, want)
	}

	if ok := reflect.DeepEqual(err.(ErrorI), want); !ok {
		t.Errorf("got = %v, want %v", err, want)
	}
}

func TestErrorS_Status(t *testing.T) {
	errorS := NewErrorS(400, "error", map[string]any{"k": "v"}, errors.New("new error"))

	want := &ErrorS{
		Status:  400,
		Message: "error",
		Details: map[string]any{"k": "v"},
		err:     errors.New("new error"),
	}
	if got := errorS.Error(); errorS.ErrorStatus() != want.ErrorStatus() {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestErrorS_Message(t *testing.T) {
	errorS := NewErrorS(400, "error", map[string]any{"k": "v"}, errors.New("new error"))

	want := &ErrorS{
		Status:  400,
		Message: "error",
		Details: map[string]any{"k": "v"},
		err:     errors.New("new error"),
	}
	if got := errorS.Error(); errorS.ErrorMessage() != want.ErrorMessage() {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestErrorS_Details(t *testing.T) {
	errorS := NewErrorS(400, "error", map[string]any{"k": "v"}, errors.New("new error"))

	want := &ErrorS{
		Status:  400,
		Message: "error",
		Details: map[string]any{"k": "v"},
		err:     errors.New("new error"),
	}
	if got := errorS.Error(); !reflect.DeepEqual(errorS.ErrorDetails(), want.ErrorDetails()) {
		t.Errorf("got = %v, want %v", got, want)
	}
}
