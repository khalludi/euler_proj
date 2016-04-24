package main

import "fmt"

func main() {
  sum := 0
  for num := 0; num < 1000; num = num + 1 {
    if num % 3 == 0 || num % 5 == 0 {
      sum = sum + num
    }
  }

  fmt.Printf("The sum of natural numbers under 1000 with multiples of 3 and 5 is: \n%d", sum)
}
