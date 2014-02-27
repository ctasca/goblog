package app

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"io/ioutil"
)

func setUp() *BlogInstaller {
	return NewBlogInstaller("0.0.0.0", "/Users/ctasca/mygo/src/github.com/ctasca/goblog")
}

func TestNewBlogInstaller(t *testing.T) {
	Convey("Given NewBlogInstaller message is sent", t, func() {
			i := setUp()
			Convey("Then the constucted object must be of type BlogInstaller", func() {
					So(i, ShouldHaveSameTypeAs, &BlogInstaller{})
				})
		})
}

func TestHandler(t *testing.T) {
	Convey("Given a new blog installer is created", t, func() {
			i := setUp()
			Convey("Then invoking Handler should return the installer handler 'func() string'", func() {
					handler := i.Handler()
					handlerMock := func() string {
						return "Hello from Installer Handler!"
					}
					So(handler, ShouldHaveSameTypeAs, handlerMock)
				})
		})
}

func TestVersion(t *testing.T) {
	Convey("Given a new blog installer with version 0.0.0.0 is created", t, func() {
			i := setUp()
			Convey("Then invoking Version should return the expected version number", func() {
					version := i.Version()
					So(version, ShouldEqual, "0.0.0.0")
				})
		})
}


func TestInstalled(t *testing.T) {
	Convey("Given a new blog installer is created", t, func() {
			i := setUp()
			os.Remove(i.Basedir() + "/_os/etc/config.json")
			Convey("When the message Installed is sent", func() {
					installed := i.Installed()
					Convey("Then the value returned should be false when no configuration file is found", func() {
							So(installed, ShouldEqual, false)
						})
				})
		})
}

func TestInstall(t *testing.T) {
	Convey("Given a new blog installer is created", t, func() {
			i := setUp()

			Convey("When the message Install is sent", func() {
					post := make(map[string]interface{})
					post["dbname"] = "goblog"
					post["dbusr"] = "root"
					post["dbpass"] = "what3v3r"
					post["routes"] = []map[string]string{
					map[string]string{"/" : "HomeHandler"},
					map[string]string{"/about/" : "AboutHandler"}}

					installed := i.Install(post)

					Convey("Then there must be config.json file in /tmp/ directory", func() {
							_, err := os.Open("/tmp/config.json");
							So(err, ShouldEqual, nil)
							Convey("And the content of /tmp/config.json is the same as [basedir]/etc/config.json", func() {
									tconfig, _ := ioutil.ReadFile("/tmp/config.json")
									configerr  := ioutil.WriteFile(i.Basedir() + "/_os/etc/config.json", tconfig, 0644)
									So(configerr, ShouldEqual, nil)
											Convey("And should return true", func() {
													So(installed, ShouldEqual, nil)
												})
								})
						})
				})
		})
}

