package main

import "fmt"

func main(){
    var intNum int = 32767
    fmt.Println(intNum)

    var floatNum float64 = 12345678.9 
    fmt.Println(floatNum)
    
    var intNum1 int = 3
    var intNum2 int = 2
    fmt.Println(intNum1/intNum2)
    fmt.Println(intNum1%intNum2)
    
    var myString string = "hello world"
    fmt.Println(myString)

    var myBoolean bool = false
    fmt.Println(myBoolean)

    var intNum3 int
    fmt.Println(intNum3)

    const myConst string = "const value"
    fmt.Println(myConst)
}
