package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    scanner.Scan()
    MESSAGE := scanner.Text()
    fmt.Fprintf(os.Stderr, "%s\n", MESSAGE)

    text := ""
    for i := 0; i < len(MESSAGE); i++ {
        text += padLeft(strconv.FormatInt(int64(MESSAGE[i]), 2), "0", 7)
    }
    fmt.Fprintf(os.Stderr, "%s\n", text)

    var previous byte
    result := ""
    for i := 0; i < len(text); i++ {
        result += findNext(text[i], previous)
        previous = text[i]
    }

    fmt.Println(result)// Write answer to stdout
}

func findNext(current byte, previous byte) string {
    var empty byte
    zero := "0"

    if current != previous {
        space := ""
        if previous != empty {
            space = " "
        }

        if current == zero[0] {
            return space + "00 0"
        } else {
            return space + "0 0"
        }
    }

    return "0"
}

func padLeft(str string, char string, length int) string {
    for len(str) < length {
        str = char + str
    }

    return str
}
