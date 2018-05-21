package main

import "fmt"
import "os"
import "sort"

func main() {
    var N int
    fmt.Scan(&N)

    strengths := make([]int, N)
    for i := 0; i < N; i++ {
        var Pi int
        fmt.Scan(&Pi)
        strengths[i] = Pi
    }
    sort.Ints(strengths)

    distance := -1
    for i:= 0; i < N - 1; i++ {
        currDist := strengths[i+1] - strengths[i]

        if distance < 0 {
            distance = currDist
        } else if currDist < distance {
            distance = currDist
        }
    }

    fmt.Fprintln(os.Stderr, strengths)
    fmt.Println(distance)// Write answer to stdout
}
