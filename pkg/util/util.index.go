package util

import (
	"xi/pkg/util/crypt"
	"xi/pkg/util/file"
	"xi/pkg/util/maps"
	"xi/pkg/util/minify"
	"xi/pkg/util/misc"
	"xi/pkg/util/str"
	"xi/pkg/util/url"
)

type UtilLib struct {
	Crypt  crypt.CryptLib
	File   file.FileLib
	Map    maps.MapsLib
	Minify minify.MinifyLib
	Misc   misc.MiscLib
	Str    str.StrLib
	Url    url.UrlLib
}

var Util = &UtilLib{
	Crypt:  crypt.CryptLib{},
	File:   file.FileLib{},
	Map:    maps.MapsLib{},
	Misc:   misc.MiscLib{},
	Minify: minify.MinifyLib{},
	Str:    str.StrLib{},
	Url:    url.UrlLib{},
}

// expose shortcuts
var (
	Crypt  = &Util.Crypt
	File   = &Util.File
	Map    = &Util.Map
	Minify = &Util.Minify
	Misc   = &Util.Misc
	Str    = &Util.Str
	Url    = &Util.Url
)
