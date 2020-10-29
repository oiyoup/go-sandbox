package main

import "fmt"

type My struct {
    Name string
    Age int
}

func main() {
    fmt.Printf("%v\n", 1)
    my := My{"Gyeongun Gim", 41}
    //구조체의 필드도 함께 표시
    fmt.Printf("%+v\n", my)
    fmt.Printf("%#v\n", my)
}
