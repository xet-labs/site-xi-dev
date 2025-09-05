package util

import (
	"xi/app/lib/util/file"
	"xi/app/lib/util/maps"
	"xi/app/lib/util/minify"
	// "xi/app/lib/util/misc"
	"xi/app/lib/util/str"
	"xi/app/lib/util/url"
)

type UtilLib struct {
	File   file.FileLib
	Map    maps.MapsLib
	Minify minify.MinifyLib
	Str    str.StrLib
	Url    url.UrlLib
}

var Util = &UtilLib{
	File:   file.FileLib{},
	Map:    maps.MapsLib{},
	Minify: minify.MinifyLib{},
	Str:    str.StrLib{},
	Url:    url.UrlLib{},
}

// expose shortcuts
var (
	File   = &Util.File
	Map    = &Util.Map
	Minify = &Util.Minify
	// Misc = misc
	Str    = &Util.Str
	Url    = &Util.Url
)

// var Misc = misc
