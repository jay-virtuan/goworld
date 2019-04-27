## create go project

jasintha@jasintha-XPS:~/repo/go$ mkdir hello.io

jasintha@jasintha-XPS:~/repo/go$ cd hello.io

jasintha@jasintha-XPS:~/repo/go/hello.io$ code main.go (edit using visual code)

## create go module
jasintha@jasintha-XPS:~/repo/go/hello.io$ go mod init main.go

### single package single file
jasintha@jasintha-XPS:~/repo/go/hello.io$ go build -o hell .
jasintha@jasintha-XPS:~/repo/go/hello.io$ ./hell
Hello Go

### two files in same package
jasintha@jasintha-XPS:~/repo/go/hello.io$ go build -o hell .
jasintha@jasintha-XPS:~/repo/go/hello.io$ ./hell
Hello Go
welcome go

### two files in same package and other file in a different package
jasintha@jasintha-XPS:~/repo/go/hello.io$ go build -o hell .
jasintha@jasintha-XPS:~/repo/go/hello.io$ ./hell
Hello Go
welcome go
Foo go

