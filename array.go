package main

import "fmt"

func main() {
  //How to declare Array Var
  //<var> := []<var type>{<value1>,<value2>,<value3>,...,<valuen>}
  avengers := []string{"Iron Man","Captain America","Hulk","Black Widow","Hawkeye","Thor"}
  
  //How to add value in Array
  // <var> = append(<var>,<value that you want to append in array var>)
  avengers = append(avengers, "Vision")
  avengers = append(avengers, "Scarlet Witch")
  avengers = append(avengers, "Doctor Strange")
  avengers = append(avengers, "Black Panther")
  avengers = append(avengers, "Winter Soldier")

  
  // fmt.Println(avengers[0])
  // fmt.Println(avengers)
  fmt.Println("List Name of Avenger Team")
  for index,avengers := range avengers{
    fmt.Println(" ",index+1,".",avengers)
  }
  
}
