package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func checkPassword(login, password string) bool {
	if login == "azaru" && password == "123"{
		return true
	}
	return false
}

func indexHandler(rnd render.Render) {
	rnd.HTML(200,"index",nil)
}

func mainHandler(rnd render.Render, r *http.Request ) {

	rnd.HTML(200,"main",nil)
}

func loginHandler(rnd render.Render, r *http.Request ) {
	login := r.FormValue("login")
	pass := r.FormValue("password")
	if checkPassword(login,pass){
		rnd.Redirect("/main")
	}else{
		rnd.Redirect("/error_Login")
	}
}


func error_LoginHandler(rnd render.Render) {
	rnd.HTML(200,"error","Didn't find user")
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                         // Specify what path to load the templates from.
		Layout:     "layout",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"},          // Specify extensions to load for templates.
		//Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8",                             // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                                // Output human readable JSON
	}))

	staticOptions := martini.StaticOptions{ Prefix : "assets"}
	m.Use(martini.Static("assets", staticOptions))

	m.Get("/",indexHandler)
	m.Get("/main",mainHandler)
	m.Post("/login",loginHandler)
	m.Get("/error_Login",error_LoginHandler)
	m.Run()
}







