package main

import "fmt"
import "os"
import "bufio"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var L int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&L)

    var H int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&H)

    scanner.Scan()
    T := scanner.Text()

    alphabet := ReadAlphabet(scanner, H, L)
    PrintASCII(T, alphabet, H)
}

func ReadAlphabet(scanner *bufio.Scanner, H int, L int) [][]string {
    alphabet := make([][]string, 27)
    for l := 0; l < 27; l++ {
        alphabet[l] = make([]string, H)
    }

    for i := 0; i < H; i++ {
        scanner.Scan()
        ROW := scanner.Text()
        for j := 0; j < 26; j++ {
            alphabet[j][i] = string(ROW[L*j:L*(j+1)])
        }
        alphabet[26][i] = string(ROW[L*26:])
    }

    return alphabet
}

func PrintASCII(T string, alphabet [][]string, H int) {
    for i := 0; i < H; i++ {
        for _, letter := range T {
            // A:65 - Z:90; a:97 - z:122
            if letter >= 97 && letter <= 122 {
                letter = letter - 32
            }

            if letter < 65 || letter > 90 {
                letter = 91
            }

            fmt.Printf("%s", alphabet[letter-65][i])
        }
        fmt.Printf("\n")
    }
}
