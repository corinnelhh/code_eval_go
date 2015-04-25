package main

import "fmt"
  
func main() {
    i := 1
    for i <= 100 {
        switch {
        case (i % 5 == 0) && (i % 3 == 0): fmt.Println("FizzBuzz")
        case (i % 5 == 0): fmt.Println("Buzz")
        case (i % 3 == 0): fmt.Println("Fizz")
        default: fmt.Println(i)
        }
        i ++
    }
}



