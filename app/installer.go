package app

import (
	"os"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

type BlogInstaller struct {
	version string
	basedir string
	etcdir  string
}

func NewBlogInstaller(version string, basedir string, etcdir string) *BlogInstaller {
	installer := new(BlogInstaller)
	installer.version = version
	installer.basedir = basedir
	installer.etcdir = etcdir
	return installer
}

func (installer *BlogInstaller) Handler() interface{} {
	return func() string {
		return "Hello from Installer Handler!"
	}
}

func (installer *BlogInstaller) Version() string {
	return installer.version
}

func (installer *BlogInstaller) Basedir() string {
	return installer.basedir
}

func (installer *BlogInstaller) Etcdir() string {
	return installer.etcdir
}

func (installer *BlogInstaller) Installed() bool {
	_, err := os.Open(installer.Basedir() + installer.Etcdir() + "/config.json");

	if err != nil {
		return false
	}

	return true
}

func (installer *BlogInstaller) Install(dat map[string]interface{}) error {
	err := installer.writeTempJsonConfigFile(dat)
	if err == nil {
		installer.copyTempJsonConfigFile()
	}
	return err
}

func (installer *BlogInstaller) writeTempJsonConfigFile(dat map[string]interface{}) error {
	f, err := os.Create("/tmp/config.json")
	d2, err := json.Marshal(dat)
	defer f.Close()
	i, err := f.Write(d2)
	f.Sync()
	if err != nil {
		panic("Could not create /tmp/config.json: " + strconv.Itoa(i) + err.Error() )
	}
	return err
}

func (installer *BlogInstaller) copyTempJsonConfigFile() error {
	tconfig, temperr := ioutil.ReadFile("/tmp/config.json")

	if temperr != nil {
		panic("Could not read temporary config.json file " + temperr.Error())
	}

	return ioutil.WriteFile(installer.Basedir() + installer.Etcdir()  + "/config.json", tconfig, 0644)
}

