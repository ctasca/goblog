package app

import (
	"net/http"
	"github.com/codegangsta/martini"
	"encoding/json"
	"io/ioutil"
)

type Installer interface {
	Basedir () string
	Handler () interface {}
	Version ()  string
	Installed () bool
	Install(map[string]interface{}) error
}

type App struct {
	main      []byte
	installer Installer
	name      string
	domain    string
	port      string
	basedir   string
	version   string
}
// Creates a new blog application
func NewApp(installer Installer) *App {
	application              := new(App)
	application.installer    = installer
	application.basedir      = application.installer.Basedir()
	main, err                := ioutil.ReadFile(application.installer.Basedir() + "/_os/etc/appconfig.json");
	var dat map[string]string

	if err != nil {
		panic("Blog application could not be created " + err.Error())
	} else {
		json.Unmarshal(main, &dat)
		application.initialize(main, dat["name"], dat["domain"], dat["port"])
	}
	return application
}

func (application *App) Run() (int, error) {
	m := martini.Classic()
	if !application.installer.Installed() {
		m.Get("/", application.installer.Handler())
		return 3, http.ListenAndServe(application.domain+":"+application.port, m)
	}
	return 1, http.ListenAndServe(application.domain+":"+application.port, m)
}

// initialise application main, name, domain and port members
func (application *App) initialize(main []byte, name string, domain string, port string) {
	application.main      = main
	application.name      = name
	application.domain    = domain
	application.port      = port
}
