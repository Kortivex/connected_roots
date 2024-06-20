package commons

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestUnwrap(t *testing.T) {
	errorS1 := NewErrorS(400, "error1", map[string]any{"k1": "v1"}, errors.New("new error1"))
	errorS2 := NewErrorS(400, "error2", map[string]any{"k2": "v2"}, fmt.Errorf("error: %w", errorS1))
	errorS3 := NewErrorS(400, "error3", map[string]any{"k3": "v3"}, fmt.Errorf("error: %w", errorS2))

	var want []ErrorCMap

	e1 := ErrorCMap{}
	e1["status"] = 400
	e1["message"] = "error1"
	e1["details"] = map[string]any{"k1": "v1"}

	e2 := ErrorCMap{}
	e2["status"] = 400
	e2["message"] = "error2"
	e2["details"] = map[string]any{"k2": "v2"}

	e3 := ErrorCMap{}
	e3["status"] = 400
	e3["message"] = "error3"
	e3["details"] = map[string]any{"k3": "v3"}

	want = append(want, e3, e2, e1)
	got := Unwrap(errorS3)

	// println(fmt.Sprintf("%v", want))

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Errorf("Unwrap() = %+v, want %+v", got, want)
	}
}
