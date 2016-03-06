/* This is my answer to Task 2, and is supposed to familiarize
   you with if/else conditions and loops and emphasizing the
   use of error checking*/


package main

import (
  "fmt"
  "time"
  "strconv"
)

type user struct {
  Name string `json:"name"`
  DOB string `json:"date_of_birth"`
  City string `json:"city"`
}

func hello(u user) {
  //Prints "Hello Betty Holberton"
  fmt.Printf("Hello %s\n", u.Name)
}

func age(u user){
  //assigns the DOB to a variable
  dYOB := u.DOB

  //Extract the last four characters in the string
  YOB := string(dYOB[len(dYOB)-4:])

  //Convert the string "1917" to an integer
  YOBInt, _ := strconv.ParseInt(YOB, 10, 0)

  now := time.Now()

  // Convert the int64 type into int type
  YOBInt2 := int(YOBInt)
  age := now.Year() - YOBInt2

  //Prints <Name> who was born in <City> would be XX years old today
  fmt.Printf("%s who was born in %s would be %d years old today\n", u.Name, u.City, age)
}

func main() {
  u := user{
    Name: "Betty Holberton",
    DOB: "March 7, 1917",
    City: "Philadelphia",
  }

  hello(u)
  age(u)

}
