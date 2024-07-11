package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := make([]int, 0)

	for scanner.Scan() {
		input := scanner.Text()
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error converting input to number:", err)
			continue
		}

		numbers = append(numbers, number)

		if len(numbers) > 1 {
			lowerLimit, upperLimit := predictRange(numbers[:len(numbers)-1])
			fmt.Println(numbers[len(numbers)-1])
			fmt.Printf("%d %d\n", lowerLimit, upperLimit)
		}
	}
}

func predictRange(numbers []int) (int, int) {
	if len(numbers) == 0 {
		return 1, 1000 // Default range if no numbers are available
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	mean := float64(sum) / float64(len(numbers))

	varianceSum := 0.0
	for _, num := range numbers {
		diff := float64(num) - mean
		varianceSum += diff * diff
	}
	variance := varianceSum / float64(len(numbers))
	stdev := math.Sqrt(variance)

	lowerLimit := int(math.Max(1, math.Floor(mean-stdev)))
	upperLimit := int(math.Ceil(mean + stdev))

	return lowerLimit, upperLimit
}
