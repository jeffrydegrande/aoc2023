package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "A brief description of your command",
	Run:   day1Main,
}

var numbersPartOne = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var numbers = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}

func day1Main(cmd *cobra.Command, args []string) {
	day1PartOne(cmd, args)
	day1PartTwo(cmd, args)
}

func day1PartOne(cmd *cobra.Command, args []string) {
	f, err := os.Open("day1.txt")
	failIf(err)

	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		count += findNumbersPartOne(line)
	}
	fmt.Println("#1:", count)

	failIf(scanner.Err())
}

func day1PartTwo(cmd *cobra.Command, args []string) {
	f, err := os.Open("day1.txt")
	failIf(err)

	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		count += findNumbersPartTwo(line)
	}
	fmt.Println("#2:", count)

	failIf(scanner.Err())
}

func findNumbersPartOne(s string) int {
	first := 0
	firstIndex := -1
	last := 0
	lastIndex := -1

	for key, value := range numbersPartOne {
		pos := strings.Index(s, key)
		if pos == -1 {
			continue
		}

		if firstIndex == -1 || pos <= firstIndex {
			first = value
			firstIndex = pos
		}
	}

	for key, value := range numbersPartOne {
		pos := strings.LastIndex(s, key)
		if pos == -1 {
			continue
		}

		if lastIndex == -1 || pos >= lastIndex {
			last = value
			lastIndex = pos
		}
	}

	return 10*first + last
}

func findNumbersPartTwo(s string) int {
	first := 0
	firstIndex := -1
	last := 0
	lastIndex := -1

	for key, value := range numbers {
		pos := strings.Index(s, key)
		if pos == -1 {
			continue
		}

		if firstIndex == -1 || pos <= firstIndex {
			first = value
			firstIndex = pos
		}
	}

	for key, value := range numbers {
		pos := strings.LastIndex(s, key)
		if pos == -1 {
			continue
		}

		if lastIndex == -1 || pos >= lastIndex {
			last = value
			lastIndex = pos
		}
	}

	return 10*first + last
}
