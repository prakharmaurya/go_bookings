package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/prakharmaurya/go_bookings/internal/config"
	"github.com/prakharmaurya/go_bookings/internal/forms"
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
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(rw, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repository) PostReservation(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")

	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(rw, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.app.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(rw, r, "/reservation-summary", http.StatusSeeOther)
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

func (m *Repository) ReservationSummary(rw http.ResponseWriter, r *http.Request) {
	reservation, ok := m.app.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get item from session")
		m.app.Session.Put(r.Context(), "error", "cant get reservation")
		http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
		return
	}
	m.app.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderTemplate(rw, r, "reservation-summary.page.tmpl", &models.TemplateData{Data: data})
}
