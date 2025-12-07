package ctrl

type (
	CtrlPkg struct {
		Blog *BlogCtrl
	}
)

var Ctrl = &CtrlPkg{
	Blog: Blog,
}