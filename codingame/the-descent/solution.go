package main

import "fmt"

func main() {
    for {
        max := 0
        res := 0
        for i := 0; i < 8; i++ {
            // mountainH: represents the height of one mountain.
            var mountainH int
            fmt.Scan(&mountainH)
            if mountainH > max {
                max = mountainH
                res = i
            }
        }
        
        fmt.Println(res) // The index of the mountain to fire on.
    }
}
