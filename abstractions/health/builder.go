package health

type Builder struct {
	components []Indicator
}

func NewHealthIndicator(compList []Indicator) *Builder {
	return &Builder{components: compList}
}

func (builder *Builder) Build() map[string]interface{} {
	root := make(map[string]interface{})
	root["status"] = "up"
	var componentStatus []map[string]interface{}
	for _, component := range builder.components {
		status := component.Health()
		componentStatus = append(componentStatus, status.Build())
		if status.GetStatus() == "down" {
			root["status"] = "down"
		}
	}
	root["components"] = componentStatus
	return root
}
