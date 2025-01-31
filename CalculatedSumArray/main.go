package main

import "fmt"


func arrayCalculate(arr []int) {
    sum := 0
    for i := 0; i < len(arr); i++ {
        sum += arr[i]
    }
    fmt.Println(sum)
}


func main() {
    arr := []int{5,5,5,5,5,5,5,5}
    arrayCalculate(arr)
}
