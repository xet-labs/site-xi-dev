package constraints

import "regexp"

var (
	unameRE = regexp.MustCompile(`^@[a-zA-Z0-9_]{3,33}$`)
	uidRE   = regexp.MustCompile(`^[0-9]{1,20}$`)
	slugRE  = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,64}$`)
)

func Uname(s string) bool { return unameRE.MatchString(s) }
func UID(s string) bool   { return uidRE.MatchString(s) }
func Slug(s string) bool  { return slugRE.MatchString(s) }
