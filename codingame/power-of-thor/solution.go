package main

import "fmt"

func main() {
    // lightX: the X position of the light of power
    // lightY: the Y position of the light of power
    // initialTX: Thor's starting X position
    // initialTY: Thor's starting Y position
    var lightX, lightY, initialTX, initialTY int
    fmt.Scan(&lightX, &lightY, &initialTX, &initialTY)

    currX := initialTX
    currY := initialTY
    for {
        // remainingTurns: The remaining amount of turns Thor can move. Do not remove this line.
        var remainingTurns int
        fmt.Scan(&remainingTurns)

        //fmt.Fprintln(os.Stderr, findX(initialTX, lightX) + findY(initialTY, lightY))
        dirY, y := findY(currY, lightY)
        dirX, x := findX(currX, lightX)

        currY += y
        currX += x

        // A single line providing the move to be made: N NE E SE S SW W or NW
        fmt.Println(dirY + dirX)
    }
}

func findY(from int, to int) (string, int) {
    if from > to {
        return "N", -1
    }

    if from < to {
        return "S", 1
    }

    return "", 0
}

func findX(from int, to int) (string, int) {
    if from > to {
        return "W", -1
    }

    if from < to {
        return "E", 1
    }

    return "", 0
}
