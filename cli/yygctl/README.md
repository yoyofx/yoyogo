# yygctl 
## install
```bash
go install github.com/yoyofx/yoyogo/cli/yygctl
```

## local install
```bash
cd yoyogo/cli/yygctl
go install
```

# Installation location:
$GOPATH

add $GOPATH to $PATH Environment variable

# Commands
There are commands working with application root folder
## new 
```bash
yygctl new <TEMPLATE> [-l|--list] [-n <PROJECTNAME>] [-p <TARGETDIR>]
```
### --list
list all templates
#### TEMPLATE LIST
console / webapi / mvc / grpc / xxl-job

### -n 
generate folder by project name <PROJECTNAME>

### -p <TARGETDIR>
output files to target directory. 

## such as 
```bash
yygctl new console -n demo -p /Projects
```

## add (Not realized)
add code snippet to the file, filepath was for default settings.
```bash
yygctl add <SNIPPET> [-l|--list] [-f|--file <filepath>]
```
#### SNIPPET LIST
dockerfile / config / controller / job-handler / hostservice / startup / web-middleware / web-filter / grpc-interceptor

## build
build current working directory
```bash
yygctl build
```

## run
running current working directory app
```bash
yygctl run
```

## version
display yoyogo version
```bash
yygctl version
```

