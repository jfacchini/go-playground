package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func abs(a int) int {
    if a < 0 {
        return a * -1
    }

    return a
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    // n: the number of temperatures to analyse
    var n int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&n)

    scanner.Scan()
    temps := scanner.Text() // the n temperatures expressed as integers ranging from -273 to 5526
    tempsVal := strings.Split(temps, " ")

    result := 5527
    for _, val := range tempsVal {
        if i, err := strconv.Atoi(val); err == nil {
            rabs := abs(result)
            iabs := abs(i)

            if result != i && rabs == iabs {
                result = rabs
            } else if iabs < rabs {
                result = i
            }
        }
    }

    if result == 5527 {
        result = 0
    }

    fmt.Println(result)// Write answer to stdout
}
