# go for devops 
Learning golang to see how it can help me as a DevOps Engineer. In this repo, you will all the practice scripts under the [/learning](./learning/) directory. And cool DevOps projects under the [/projects](./projects/) directory. Golang basics and study notes can be found in this readme.

# Table of Index


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

# Conditionals
- `if/else`
- `switch`

  ```go
  if [evaluation] {
      // execute/do something
  } else if [evaluation] {
      // execute/do something else
  } else {
      // execute/do something else
  }
  ```

> go has the ability to execute a statement within the if scope before the evaluation is made.
> we can call it shortly - "short-circuit evaluation" or "statement in condition".

# Function

- multiple returns values supported
- **variadic argument** - provide unknown number of argument to a function
- named return values

```go
func sum(a int, b int) int {
	return a + b
}
```

named returns -

```go
func multiple(a int, s string) (result int, name string) {
	result = a * 2
	name = s + " boss"

	return result, name
}
```

- named returns are already created and ready to use inside the function

# Public and Private

- Declaring constants/variables/functons/methods - when they can be called.
- Public 
  - Exported
  - Must be declared outside functions/methods
  - must start with Capital letterv
- Private

# Data types
(most common)
- int
- float
- string
- bool
- **list**
  - fixed size, can not grow
  - `[<size>]<data_type>{value/empty}` - `[3]string{"min", "ha", "j"}`
- **slice**
  - array but with no size constraint, can grow.
  - `[<empty>]<data_type>{value/empty}` - `[]string{"min", "ha", "j"}`
- **map**
  - key value pairs (dict in Python)
  - can grow beyond it's size declared. So size is optional here.
  - maps have `non-deterministic` order - on each iteration data order will be different. why?
  - declare using *make* `make(map[<key_data_type>]<value_data_type>, size)` - `make(map[string]int, 10)`
  - declare using composite literal - `map[int]string{}`
- **pointers**
  - under the hood of each variable, the memory allocator allocates a space (memory address) to store value to.
  - prepend `&` before a variable to see the mem addr. `print(`&a`)` 
  - function arguments are copies. Changing a value passed to the functon inside the the function does not change the original.
  - *Pointer stores the address of a value, not hte value itself*.
    - *stack* - for exclusive use for function/method call. is faster than heap.
    - *heap allocation* - when it cant determine to live exclusively in a function call.
  - *Use a. pointer on long live objects and where the copy might be expensive*.

# Common bugs / things to remember
- you can only create a variable that does not exist
  - a declared variable must be used or assigned to nothing - ` _ = a`
- **variable shadowing**
  - solution: don't use variables of same name in multiple places.
- a declared package must be used
- a declared variable must be used
- **for loops open braces must be on the same line as the for keyword.**
- for `if else` conditionals, if there is another statement in chain, it must start on the same line at the previous one ends.
- `println` vs `fmt.Println` 
  - println is built into program for developers only.
    - not meant for production use
    - reports to `stderr`
  - fmt.Println is for production use
    - reports to `stdout`
- get data type of a variable
  - `fmt.Println("%T\n", name)         // does not work`
  - `fmt.Println(reflect.TypeOf(name)) // works`
- go is a garbage collected language
  - how to solve garbage collection? [see what go officially says](https://go.dev/doc/gc-guide)
- go's memory model is heap/stack model.
  - compiler does `escape analysis` - what is that? 
    - a compiler optimization technique that determines the appropriate memory location (stack or heap) for a variable based on its lifetime
  - *stack* - for exclusive use for function/method call. is faster than heap.
  - *heap allocation* - when it cant determine to live exclusively in a function call.
> `go1.25.3` ashraf-minhaj was here