## Fun Facts in Go


### return interface{}

Did you know? In Go, the interface{} type serves as a wildcard, allowing functions to return values of any type. For example, imagine a function called functionOfSomeType() that returns interface{}. This means it can return a value of any type! This flexibility enables Go programs to handle a wide range of data types seamlessly, making coding in Go both dynamic and exciting!
e.g.
```go
func functionOfAnyType() interface{} {
	return 42
}
```
can also be
```go
func functionOfAnyType() interface{} {
	return true
}
```
### named return values

In Go, functions can return multiple values, but by **giving them names**, you can make the code even more snappier!
When you name the return values of a function, Go can automatically return those named values without you having to specify anything in the return statement.
e.g.
```go
func namedReturnValues(s int) (n int, f bool) {
	n = s + 1
	f = true
	return // here, you didn't have to specify the return types
}
```
and here is the cool part! if you forget to change a named return value inside the function, Go doesn't panic! It simply returns the **zero value** of that return type. 
e.g.
```go
func namedReturnValues(s int) (n int, f bool) {
	return // this will return 0 false
}
```




