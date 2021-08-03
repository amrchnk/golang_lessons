package main

import "fmt"
const pi=3.1415
func main(){
    printCircleArea()
}

func printCircleArea(){
    circleRadius:=2
    CircleArea:=float32(circleRadius)*float32(circleRadius)*pi

    fmt.Printf("Radius: %d \n",circleRadius)
    fmt.Printf("Area: %f32",CircleArea)
}