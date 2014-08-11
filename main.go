package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"labix.org/v2/mgo"
	"crypto/rand"
	"fmt"
	"Projidea/models"
	
)

func generateId() string{
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x",b)
}

func checkPassword(r *http.Request, login, password string) bool {
	users := []models.UserDocument{}
	usersCollection.Find(nil).All(&users)
    for _,user := range users{
		if user.Name == login && user.Password == password{
			return true
		}
	}

	return false
}

func indexHandler(rnd render.Render) {
	rnd.HTML(200,"login",nil)
}

func mainHandler(rnd render.Render, r *http.Request ) {
	rnd.HTML(200, "main", nil)
}

func loginHandler(rnd render.Render, r *http.Request ) {
	login := r.FormValue("login")
	pass := r.FormValue("password")
	if checkPassword(nil,login,pass){
		rnd.Redirect("/main")
	}else{
		rnd.Redirect("/error_Login")
	}
}


func error_LoginHandler(rnd render.Render) {
	rnd.HTML(200,"error","Didn't find user")
}

func regestrationHandler(rnd render.Render, r *http.Request ) {
	id := generateId()
	name := r.FormValue("login_reg")
	password := r.FormValue("password_reg")
	password_2 := r.FormValue("password_reg_2")
	if password == password_2 {
		usersCollection.Insert(&models.UserDocument{id, name, password})
	}else{
		rnd.HTML(200,"login","Passwords don't coincide!!!")
	}
	rnd.Redirect("/main")

}



var usersCollection *mgo.Collection

func main() {

	session, err := mgo.Dial("localhost")
	if err != nil{
		panic(err)
	}

	usersCollection = session.DB("projidea").C("users")

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
	m.Post("/regestration", regestrationHandler)
	m.Run()
}









