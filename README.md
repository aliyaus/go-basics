# Golang explained!
Golang, or Go, is exceptionally well-suited for building scalable and efficient systems. Its lightweight concurrency model, powered by goroutines and channels, makes it ideal for handling high-performance applications, such as distributed systems and microservices. Go emphasizes simplicity and readability, with a minimalistic syntax that reduces development overhead and enhances maintainability. It is also highly efficient, as it compiles to machine code, offers low-level memory control, and has fast startup times, making it perfect for modern cloud-native applications. Additionally, Go excels at networking and API development, with a robust standard library that simplifies creating web servers, RESTful APIs, and other networked services. These strengths make it a top choice for developers working on high-scale, reliable, and performant software.

## Short Variable Declaration (:=)
For local variables, you can use the := shorthand to declare and assign a variable without explicitly specifying the type. The type is inferred from the value.
`variableName := value`

## Declare a Variable with Explicit Type
You can declare a variable with a specific type and then assign a value.
```
var variableName type
variableName = value

var age int
age = 25
```

## Declare and Assign in a Single Statement
You can declare a variable and assign a value at the same time.
```
var variableName type = value

var name string = "John"
var isActive bool = true
```

## Passing arguments to a function
Passing by Value
```
func functionName(param Type) {
    // function body                                        # Call By Value
}
```
Passing by References
```
func functionName(param *Type) {
    // function body                                       # Call By Reference
}
```

## Arrays
In Go, arrays have a fixed size, so you cannot create an array without specifying its length. However, Go provides slices, which are more commonly used because they provide dynamic resizing and are more flexible than arrays.
Creating an Array (fixed size)
```
var arr [5]int // Array of size 5
arr[0] = 1
fmt.Println(arr) // Output: [1 0 0 0 0]
```
Creating a Slice (dynamic size)
```
var slice []int // A slice without a predefined size
slice = append(slice, 1, 2, 3) // Append dynamically
fmt.Println(slice) // Output: [1 2 3]
```

## Loops
Golang only has 1 loop ---> the for-loop. But a while loop can be done with the for loop: 
```
// Basic for loop
for i := 0; i < 5; i++ {
        fmt.Println(i)
}

// While loop as for loop
i := 0
    for i < 5 { // Equivalent to while(i < 5)
        fmt.Println(i)
        i++
    }
```

## Maps
```
myMap := map[string]int{"a": 1, "b": 2}

    for key, value := range myMap {
        fmt.Println("Key:", key, "Value:", value)
    }
```
Adding to a map
```
Adding elements
    myMap["apple"] = 5
    myMap["banana"] = 3

    // Updating an element
    myMap["banana"] = 10
```
Deleting from a map 
```
 myMap := map[string]int{"apple": 5, "banana": 3}

delete(myMap, "banana") // Removes "banana"
```

## Basic Logging
```
import ("log"
)
log.Println("Log Message")
```

# `go.mod` explained
The go.mod file defines the module’s metadata and dependencies. It is the primary file for Go's module system.

## Key Roles of `go.mod`:
* Declares the Module Name:
The first line defines the module's name, typically the import path for your project.
Example:
`module github.com/username/project`
* Specifies Dependencies:
Lists all the dependencies (third-party modules) required by your project and their versions.
Example:
```
require (
    github.com/gin-gonic/gin v1.8.1
    golang.org/x/exp v0.0.1-20250121-abcdef123456
)
```
* Manages Go Version:
Indicates the version of Go required to build the module.
Example: `go 1.20`
* Reproducibility:
By locking the dependency versions, it ensures your project builds the same way regardless of the machine or environment.
How It's Created and Updated:

When you initialize a module using go mod init, a go.mod file is created.
It is updated automatically when you add, remove, or upgrade dependencies using commands like go get.

# `go.sum` explained
The go.sum file records the checksums of the dependencies listed in go.mod. It is used for verifying the integrity of downloaded modules.

## Key Roles of `go.sum`:
* Verifies Dependency Integrity: Ensures that the modules downloaded match their expected content using cryptographic hashes.
* Records Checksums for Transitive Dependencies: Stores checksums not only for direct dependencies but also for all transitive dependencies (dependencies of dependencies).
* Reproducibility and Security: Helps prevent supply chain attacks by ensuring the modules are not tampered with.
* How It’s Managed: Automatically updated by Go commands like go mod tidy or go get.
You should commit it to version control (e.g., Git) along with the go.mod file.

# Best Practices for `go.mod` & `go.sum`
1. Always Commit Both Files: These files ensure reproducibility and are essential for team collaboration.
2. Use go mod tidy Regularly: Cleans up unused dependencies and ensures the go.mod and go.sum files are up-to-date.
3. Avoid Manually Editing go.sum: Let Go commands manage this file automatically.
4. Verify Dependencies: Use commands like go mod verify to ensure the integrity of downloaded modules.

# `go mod tidy`
`go mod tidy` cleans up the go.mod and go.sum files by adding missing dependencies, removing unused ones, and ensuring all modules are correctly verified and consistent with your code imports.

## When Should You Use `go mod tidy?`
* After Adding or Removing Imports: If you manually add or remove imports in your code, run go mod tidy to update the dependencies.
* Before Committing Changes: Running go mod tidy ensures that your go.mod and go.sum files are clean and accurate before pushing code to version control.
* To Clean Up Legacy Dependencies: If your project has unused dependencies lingering from earlier development stages, go mod tidy will remove them.
* After Switching Branches: Different branches may have different dependencies. Running go mod tidy ensures the current branch's dependencies are correct.

# Notes
- In Go, only identifiers (functions, variables, etc.) that start with an uppercase letter are exported and accessible from other files or packages.
- When you pass without pointers (Pass By Value), a copy of the variable is created which is very inefficient. Passing arguments by pointer (*) in Go allows functions to directly modify the original data, improves efficiency by avoiding unnecessary copies of large structures, and enables sharing the same instance across functions. It is essential when working with large data, making persistent changes, or coordinating shared data. 