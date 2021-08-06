package main

import (
    "fmt"
    "net/http"
    "math"
//     "reflect"
    "strconv"
)

func main(){
    http.HandleFunc("/hw",func(w http.ResponseWriter, r *http.Request){
        radius,err:=strconv.Atoi(r.URL.Query().Get("radius"))
        if  err != nil || radius<1{
            return
        }
        fmt.Fprint(w,math.Pow(float64(radius),2)*math.Pi)

    })

    http.ListenAndServe(":80",nil)
}