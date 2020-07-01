package Mvc

const (
	DefalueMvcTemplate = "/{controller}/{action}"
)

type Options struct {
	Template string
}

// NewMvcOptions new mvc options.
func NewMvcOptions() Options {
	return Options{Template: DefalueMvcTemplate}
}

// MapRoute map url route to mvc url template.
func (options Options) MapRoute(templte string) {
	options.Template = templte
}
