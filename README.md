# go for devops 

## Why should I learn Golang as a DevOps Engineer?

* If you stay only in Python, you will **use** DevOps tools.
  If you learn Go, you can **build and control** them.

* Kubernetes is not Python-native.
  Writing operators/controllers → Go is the standard path.
  Python options exist but are limited and not widely used.

* Real constraint: **Python GIL**

  * 1 process = effectively 1 CPU core for CPU work
  * Threads won’t scale CPU usage
  * To scale → multiple processes → more RAM + complexity

* Go:

  * No GIL → uses all CPU cores in one process
  * Goroutines are cheap → thousands of concurrent tasks
  * No need to manage processes manually

* Day-to-day impact:

  * Parallel API calls, SSH, scraping → Go is simpler and faster
  * Log/metrics processing → Go handles high throughput better
  * Python starts slowing or needs redesign (multiprocessing, queues)

* Memory difference (practical):

  * Python multiprocessing → each worker = separate memory
  * Go → thousands of goroutines inside one process
  * Result: Go uses significantly less RAM under load

* Deployment reality:

  * Python → need runtime + dependencies + env setup
  * Go → single binary → copy and run anywhere
  * In CI/CD, containers, remote servers → this matters a lot

* Where Python still wins:

  * Quick scripts
  * Data work, ML, automation glue
  * You should keep using it

* Why learn Go anyway:

  * To move from “automation scripts” → “infrastructure tooling”
  * To work deeper with Kubernetes ecosystem
  * To build production-grade tools instead of scripts

**Bottom line:**
You don’t replace Python.
You add Go when scale, performance, or Kubernetes-level control becomes a requirement.


# Run a Go file
GO baased DevOps projects & learning

# Basic
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

Or, **Build and Run as Binary:**

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

# Loop
- for loop
  
  ```go
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }
  ```
- for without the post statement - **while** loop

  ```go
  i := 0
  for i < 10 {
    fmt.Println(i)
    i++
  }
  ```
  
  or,

  ```go
  run = true
  for run {
    fmt.Println("Running since I was born")
  }
  ```

- for without anything - **infinite** loop

  ```go
  for {
    fmt.Println("shit")
  }
  ```

# Common bugs / things to remember
- **variable shadowing**
  - solution: don't use variables of same name in multiple places.
- a declared package must be used
- a declared variable must be used
- **for loops open braces must be on the same line as the for keyword.**


> minhaj was here