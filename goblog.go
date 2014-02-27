package main

import (
	"github.com/ctasca/goblog/app"
	"fmt"
	"strconv"
	"os"
)

func main() {
	wd, err := os.Getwd()

	if(err != nil) {
		panic("Cannot get OS working directory")
	}

	installer := app.NewBlogInstaller("0.0.0.1", wd);
	goblog := app.NewApp(installer)

	status, error := goblog.Run()

	if error != nil {
		fmt.Println("App run error: " + error.Error())
	} else {
		fmt.Println("App run status code: " + strconv.Itoa(status))
	}
}
