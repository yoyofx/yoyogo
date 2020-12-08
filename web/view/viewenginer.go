package view

type IViewEngine interface {
	ViewHtml(viewName string, viewDataset ...interface{}) (string, error)
	AddIncludeTmpl(viewName string)
	SetTemplatePath(option *Option)
}
