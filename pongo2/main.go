package main

import (
    "fmt"
    "github.com/flosch/pongo2"
    "log"
)

func main() {
    tpl, err := pongo2.FromString("Hello {{name|capfirst}}!")
    if err != nil {
        log.Fatal(err)
    }
    
    out, err := tpl.Execute(pongo2.Context{"name": "florian"})
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(out)
}
