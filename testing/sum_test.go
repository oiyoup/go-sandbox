package main

import "testing"

//go test

func TestSum(t *testing.T) {
    s := Sum(1, 2, 3)

    if s != 6 {
        t.Error("Wrong result")
    }
}
