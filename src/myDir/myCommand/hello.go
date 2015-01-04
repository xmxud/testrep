package main

import (
    "fmt"
    "myDir/pkgDir"
)

func main(){
    fmt.Print("main package 111~\n")
    pkgDir.TestPrint()
}