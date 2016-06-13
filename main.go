package main

import (
  "fmt"
  "math"
)

func main() {
  //Base prime numbers
  p, q := getPrimeNumbers()
  //Length of the set
  n := p * q
  //Euler's totient function result
  phi := getTotientFuncion(p, q)
  // Co-prime of n. MDC(φ(n), e) = 1, being e > 1
  e := getPublicKey(n, 15)

  message := "Oloko bixu"

  fmt.Println("Message %s /n Cyphered Message %s", message, cypherMessage(message, e, n))
}

func getPrimeNumbers() (int, int){
  return 17, 41
}

func getTotientFuncion(p, q int) int {
  return (p - 1) * (q - 1)
}

func getPublicKey(n, coPrimeMinSize int) int {
  return 13
}

func cypherMessage(message string, e int, n int) string {
  var cypheredMessage string

  for i := 0; i < len(message); i++ {
    cypheredMessage[i] = math.Pow(float64(message[i]), float64(e)) % float64(n)
  }

  return cypheredMessage
}
