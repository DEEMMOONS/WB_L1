package main

import "fmt"

type Human struct {
  name string
  age string
}

type Action struct {
  Man Human
}

func (h *Human) info() (string, string) {
  return h.name, h.age
}

func (a *Action) ans() string {
  str1, str2 := a.Man.info()
  return str1 + ": " + str2
}

func main() {
  ann := Human{name: "Ann", age: "25",}
  act := Action{Man: ann}
  fmt.Printf(act.ans())
}
