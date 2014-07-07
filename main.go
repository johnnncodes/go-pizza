package main

import (
    "fmt"
    "time"
    "math/rand"
)

type Maker struct {
    quantity int
}

func (m *Maker) Bake(c chan<- string) {
    rand.Seed(time.Now().UnixNano())
    amt := time.Duration(rand.Intn(10))
    time.Sleep(time.Second * amt)
    c <- fmt.Sprintf("Pizza created in %d second(s).\n", amt)
}

func (m *Maker) Make() {
    c := make(chan string)

    for i := 0; i < m.quantity; i++ {
        go m.Bake(c)
    }

    for i := 0; i < m.quantity; i++ {
        fmt.Printf(<- c)
    }
}

func main() {
    fmt.Printf("Enter the number of pizzas you want to be created: ")

    var input int
    fmt.Scanln(&input)

    maker := Maker{
        quantity: input,
    }

    maker.Make()

    fmt.Printf("Done creating %d pizza(s). Press enter to quit.", maker.quantity)

    fmt.Scanln(&input)
}
