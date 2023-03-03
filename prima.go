package main

import (
    "fmt"
)

//var isPrima int

func main() {

    var y int
    var a int
    var g int
    fmt.Print("\n\t\tDeretan Prima-> ")
    for y=1; y<=1000; y++ {
        a = 0;
            for g=1; g<=y; g++ {
                if y % g == 0 {
                a++;
            }
        }

        if a == 2 {
            fmt.Print(y, ", ")
        }
       
    }
}//end syntax-nya