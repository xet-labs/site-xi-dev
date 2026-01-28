package web

import "xi/pkg"

func (s *WebApp) Init() {
	s.once.Do(s.InitForce)
}

func (s *WebApp) InitForce() {
	pkg.Util.Minify.Init()
}
