package app

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Installer interface {
	Basedir () string
	Etcdir () string
	HttpHandlers()
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
	etcdir   string
	version   string
}
// Creates a new blog application
func NewApp(installer Installer) *App {
	application              := new(App)
	application.installer    = installer
	application.basedir      = application.installer.Basedir()
	application.etcdir       = application.installer.Etcdir()
	main, err                := ioutil.ReadFile(application.coreJsonPath());
	var dat map[string]string

	if err != nil {
		panic("Blog application could not be created " + err.Error())
	} else {
		json.Unmarshal(main, &dat)
		application.initialize(main, dat["name"], dat["domain"], dat["port"])
	}
	return application
}

// Invoked in main, runs the application
// If the application is not installed, serves installer handler
func (application *App) Run() (int, error) {
	if !application.installer.Installed() {
		application.installer.HttpHandlers()
		return 3, http.ListenAndServe(application.domain+":"+application.port, nil)
	}
	return 1, http.ListenAndServe(application.domain+":"+application.port, nil)
}

// initialise application main, name, domain and port members
func (application *App) initialize(main []byte, name string, domain string, port string) {
	application.main      = main
	application.name      = name
	application.domain    = domain
	application.port      = port
}

func (application *App) coreJsonPath() string {
	return application.installer.Basedir() + application.etcdir + "/core.json"
}
