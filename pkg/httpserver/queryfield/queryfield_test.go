package queryfield_test

import (
	"testing"

	"github.com/Kortivex/connected_roots/pkg/httpserver/queryfield"
)

func TestDecodeType(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "hello", want: `"hello"`},
		{input: "hello+e", want: `"hello+e"`},
		{input: "hel5", want: `"hel5"`},
		{input: "5", want: "5"},
		{input: "5.45", want: "5.45"},
	}

	for i, tc := range tests {
		got := queryfield.DecodeType(tc.input)
		if got != tc.want {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}

func TestDecodeOperator(t *testing.T) {
	tests := []struct {
		input     string
		wantOp    string
		wantField string
	}{
		{input: "eq:field:subfield", wantOp: "==", wantField: "field:subfield"},
		{input: "field:subfield", wantOp: "==", wantField: "field:subfield"},
		{input: "field", wantOp: "==", wantField: "field"},
		{input: "neq:field:subfield", wantOp: "!=", wantField: "field:subfield"},
		{input: "gt:field:subfield", wantOp: ">", wantField: "field:subfield"},
		{input: "gte:field:subfield", wantOp: ">=", wantField: "field:subfield"},
		{input: "lt:field:subfield", wantOp: "<", wantField: "field:subfield"},
		{input: "lte:field:subfield", wantOp: "<=", wantField: "field:subfield"},
	}

	for i, tc := range tests {
		gotOp, gotField := queryfield.DecodeOperator(tc.input)
		if gotOp != tc.wantOp {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.wantOp, gotOp)
		}

		if gotField != tc.wantField {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.wantField, gotField)
		}
	}
}

func TestDecodeEmbeddedName(t *testing.T) {
	tests := []struct {
		name   string
		prefix string
		want   string
	}{
		{name: "property:field", prefix: "property:", want: "field"},
		{name: "property:field:subfield", prefix: "property:", want: "field.subfield"},
		{name: "field", prefix: "", want: "field"},
		{name: "field:subfield", prefix: "", want: "field.subfield"},
	}

	for i, tc := range tests {
		got := queryfield.DecodeEmbeddedName(tc.name, tc.prefix)
		if got != tc.want {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		prefix    string
		wantName  string
		wantOp    string
		wantValue string
	}{
		{name: "property:field", value: "neq:5", prefix: "property:", wantName: "field", wantOp: "!=", wantValue: "5"},
	}

	for i, tc := range tests {
		gotName, gotOp, gotValue := queryfield.Decode(tc.name, tc.value, tc.prefix)
		if gotName != tc.wantName {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.wantName, gotName)
		}

		if gotOp != tc.wantOp {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.wantOp, gotOp)
		}

		if gotValue != tc.wantValue {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.wantValue, gotValue)
		}
	}
}
