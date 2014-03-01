package app

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func setUpPager() *Page {
	return NewPage("Title")
}

func TestNewPage(t *testing.T) {
	Convey("Given NewPage message is sent", t, func() {
			pager := setUpPager()
			Convey("Then the constucted object must be of type Page", func() {
					So(pager, ShouldHaveSameTypeAs, &Page{})
				})
		})
}
