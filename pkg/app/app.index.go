package app

import "xi/pkg/app/handler"

type AppPkg struct {
	Err *handler.AppErr
}

var App = &AppPkg{
	Err: handler.Err,
}

// shortcuts
var Err = App.Err
