package main

import "fmt"

func main(){
    type employee struct{
        name string
        sex string
        age int
        salary int
    }
    emp1:=employee{
        name:"Anek",
        sex:"F",
        age:20,
        salary:100000000,
    }
    emp2:=employee{
            name:"Tem",
            sex:"M",
            age:22,
            salary:320000,
        }
    fmt.Printf("%+v\n",emp1)
    fmt.Printf("%+v\n",emp2)
}