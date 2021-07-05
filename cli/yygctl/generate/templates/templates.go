package templates

import (
	"sort"
	"sync"
	"yygctl/generate/projects"
)

var projectsTmpMap map[string]*projects.Project
var mutex sync.Mutex

func registerProject(name string, project *projects.Project) {
	mutex.Lock()
	defer mutex.Unlock()

	if projectsTmpMap == nil {
		projectsTmpMap = make(map[string]*projects.Project)
	}
	projectsTmpMap[name] = project
}

func GetProjectByName(name string) *projects.Project {
	return projectsTmpMap[name]
}

func GetProjectList() []string {
	j := 0
	keys := make([]string, len(projectsTmpMap))
	for k := range projectsTmpMap {
		keys[j] = k
		j++
	}
	sort.Strings(keys)
	return keys
}
