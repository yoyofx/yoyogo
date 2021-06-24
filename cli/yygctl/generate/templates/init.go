package templates

import (
	"github.com/yoyofx/yoyogo/cli/yygctl/generate/projects"
	"sync"
)

var projectsTmpMap map[string]*projects.Project
var mutex sync.Mutex

func init() {
	registerProject("demo", demoProject)

}

func registerProject(name string, project *projects.Project) {
	mutex.Lock()
	defer mutex.Unlock()

	if projectsTmpMap == nil {
		projectsTmpMap = make(map[string]*projects.Project)
	}
	projectsTmpMap[name] = project
}

func GetProject(name string) *projects.Project {
	return projectsTmpMap[name]
}
