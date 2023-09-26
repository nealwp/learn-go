package main

import "fmt"

type gasEngine struct {
    mpg uint8
    gallons uint8
}

type electricEngine struct {
    mpkwh uint8
    kwh uint8
}

func (e electricEngine) milesLeft() uint8 {
    return e.kwh * e.mpkwh
}

func (e gasEngine) milesLeft() uint8 {
    return e.gallons * e.mpg
}

type engine interface {
    milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
    if miles <= e.milesLeft(){
        fmt.Println("You can make it")
    } else {
        fmt.Println("Need to fuel up")
    }
}


func main() {
    var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
    fmt.Println(myEngine.mpg, myEngine.gallons)
    fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())

    canMakeIt(myEngine, 50)
}
