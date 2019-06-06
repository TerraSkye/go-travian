package main

import (
	"fmt"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"

	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// parseTemplate applies a given file to the body of the base template.
func parseTemplate(filename string) *appTemplate {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("html/template", html.Minify)

	tmpl := template.Must(template.ParseFiles("templates/base.html"))

	// Put the named file into a template called "body"
	path := filepath.Join("templates", filename)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("could not read template: %v", err))
	}

	mb, err := m.Bytes("html/template", b)
	if err != nil {
		panic(fmt.Errorf("minify error for: %v", err))
	}

	template.Must(tmpl.New("body").Parse(string(mb)))

	return &appTemplate{tmpl.Lookup("base.html")}
}

// appTemplate is a user login-aware wrapper for a html/template.
type appTemplate struct {
	t *template.Template
}

// Execute writes the template using the provided data, adding login and user
// information to the base template.
func (tmpl *appTemplate) Execute(w http.ResponseWriter, r *http.Request, data interface{}) error {

	d := struct {
		Data        interface{}
		AuthEnabled bool
		//Profile     *Profile
		//LoginURL    string
		//LogoutURL   string
	}{
		Data:        data,
		AuthEnabled: false,
		//AuthEnabled: bookshelf.OAuthConfig != nil,
		//LoginURL:    "/login?redirect=" + r.URL.RequestURI(),
		//LogoutURL:   "/logout?redirect=" + r.URL.RequestURI(),
	}

	//if d.AuthEnabled {
	//	// Ignore any errors.
	//	d.Profile = profileFromSession(r)
	//}

	if err := tmpl.t.Execute(w, d); err != nil {
		return error(err)
	}
	return nil
}
