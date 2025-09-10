package util

import (
	"xi/pkg/util/file"
	"xi/pkg/util/maps"
	"xi/pkg/util/minify"
	"xi/pkg/util/misc"
	"xi/pkg/util/str"
	"xi/pkg/util/url"
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
