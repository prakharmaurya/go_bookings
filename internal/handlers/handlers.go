package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prakharmaurya/go_bookings/internal/config"
	"github.com/prakharmaurya/go_bookings/internal/models"
	"github.com/prakharmaurya/go_bookings/internal/render"
)

var Repo *Repository

type Repository struct {
	app *config.AppConfig
}

func NewRepositor(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.app.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(rw, r, "home.page.tmpl", &models.TemplateData{IsAuthenticated: 0})
}

func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["greet"] = "sdvndfvkjndf"
	stringMap["remote_ip"] = m.app.Session.GetString(r.Context(), "remote_ip")
	render.RenderTemplate(rw, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservation(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "make-reservation.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Majors(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "majors.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Generals(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "generals.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Availability(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "search-availability.page.tmpl", &models.TemplateData{})
}
func (m *Repository) PostAvailability(rw http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	rw.Write([]byte("start date is " + start + " end date " + end))

}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(rw http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}
	out, err := json.MarshalIndent(resp, "", "	")

	if err != nil {
		fmt.Println("failed to masrshel", err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(out)
}
func (m *Repository) Contact(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "contact.page.tmpl", &models.TemplateData{})
}
