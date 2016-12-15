package main

import (
	"image/png"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testGet(uri string) (rr *httptest.ResponseRecorder) {
	req := httptest.NewRequest("GET", uri, nil)
	rr = httptest.NewRecorder()
	handler := http.HandlerFunc(qrHandler)
	handler.ServeHTTP(rr, req)
	return
}

// Tests that qrHandler returns Bad Request if data param is missing
func TestMissingData(t *testing.T) {
	rr := testGet("/")
	if rr.Code != 400 {
		t.Error("handler didn't return Bad Request on missing data")
	}
}

// Tests that qrHandler returns Bad Request if size is not an integer value
func TestInvalidSize(t *testing.T) {
	rr := testGet("/?data=test&size=large")
	if rr.Code != 400 {
		t.Error("handler didn't return Bad Request on invalid size")
	}
}

// Tests that qrHandler returns Bad Request if size is too large
func TestTooLarge(t *testing.T) {
	rr := testGet("/?data=test&size=8192")
	if rr.Code != 400 {
		t.Error("handler didn't return Bad Request on image size too large")
	}
}

// Tests that qrHandler generates and returns the QR code image as PNG when
// all the params are set
func TestGeneratesImage(t *testing.T) {
	rr := testGet("/?data=test&size=100")
	if rr.Code != 200 {
		t.Error("handler didn't return status code 200 OK")
	}

	img, err := png.Decode(rr.Body)
	if err != nil {
		t.Errorf("handler didn't return a PNG image: %s", err.Error())
	}

	size := img.Bounds().Max
	if size.X != 100 || size.Y != 100 {
		t.Errorf("expected PNG of size (100,100), got size %s instead", size)
	}
}
