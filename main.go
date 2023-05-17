package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculateMedian(data []int) float64 {
	// Sort the data in ascending order
	sort.Ints(data)

	length := len(data)
	if length%2 == 0 {
		// If the length is even, average the middle two elements
		middle1 := data[length/2-1]
		middle2 := data[length/2]
		return float64(middle1+middle2) / 2.0
	} else {
		// If the length is odd, return the middle element
		middle := data[length/2]
		return float64(middle)
	}
}
func calculateAverage(data []int) float64 {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return float64(sum) / float64(len(data))
}
func calculateVariance(data []int) float64 {
	average := calculateAverage(data)
	sumOfSquares := 0.0
	for _, value := range data {
		deviation := float64(value) - average
		sumOfSquares += deviation * deviation
	}
	return sumOfSquares / float64(len(data))
}
func calculateStandardDeviation(data []int) float64 {
	variance := calculateVariance(data)
	return math.Sqrt(variance)
}
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path as an argument.")
	}

	filePath := os.Args[1]

	content, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	lines := strings.Split(string(content), "\n")
	data := make([]int, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Invalid data format in file: %s", filePath)
		}

		data = append(data, value)
	}
	median := calculateMedian(data)
	average := calculateAverage(data)
	variance := calculateVariance(data)
	standardDeviation := calculateStandardDeviation(data)
	fmt.Println("Median", median)
	fmt.Println("Avarage", average)
	fmt.Println("Variance", variance)
	fmt.Println("Standard Deviation:", standardDeviation)

}
