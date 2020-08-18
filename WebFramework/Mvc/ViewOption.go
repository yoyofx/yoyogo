package Mvc

import "html/template"

type ViewOption struct {
	Files   []string
	Pattern string
	FuncMap template.FuncMap
}
