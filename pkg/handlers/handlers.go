package handlers

import (
	"net/http"

	"github.com/anandureghu/bookings/pkg/config"
	"github.com/anandureghu/bookings/pkg/models"
	"github.com/anandureghu/bookings/pkg/render"
)

// the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repostory for the new handlers
func NewHandler(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some businness logic
	stringMap := make(map[string]string)
	stringMap["name"] = "Name"
	stringMap["email"] = "Email@email.com"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP
	//send that data to template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
