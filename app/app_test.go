package app

import (
	 "testing"
	 . "github.com/smartystreets/goconvey/convey"
)
/*
func TestItCollectRoutes(t *testing.T) {
	Convey("Given a new blog application is created", t, func() {
			a := NewApp();

			Convey("When the message CollectRoutes is sent", func() {
					routes := a.CollectRoutes()

					Convey("The value returned should be a map of routes with handlers", func() {
							expected := make(map[string]string)
							expected["/"] = "HomeHandler"
							expected["/articles/"] = "ArticlesHandler"
							So(routes, ShouldEqual, expected)
						})
				})
		})
}
*/

func TestNewApp (t *testing.T) {
	Convey("Given NewApp message is sent", t, func() {
			installer := NewBlogInstaller("0.0.0.1", "/Users/ctasca/mygo/src/github.com/ctasca/goblog")
			a := NewApp(installer);
			Convey("The created app must be of type App", func() {
					So(a, ShouldHaveSameTypeAs, &App{})
				})
	})
}
