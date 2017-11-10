package main

import "fmt"

func pingpong(num int) {
  for num > 0 {
    if num % 15 == 0 {
      fmt.Println("ping-pong")
    } else if num % 5 == 0 {
      fmt.Println("pong")
    } else if num % 3 == 0 {
      fmt.Println("ping")
    } else {
      fmt.Println(num)
    }
    num = num -1
  }
}

func main() {
  var inputNum int;
  fmt.Println("Enter a number to start PingPong: ");
  fmt.Scanf("%d", &inputNum)
  pingpong(inputNum)
}
