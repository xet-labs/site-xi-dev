package web

import "sync"

type WebApp struct {
	once sync.Once
}

var App = &WebApp{}
