package web

import (
	"errors"
	"html"
	"html/template"
	"net/url"
	"strings"
	"xi/internal/app/lib/util"
	"xi/internal/app/lib/web/htmlfn"

	"github.com/rs/zerolog/log"
)

var HtmlFn = template.FuncMap{
	"formatTime":   htmlfn.FormatTime,
	"metaGen":      htmlfn.Meta.Gen,
	"isSlice":      htmlfn.IsSlice,
	// "len":          htmlfn.Len,
	"linkCss":      htmlfn.LinkCss,
	"linkCssSlice": htmlfn.LinkCssSlice,
	"linkJs":       htmlfn.LinkJs,
	"linkJsSlice":  htmlfn.LinkJsSlice,
	"linkLib":      htmlfn.LinkLib,
	"linkLibSlice": htmlfn.LinkLibSlice,
	"slice":        htmlfn.Slice,

	"htmlEscape":   html.EscapeString,
	"join":         strings.Join,
	"urlEscape":    url.QueryEscape,
}

func (v *WebLib) NewTmpl(name, ext string, dirs ...string) *template.Template {
	if name == "" { name = "main"}

	files, err := util.File.GetWithExt(ext, dirs...)
	if err != nil {
		log.Error().Caller().Err(err).Str("cli", name).Str("template-dir", strings.Join(dirs, ", ")).
			Msg("web.NewTmpl: couldnt get template files")
	}
	if len(files) == 0 {
		log.Fatal().Caller().Err(errors.New("web.NewTmpl: no template files found")).
			Str("cli", name).Str("template-dir", strings.Join(dirs, ", ")).
			Msg("web.NewTmpl: couldnt get template files")
	}

	tcli := template.Must(template.New(name).
		Funcs(HtmlFn).
		// Funcs(HtmlFuncs).
		// Funcs(timeutil.Funcs).
		ParseFiles(files...),
	)

	// Store instance globally so it can be used alter by other functions for rendering pages
	if v.Tcli == nil {
		v.Tcli = tcli
		rawTcli, err := tcli.Clone()
		if err != nil {
			log.Error().Caller().Err(err).Str("cli", name)
		}
		v.RawTcli = rawTcli
	}
	return tcli
}
