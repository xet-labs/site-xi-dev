package server

import "xi/pkg/lib"

func (s *ServerApp) Init() {
	lib.Util.Minify.Init()
}
