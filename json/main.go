package main

import (
    "encoding/json"
    "fmt"
)

type Response1 struct {
    Page int
    Fruits []string
}

type Response2 struct {
    Page int `json:"page"`
    Fruits []string `json:"fruits"`
}

func main()  {
    boolB, _ := json.Marshal(true)
    fmt.Println(string(boolB))
    
    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    slcD := []string{"hello", "world", "good"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"hello":23, "world":45, "good":99}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))
    
    res1D := &Response1{
        Page:   1,
        Fruits: []string{"hello", "world", "good"},
    }
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    res2D := &Response2{
        Page:   1,
        Fruits: []string{"hello", "world", "good"},
    }
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))
}
