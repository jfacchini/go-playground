package main

import "fmt"
import "os"
import "bufio"
import "strings"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    // N: Number of elements which make up the association table.
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)

    // Q: Number Q of file names to be analyzed.
    var Q int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&Q)

    mimeTypes := make(map[string]string, N)

    for i := 0; i < N; i++ {
        // EXT: file extension
        // MT: MIME type.
        var EXT, MT string
        scanner.Scan()
        fmt.Sscan(scanner.Text(),&EXT, &MT)
        mimeTypes[strings.ToLower(EXT)] = MT
        fmt.Fprintln(os.Stderr, EXT, mimeTypes[EXT])
    }
    for i := 0; i < Q; i++ {
        scanner.Scan()
        FNAME := scanner.Text() // One file name per line.
        result := "UNKNOWN"
        lastIndex := strings.LastIndex(FNAME, ".")
        if lastIndex > -1 {
            suffix := strings.ToLower(FNAME[lastIndex + 1:])
            fmt.Fprintln(os.Stderr, suffix)
            if mType, ok := mimeTypes[suffix]; ok {
                result = mType
            }
        }
        fmt.Println(result)
    }
}
