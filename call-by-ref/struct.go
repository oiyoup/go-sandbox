package main

type Person struct {
    name string
    age  int
}

func (p *Person) toString() {
    p.age++
    println(p.name, "addr:", p)
}

func main() {
    p1 := Person{name: "p1", age: 32}
    p2 := Person{"p2", 29}

    println("p1 addr:", &p1, ", p2 addr:", &p2)
    
    p1.toString()
    p2.toString()

    println("p1 age:", p1.age, ", p2 age:", p2.age)
}
