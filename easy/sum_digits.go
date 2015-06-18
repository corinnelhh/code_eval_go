package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

// Given a positive integer, find the sum of its constituent digits.

// INPUT SAMPLE:

// The first argument will be a path to a filename containing positive integers, one per line. E.g.

// 23
// 496
// OUTPUT SAMPLE:

// Print to stdout, the sum of the numbers that make up the integer, one per line. E.g.

// 5
// 19

func readLines() {
    fileName := os.Args[1]
    inFile, _ := os.Open(fileName)
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        stringSlice := strings.Split(scanner.Text(), "")
        out := 0
        for _, num := range stringSlice {
            i, _ := strconv.Atoi(num)
            out += i
        }
        fmt.Println(out)
    }
}

func main(){
    readLines()
}