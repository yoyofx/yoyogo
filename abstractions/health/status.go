package health

type ComponentStatus struct {
	status  map[string]interface{}
	details map[string]interface{}
}

func newStatus(name string) ComponentStatus {
	cs := ComponentStatus{}
	cs.status = make(map[string]interface{})
	cs.details = make(map[string]interface{})
	if name != "" {
		cs.status["name"] = name
	}
	return cs
}

func Down(name string) ComponentStatus {
	cs := newStatus(name)
	cs.status["status"] = "down"
	return cs
}

func Up(name string) ComponentStatus {
	cs := newStatus(name)
	cs.status["status"] = "up"
	return cs
}

func (cs ComponentStatus) WithDetail(key string, detail interface{}) ComponentStatus {
	cs.details[key] = detail
	return cs
}

func (cs ComponentStatus) GetName() string {
	return cs.status["name"].(string)
}

func (cs ComponentStatus) GetStatus() string {
	return cs.status["status"].(string)
}

func (cs ComponentStatus) SetStatus(status string) {
	cs.status["status"] = status
}

func (cs ComponentStatus) Build() map[string]interface{} {
	cs.status["details"] = cs.details
	return cs.status
}
