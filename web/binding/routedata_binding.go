package binding

type routeDataBinding struct{}

func (routeDataBinding) Name() string {
	return "path"
}

func (routeDataBinding) BindUri(m map[string][]string, obj interface{}) error {
	if err := mapPath(obj, m); err != nil {
		return err
	}
	return validate(obj)
}
