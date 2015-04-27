package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

// The Major Element

// Challenge Description:

// The major element in a sequence with the length of L is the element which appears in a sequence more than L/2 times. The challenge is to find that element in a sequence.
// Input sample:

// Your program should accept as its first argument a path to a filename. Each line of the file contains a sequence of integers N separated by comma. E.g.
// Output sample:

// For each sequence print out the major element or print "None" in case there is no such element. E.g.

// Constraints:
// N is in range [0, 100]
// L is in range [10000, 30000]
// The number of test cases <= 40


func printMajor(s []string){
    n := (len(s) / 2)
    majorMap := make(map[string]int)
    for _, val := range s {
        if _, ok := majorMap[val]; ok{
            majorMap[val] ++
        } else {
            majorMap[val] = 1
        }
    }

    l := make([]string, 0)

    for k, v := range majorMap{
        if v > n {
            l = append(l, k)
        }
    }

    switch{
        case len(l) > 0: for _, v := range(l){fmt.Println(v)}
        default: fmt.Println("None")
    }
}

func readLines() {
    fileName := os.Args[1]
    inFile, _ := os.Open(fileName)
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        stringSlice := strings.Split(scanner.Text(), ",")
        printMajor(stringSlice)
    }
}

func main(){
    readLines()
}