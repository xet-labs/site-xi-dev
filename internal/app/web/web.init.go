package web

import "xi/pkg"

func (s *WebApp) Init() {
	pkg.Util.Minify.Init()
}
