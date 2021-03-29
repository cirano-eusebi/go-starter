module github.com/cirano-eusebi/go-starter/src/3-create-module/hello

go 1.16

replace github.com/cirano-eusebi/go-starter/src/3-create-module/greetings => ../greetings

require github.com/cirano-eusebi/go-starter/src/3-create-module/greetings v0.0.0-00010101000000-000000000000
