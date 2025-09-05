package util

import (
	"xi/internal/app/lib/util/file"
	"xi/internal/app/lib/util/maps"
	"xi/internal/app/lib/util/minify"
	"xi/internal/app/lib/util/misc"
	"xi/internal/app/lib/util/str"
	"xi/internal/app/lib/util/url"
)

type UtilLib struct {
	File   file.FileLib
	Map    maps.MapsLib
	Minify minify.MinifyLib
	Misc   misc.MiscLib
	Str    str.StrLib
	Url    url.UrlLib
}

var Util = &UtilLib{
	File:   file.FileLib{},
	Map:    maps.MapsLib{},
	Misc:   misc.MiscLib{},
	Minify: minify.MinifyLib{},
	Str:    str.StrLib{},
	Url:    url.UrlLib{},
}

// expose shortcuts
var (
	File   = &Util.File
	Map    = &Util.Map
	Minify = &Util.Minify
	Misc   = &Util.Misc
	Str    = &Util.Str
	Url    = &Util.Url
)
