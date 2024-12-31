package main

import (
    "fmt"
    "recursive-number-algo/src/algorithm"
)

func main() {
    // Test different inputs
    inputs := []int{9, 10, 100}
    
    for _, input := range inputs {
        result := algorithm.RecursiveFunction(input)
        fmt.Printf("Input: %d\nOutput: %v\n\n", input, result)
    }
}