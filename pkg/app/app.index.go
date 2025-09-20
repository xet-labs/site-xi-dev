package app

import "xi/pkg/app/err"

type AppPkg struct{
	Err *err.AppErr
}

var App = &AppPkg{
	Err: err.Err,
}

// shortcuts
var Err = App.Err