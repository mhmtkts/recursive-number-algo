package tests

import (
    "testing"
    "reflect"
    "recursive-number-algo/src/algorithm"
)

func TestRecursiveFunction(t *testing.T) {
    testCases := []struct {
        name     string
        input    int
        expected []int
    }{
        {"Input 9", 9, []int{2, 4, 9}},
        {"Input 10", 10, []int{2, 4, 10}},
        {"Input 1", 1, []int{1}},
        {"Input 8", 8, []int{2, 4, 8}},
        {"Input 100", 100, []int{2, 4, 8, 16, 32, 64, 100}},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := algorithm.RecursiveFunction(tc.input)
            if !reflect.DeepEqual(result, tc.expected) {
                t.Errorf("For input %d, expected %v but got %v",
                    tc.input, tc.expected, result)
            }
        })
    }
}