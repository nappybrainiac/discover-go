package main

import (
  "fmt"
  "time"
  "strconv"
)

type user struct {
  Name string
  DOB string
  City string
}

func main() {
  u := user{
    Name: "Betty Holberton",
    DOB: "March 7, 1917",
    City: "Philadelphia",
  }

dYOB := u.DOB
YOB := string(dYOB[len(dYOB)-4:])
YOBInt, _ := strconv.ParseInt(YOB, 10, 0)

/*if s, err := strconv.Atoi(YOB); err == nil {
		fmt.Printf("%v", s)
    return s
  }*/

now := time.Now()
YOBInt2 := int(YOBInt)
age := now.Year() - YOBInt2


fmt.Printf("Hello %s\n", u.Name)
fmt.Printf("%s who was born in %s would be %d years old today\n", u.Name, u.City, age)


}
