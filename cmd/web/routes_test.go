package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/prakharmaurya/go_bookings/internal/config"
)

func TestRoute(t *testing.T) {
	var app config.AppConfig

	h := routes(&app)

	switch v := h.(type) {
	case *chi.Mux:
	default:
		t.Error(fmt.Sprintf("Type of returned Value dosen't match, but it is %T", v))
	}
}
