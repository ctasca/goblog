package app

type Pager interface {
	Save () error
}

type PageLoader interface {
	LoadPage ()
}

type Page struct {
	Title string
	Head string
	Body string
}

func NewPage(title string) *Page {
	page                 := new(Page)
	page.Title           = title
	return page
}
