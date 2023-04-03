package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const averageWordLength = 5
const theSpaceSelimiter = "."

func randomNumberGenerator() int {
	rand.Seed(time.Now().UnixNano())
	textLength := rand.Intn(81) + 10 // Generate a random integer between 10 and 90
	return textLength
}

func getRandomElements(dictionary []string, textLength int) string {
	selected := make(map[int]bool)
	finalText := ""
	for len(selected) < textLength {
		index := rand.Intn(len(dictionary)) // generate a random index within the range of the dictionary array
		if !selected[index] {               // check if the index has already been selected
			selected[index] = true // if not selected, mark it as selected and print the corresponding element
			finalText += dictionary[index] + theSpaceSelimiter
		}
	}
	finalText = strings.TrimRight(finalText, " ")
	return finalText
}

func readFile(fileName string) string {
	var dictionary []string
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	scannedText := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scannedText = scanner.Text()
		dictionary = append(dictionary, scannedText)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	textLength := randomNumberGenerator()
	finalText := getRandomElements(dictionary, textLength)
	return finalText
}

func MetricsCalculation(charInput string, text string, averageWordLength int, elapsed float32) {
	correct := 0

	fmt.Printf("\033[32m%s\033[0m\n", text)
	for i := 0; i < len(charInput); i++ {
		if charInput[i] == text[i] {
			fmt.Printf("\033[32m%c\033[0m", charInput[i])
			correct++
		} else {
			fmt.Printf("\033[31m%c\033[0m", charInput[i])
		}
	}
	fmt.Println()

	wpm := int((float32(len(text)) / float32(averageWordLength)) / float32(elapsed))
	fmt.Printf("\033[33mWPM: %d\033[0m\n", wpm)

	acc := int((float32(correct) / float32(len(charInput))) * 100)
	fmt.Printf("\033[33mACC: %d\033[0m\n", acc)

	raw := int(float32(len(charInput)) / float32(elapsed*60))
	fmt.Printf("\033[33mRaw: %d\033[0m\n", raw)

}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clear()
	println()
	text := readFile("Words.txt") // text to have the test on
	fmt.Print(text)
	fmt.Println("\033[0;5H") // Set cursor position to row 0, column 5 (6th character)

	start := time.Now()

	charInput := ""
	fmt.Scanf("%s", &charInput)

	elapsed := time.Since(start).Minutes()
	clear()
	MetricsCalculation(charInput, text, averageWordLength, float32(elapsed))
}
