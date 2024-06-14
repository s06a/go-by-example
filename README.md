# GO by example

Examples and notes from [gobyexample.com](https://gobyexample.com/)

## abridged version

**variabels** (more in [variables](examples%20and%20notes/02_variables.go))
```go
var a, b int = 1, 2
c := "value"
var d, e = 1, true
```

**constants** (more in [constants](examples%20and%20notes/03_constants.go))
```go
const pi float32 = 3.14
```

**for loop** (more in [for loop](examples%20and%20notes/04_for_loop.go))
```go
for i := range 4 {
    // do something
}

for i := 0; i < 5; i++ {
    // do something
}
```

**if else** (more in [if else](examples%20and%20notes/05_if_else.go))
```go
if a := 7; a%2 == 0 {
    // do something
} else if a%2 == 0 || b%2 == 0 {
    // do something
} else {
    // do something
}
```

**switch** (more in [switch](examples%20and%20notes/06_switch.go))
```go
switch time.Now().Weekday() {
case time.Friday:
    // do something
default:
    // do something
}
```

**arrays** (more in [arrays](examples%20and%20notes/07_arrays.go))
```go
var a [5]int // or a := [5]int, also a := [...]int{1, 2, 3}
a[4] = 2
b := [...]int{1:2, 0:4, 2:6} // index:value
twoD := [4][4]int // call or set with twoD = [i][j]
```