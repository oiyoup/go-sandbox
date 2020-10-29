package main

func main() {
    i := 23
    
    bcd(&i)
    
    println(i)
}

func bcd(i *int) {
    *i = 99
}
