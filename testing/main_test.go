package main

import (
    "fmt"
    "testing"
)

//go test testing/main_test.go

func TestPrintSomething(t *testing.T) {
    fmt.Println("Say hi")
    t.Log("Say bye")
}