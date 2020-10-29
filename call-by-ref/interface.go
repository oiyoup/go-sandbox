package main

import "fmt"

type Str1 struct {
    id int
    name string
}

func main() {
    var str1 = new(Str1)
    
    abc(str1)
    
    fmt.Println(str1.id)
}

func abc(i interface {}) {
    var str1 = i.(*Str1) 
    str1.id = 23
}
