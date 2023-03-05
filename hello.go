package main

import (
  "fmt"
  "log"
  "github.com/crokenx/greetings"
)

func main(){
  log.SetPrefix("greetings: ")
  log.SetFlags(0)

  var names = []string{"Jesus","Erika","Malone"}

  var messages, err = greetings.Hellos(names)

  if err != nil {
    log.Fatal(err)
  }

  // var message, err = greetings.Hello("Jesus")

  // if err != nil {
  //   log.Fatal(err)
  // }

  for _, message := range messages {
    fmt.Println(message);
  }
}
