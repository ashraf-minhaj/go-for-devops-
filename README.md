# go for devops 
GO baased DevOps projects & learning

# Hello world 
1. Install Go.

2. Create a file `main.go`

```go
package main
import "fmt"

func main() {
    fmt.Println("Assalamualaikum Habibi")
}
```

3. Run it:

```
go run main.go
```

That is the simplest way.

Alternative (compile first):

```
go build main.go
./main
```

// package name should be same as the directory name
// all go files in a directory should belong to the same package

# go is statically types

- once a variable is declared it the value type can not be changed. 
- the lack of awareness of variable type makes python to run the whole code and show error during runtime.
- Btter to find the error at compilation rather than while running it.

> minhaj was here