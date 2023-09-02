package main

import "fmt"

var prefix = map[string]string{
  "":        "Hello",
  "Spanish": "Hola",
  "French":  "Bonjour",
}

func Hello(name string, language string) string  {
  if name == ""{
    name = "World"
  }
  return fmt.Sprintf("%s, %s.", prefix[language], name)
}

func main() {
  fmt.Println(Hello("world",""))
}
