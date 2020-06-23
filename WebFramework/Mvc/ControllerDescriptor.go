package Mvc

type ControllerDescriptor struct {
	ControllerName string
	ControllerType interface{} // ctor func of controller
}

func NewControllerDescriptor(name string, controllerType interface{}) ControllerDescriptor {
	return ControllerDescriptor{name, controllerType}
}
