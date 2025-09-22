package util

import (
	"xi/pkg/lib/util/crypt"
	"xi/pkg/lib/util/file"
	"xi/pkg/lib/util/maps"
	"xi/pkg/lib/util/minify"
	"xi/pkg/lib/util/misc"
	"xi/pkg/lib/util/str"
	"xi/pkg/lib/util/url"
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
