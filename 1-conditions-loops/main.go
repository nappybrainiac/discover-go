
//This has been created for the purpose of completing task 1 in Discovering Go.
//It explores declaring variables and using them in print statements

package main

import (
  "math/rand"
  "fmt"
)

func main() {

  //rand.Intn returns a random int n, 0 <= n < 100
  randomNumber := rand.Intn(100)

  //Is randomNumber greater than 50?
  if randomNumber > 50 {
    fmt.Printf("my random number is %d and is greater than 50\n", randomNumber)
  } else {
    fmt.Printf("my random number is %d and is equal to or less than 50\n", randomNumber)
  }


  school := "Holberton School"

  //Is the string school equal to Holberton School?
  if school == "Holberton School" {
    fmt.Printf("I am a student of the %s\n", school)
  }

  //Declaring boolean values
  beautifulWeather := true

  //Is the weather  beautiful?
  if beautifulWeather == true {
    fmt.Println("It's a beautiful weather!")
  }

  //This creates a range.
  holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}

  //Loop through the range and print the following text
  for _, each := range holbertonFounders {
    fmt.Printf("%s is a founder\n", each)
  }
}
