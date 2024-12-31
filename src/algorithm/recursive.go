package algorithm

// RecursiveFunction handles two cases:
// 1. For n <= 10: Powers of 2 up to n/2, then n
// 2. For n > 10: Powers of 2 up to n, then n
func RecursiveFunction(n int) []int {
    if n <= 1 {
        return []int{n}
    }
    
    var result []int
    current := 2
    
    if n <= 10 {
        // Case 1: Powers of 2 up to n/2
        for current <= n/2 {
            result = append(result, current)
            current *= 2
        }
    } else {
        // Case 2: Powers of 2 up to n
        for current < n {
            result = append(result, current)
            if current*2 > n {
                break
            }
            current *= 2
        }
    }
    
    // Add input number if not already in sequence
    if len(result) == 0 || result[len(result)-1] != n {
        result = append(result, n)
    }
    
    return result
}