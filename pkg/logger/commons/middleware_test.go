package commons

import (
	"errors"
	"reflect"
	"testing"
)

func TestSkipExcludeUrisErr(t *testing.T) {
	middleware := SkipExcludeUris([]string{"/test"})

	ctxRequest := NewContext()
	ctxRequest.AddRequestContext(RequestContext{
		Host:      "localhost",
		Method:    "GET",
		Uri:       "/test",
		RemoteIp:  "0.0.0.0",
		UserAgent: "Testing",
		Latency:   "0",
		LatencyH:  "0",
		Status:    0,
		BytesIn:   0,
		BytesOut:  0,
	})

	message, cid, tag, ctx, err := middleware("message", Cid("1234567890"), TagPlatformApp, *ctxRequest)

	if message != "message" {
		t.Errorf("SkipExcludeUris() message = %v, want %v", message, "message")
	}

	if cid != Cid("1234567890") {
		t.Errorf("SkipExcludeUris() cid = %v, want %v", cid, Cid("1234567890"))
	}

	if tag != TagPlatformApp {
		t.Errorf("SkipExcludeUris() tag = %v, want %v", tag, TagPlatformApp)
	}

	if !reflect.DeepEqual(ctx, *ctxRequest) {
		t.Errorf("SkipExcludeUris() ctx = %v, want %v", ctx, *ctxRequest)
	}

	if err.Error() != errors.New("/test should not logged").Error() {
		t.Errorf("SkipExcludeUris() error = %v, want %v", err.Error(), errors.New("/test should not logged").Error())
	}
}

func TestSkipExcludeUrisOk(t *testing.T) {
	middleware := SkipExcludeUris([]string{"/test"})

	ctxRequest := NewContext()
	ctxRequest.AddRequestContext(RequestContext{
		Host:      "localhost",
		Method:    "GET",
		Uri:       "/other",
		RemoteIp:  "0.0.0.0",
		UserAgent: "Testing",
		Latency:   "0",
		LatencyH:  "0",
		Status:    0,
		BytesIn:   0,
		BytesOut:  0,
	})

	message, cid, tag, ctx, err := middleware("message", Cid("1234567890"), TagPlatformApp, *ctxRequest)

	if message != "message" {
		t.Errorf("SkipExcludeUris() message = %v, want %v", message, "message")
	}

	if cid != Cid("1234567890") {
		t.Errorf("SkipExcludeUris() cid = %v, want %v", cid, Cid("1234567890"))
	}

	if tag != TagPlatformApp {
		t.Errorf("SkipExcludeUris() tag = %v, want %v", tag, TagPlatformApp)
	}

	if !reflect.DeepEqual(ctx, *ctxRequest) {
		t.Errorf("SkipExcludeUris() ctx = %v, want %v", ctx, *ctxRequest)
	}

	if err != nil {
		t.Errorf("SkipExcludeUris() error = %v, want %v", err, nil)
	}
}
