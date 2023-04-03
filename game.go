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
			finalText += dictionary[index] + " "
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

func MetricsCalculation(text string, averageWordLength int, elapsed float32) {
	wpm := int((float32(len(text)) / float32(averageWordLength)) / float32(elapsed))
	fmt.Println("WPM:", wpm)
	/*
		acc :=
		fmt.Println("ACC:", acc)
		raw :=
		fmt.Println("Raw:", raw)
	*/
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	text := readFile("Text.txt") // text to have the test on
	fmt.Print(text)
	fmt.Println("\033[0;5H") // Set cursor position to row 0, column 5 (6th character)

	start := time.Now()

	charInput := ""
	fmt.Scanf("%s", charInput)
	// if text[0] == charInput[0] {
	// 	clear()
	// }

	elapsed := time.Since(start).Minutes()
	clear()

	MetricsCalculation(text, averageWordLength, float32(elapsed))

}
