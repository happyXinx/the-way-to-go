自从1.17开始， go get将不再进行编译包，仅仅用来添加、更新或者删除go.mod文件中的依赖，

在当前module中安装依赖，go install, 不带版本号

在gopath中安装依赖，go install,带版本号


go 1.18 
* install: compile and install packages and dependencies
* get: add dependencies to current module and install them

