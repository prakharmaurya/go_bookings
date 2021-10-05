package handlers

import (
	"net/http"

	"github.com/prakharmaurya/go_bookings/pkg/config"
	"github.com/prakharmaurya/go_bookings/pkg/models"
	"github.com/prakharmaurya/go_bookings/pkg/render"
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
	render.RenderTemplate(rw, "home.page.tmpl", &models.TemplateData{IsAuthenticated: 0})
}

func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["greet"] = "sdvndfvkjndf"
	stringMap["remote_ip"] = m.app.Session.GetString(r.Context(), "remote_ip")
	render.RenderTemplate(rw, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservation(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "make-reservation.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Majors(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "majors.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Generals(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "generals.page.tmpl", &models.TemplateData{})
}
