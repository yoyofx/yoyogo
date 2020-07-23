package Mvc

const (
	DefaultMvcTemplate = "v1/{controller}/{action}"
)

type Options struct {
	Template RouteTemplate
}

// NewMvcOptions new mvc options.
func NewMvcOptions() Options {
	return Options{Template: NewRouteTemplate(DefaultMvcTemplate)}
}

// MapRoute map url route to mvc url template.
func (options Options) MapRoute(template string) {
	options.Template = NewRouteTemplate(template)
}
