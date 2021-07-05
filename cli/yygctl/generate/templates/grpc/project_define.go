package grpc

import "yygctl/generate/projects"

var Project = projects.NewEmptyProject("grpc", "Grpc Application").With(func(root *projects.ProjectItem) {
	clientDir := root.AddDir("client")
	clientDir.AddFileWithContent("api.go", Client_Api_Tel)
	clientDir.AddFileWithContent("clientservice.go", Client_Service_Tel)
	clientDir.AddFileWithContent("main.go", Client_Main_Tel)
	protoDir := root.AddDir("proto")
	protoDir.AddFileWithContent("helloworld.proto", Hello_World_Tel)
	protoDir.AddDir("helloworld").AddFileWithContent("helloworld.pd.go", Hello_World_PD_TEL)
	serviceDir := root.AddDir("services")
	serviceDir.AddFileWithContent("demo.go", Demo_Tel)
	serviceDir.AddFileWithContent("greeterservice.go", Greeter_Server_Tel)
	root.AddFileWithContent("config_dev.yml", Config_Tel)
	root.AddFileWithContent("go.mod", Mod_Tel)
	root.AddFileWithContent("main.go", Main_Tel)
})
