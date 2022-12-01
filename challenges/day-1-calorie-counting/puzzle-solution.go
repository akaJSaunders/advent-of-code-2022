package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "sort"
)

func main() {
    readFile, err := os.Open("challenges/day-1-calorie-counting/puzzle-input.txt")
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    elf_calories := make([]int, 0)
    var elf_calorie_count = 0

    for fileScanner.Scan() {

        if fileScanner.Text() == "" {
            elf_calories = append(elf_calories, elf_calorie_count)
            elf_calorie_count = 0
        } else {
            calorie, _ := strconv.Atoi(fileScanner.Text())
            if err != nil {
                fmt.Println("Error during conversion")
                return
            }

            elf_calorie_count += calorie
        }

    }

    readFile.Close()
    
    sort.Sort(sort.Reverse(sort.IntSlice(elf_calories)))

    // Part 1
    fmt.Println("Most calories being carried: ", elf_calories[0])

    // Part 2
    fmt.Println("Total calories carried by the top 3:", elf_calories[0] + elf_calories[1] + elf_calories[2])
}
