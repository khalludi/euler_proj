package main

import "fmt"

func main() {
  sum := 0
  for num := 0; num < 1000; num = num + 1 {
    if num % 3 == 0 || num % 5 == 0 {
      sum = sum + num
    }
  }

  fmt.Printf("%d", sum)
}
