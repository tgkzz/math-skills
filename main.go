package main

import (
	"bufio"
	"fmt"
	"math"
	"math-skills/internal"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Error: you need to have data.txt file")
	}

	// check the name extension of the file
	ext1 := regexp.MustCompile(`\.([^\s]+)$`).FindStringSubmatch(os.Args[1])

	if ext1 == nil {
		fmt.Println("Error: file don't have a valid extension.")
		return
	}
	if ext1[1] != "txt" {
		fmt.Println("Error: file extensions is not txt.")
		return
	}

	file, err := os.Open(os.Args[1])
	internal.Logfatal(err)
	defer file.Close()

	numbers := []float64{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for _, item := range strings.Fields(line) {
			number, err := strconv.ParseFloat(item, 64)
			internal.Logfatal(err)
			numbers = append(numbers, number)
		}
	}

	sort.Float64s(numbers)

	avg := internal.Average(numbers)
	med := internal.Median(numbers)
	variance := internal.Variance(numbers, avg)
	stdDev := internal.StdDev(variance)

	fmt.Println("Average:", int(math.Round(avg)))
	fmt.Println("Median:", int(math.Round(med)))
	fmt.Println("Variance:", int(math.Round(variance)))
	fmt.Println("Standard Deviation:", int(math.Round(stdDev)))
}
