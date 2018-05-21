package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "math"
import "strconv"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var LON string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&LON)
    rLON, err := strconv.ParseFloat(strings.Replace(LON, ",", ".", 1), 64)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
    rLON = DegreeToRadian(rLON)

    var LAT string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&LAT)
    rLAT, err := strconv.ParseFloat(strings.Replace(LAT, ",", ".", 1), 64)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
    rLAT = DegreeToRadian(rLAT)

    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)

    //address := ""
    distance := -1.0
    id := 0
    defibs := make(map[int][]string, N)
    for i := 0; i < N; i++ {
        scanner.Scan()
        DEFIB := scanner.Text()
        splitDefib := strings.Split(DEFIB, ";")

        defibId, err := strconv.Atoi(splitDefib[0])
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        defibs[defibId] = splitDefib
    }

    for i, tab := range defibs {
        currLon, err := ReadCoordinate(tab[4])
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
        currLon = DegreeToRadian(currLon)
        currLat, err :=ReadCoordinate(tab[5])
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
        currLat = DegreeToRadian(currLat)

        distLon := (rLON - currLon) * math.Cos((currLat + rLAT) / 2)
        distLat := rLAT - currLat
        currDist := math.Sqrt(math.Pow(distLon, 2) + math.Pow(distLat, 2)) * 6371

        if distance < 0 || distance > currDist {
            distance = currDist
            id = i
        }
    }

    fmt.Println(defibs[id][1])// Write answer to stdout
}

func DegreeToRadian(degree float64) float64 {
    return degree * math.Pi / 180
}

func ReadCoordinate(coordinate string) (float64, error) {
    return strconv.ParseFloat(strings.Replace(coordinate, ",", ".", 1), 64)
}
