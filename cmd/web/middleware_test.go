package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mh myHandler
	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("Type of returned Value dosen't match, but it is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var mh myHandler
	h := SessionLoad(&mh)

	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("Type of returned Value dosen't match, but it is %T", v))
	}
}
func TestWriteToConsole(t *testing.T) {
	var mh myHandler
	h := writeToConsole(&mh)

	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("Type of returned Value dosen't match, but it is %T", v))
	}
}
