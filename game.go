package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
	_ "github.com/lib/pq"
	"github.com/micmonay/keybd_event"
)

const averageWordLength = 5
const theSpaceDelimiter = " "
const accessStringConcatentaionValue = "`y8q@4weygaskldn0t^"

func randomNumberGenerator(level string) int {
	minTextRange, maxTextRange := 10, 81 //default values
	if level == "Easy" {
		minTextRange = 10
		maxTextRange = 31
	} else if level == "Medium" {
		minTextRange = 32
		maxTextRange = 61
	} else {
		minTextRange = 62
		maxTextRange = 81
	}
	rand.Seed(time.Now().UnixNano())
	textLength := rand.Intn(maxTextRange) + minTextRange // Generate a random integer in the specified range
	return textLength
}

func getRandomElements(dictionary []string, textLength int) string {
	selected := make(map[int]bool)
	finalText := ""
	for len(selected) < textLength {
		index := rand.Intn(len(dictionary)) // generate a random index within the range of the dictionary array
		if !selected[index] {               // check if the index has already been selected
			selected[index] = true // if not selected, mark it as selected and print the corresponding element
			finalText += dictionary[index] + theSpaceDelimiter
		}
	}
	finalText = strings.TrimRight(finalText, theSpaceDelimiter)
	return finalText
}

func readFile(level string, fileName string) string {
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

	textLength := randomNumberGenerator(level)
	finalText := getRandomElements(dictionary, textLength)
	return finalText
}

func levelSelector(levelChoice int) string {
	if levelChoice == 0 {
		return "Easy"
	} else if levelChoice == 1 {
		return "Medium"
	} else {
		return "Hard"
	}
}

func MetricsCalculation(username string, password string, input string, text string, averageWordLength int, elapsed float32) {
	correct := 0

	fmt.Printf("\033[34m▀ \033[0m")
	fmt.Printf("\033[32m%s\033[0m\n\n", text)
	fmt.Printf("\033[34m▀ \033[0m")
	for i := 0; i < len(text); i++ {
		if i < len(input) && input[i] == text[i] {
			fmt.Printf("\033[32m%c\033[0m", text[i])
			correct++
		} else {
			fmt.Printf("\033[31m%c\033[0m", text[i])
		}
	}
	fmt.Print("\n-------------------------------------------------------------------------------------------\n\n")

	fmt.Println("\033[32m▀ Right\033[0m")
	fmt.Println("\033[31m▀ Wrong\033[0m")

	wpm := int((float32(len(input)) / float32(averageWordLength)) / float32(elapsed))
	fmt.Printf("\033[33mWPM: %d\033[0m\n", wpm)

	acc := int((float32(correct) / float32(len(input))) * 100)
	if acc < 0 {
		acc = 0
	}
	fmt.Printf("\033[33mACC: %d\033[0m\n", acc)

	raw := int(float32(len(input)) / (float32(elapsed * 60)))
	fmt.Printf("\033[33mRaw: %d\033[0m\n", raw)

	updateDataBase(username, password, strconv.Itoa(wpm), strconv.Itoa(acc), strconv.Itoa(raw))
}

func CommandLineOptionsSetter(options []string, usage string) string {
	selectedIndex := 0

	selectOptions(options, selectedIndex) // Print the initial options

	// Listen for arrow key inputs
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyArrowUp {
			selectedIndex--
			if selectedIndex < 0 {
				selectedIndex = len(options) - 1
			}
			selectOptions(options, selectedIndex)
		} else if key == keyboard.KeyArrowDown {
			selectedIndex++
			if selectedIndex >= len(options) {
				selectedIndex = 0
			}
			selectOptions(options, selectedIndex)
		} else if key == keyboard.KeyEnter {
			// User has made a selection, exit loop
			break
		} else {
			fmt.Println("Invalid key")
		}
	}
	clear()

	if usage == "Level" {
		keyStrokeHelper("Level")
		return levelSelector(selectedIndex)
	} else if usage == "Access" {
		username := ""
		password := ""
		if selectedIndex == 0 {
			fmt.Print("Enter Username : ")
			fmt.Scanln(&username)

			fmt.Print("Enter Password : ")
			fmt.Scanln(&password)
			SignUp(username, password)
		} else {
			loginFlag := false
			for !loginFlag {
				fmt.Print("Enter Username : ")
				fmt.Scanln(&username)

				fmt.Print("Enter Password : ")
				fmt.Scanln(&password)

				loginFlag = Login(username, password)
			}
			time.Sleep(1 * time.Second)
			keyStrokeHelper("Access")
		}
		return username + accessStringConcatentaionValue + password
	} else if usage == "Players" {
		keyStrokeHelper("Players")
		if selectedIndex == 0 {
			return "1Player"
		} else {

			return "2Players"
		}
	}

	return ""
}

func selectOptions(options []string, selectedIndex int) {
	fmt.Print("\033[H\033[2J") // Clear the console before printing the options

	for i, option := range options {
		if i == selectedIndex {
			fmt.Printf("\033[1m\033[7m> %s\033[0m\n", option) // Use ANSI escape codes to set the selected option to a different color
		} else {
			fmt.Printf("  %s\n", option)
		}
	}
}

func keyStrokeHelper(usage string) {
	if usage == "Level" {
		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			panic(err)
		}
		kb.SetKeys(keybd_event.VK_LEFT)

		// Press the selected keys
		err = kb.Launching()
		if err != nil {
			panic(err)
		}
	} else if usage == "Access" {
		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			panic(err)
		}
		kb.SetKeys(keybd_event.VK_DOWN)

		// Press the selected keys
		err = kb.Launching()
		if err != nil {
			panic(err)
		}
	} else if usage == "Players" {
		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			panic(err)
		}
		kb.SetKeys(keybd_event.VK_DOWN)

		// Press the selected keys
		err = kb.Launching()
		if err != nil {
			panic(err)
		}
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clear()
	println()

	// "GoTyper"
	fmt.Println("    ▄████  ▒█████  ▄▄▄█████▓▓██   ██▓ ██▓███  ▓█████  ██▀███  ")
	fmt.Println("   ██▒ ▀█▒▒██▒  ██▒▓  ██▒ ▓▒ ▒██  ██▒▓██░  ██▒▓█   ▀ ▓██ ▒ ██▒")
	fmt.Println("  ▒██░▄▄▄░▒██░  ██▒▒ ▓██░ ▒░  ▒██ ██░▓██░ ██▓▒▒███   ▓██ ░▄█ ▒")
	fmt.Println("  ░▓█  ██▓▒██   ██░░ ▓██▓ ░   ░ ▐██▓░▒██▄█▓▒ ▒▒▓█  ▄ ▒██▀▀█▄  ")
	fmt.Println("  ░▒▓███▀▒░ ████▓▒░  ▒██▒ ░   ░ ██▒▓░▒██▒ ░  ░░▒████▒░██▓ ▒██▒")
	fmt.Println("   ░▒   ▒ ░ ▒░▒░▒░   ▒ ░░      ██▒▒▒ ▒▓▒░ ░  ░░░ ▒░ ░░ ▒▓ ░▒▓░")
	fmt.Println("    ░   ░   ░ ▒ ▒░     ░     ▓██ ░▒░ ░▒ ░      ░ ░  ░  ░▒ ░ ▒░")
	fmt.Println("  ░ ░   ░ ░ ░ ░ ▒    ░       ▒ ▒ ░░  ░░          ░     ░░   ░ ")
	fmt.Println("        ░     ░ ░            ░ ░                 ░  ░   ░     ")
	fmt.Println("                             ░ ░                              ")

	time.Sleep(2 * time.Second)
	clear()
	println()

	username := ""
	password := ""

	scanner := bufio.NewScanner(os.Stdin)

	accessOptions := []string{"SignUp", "Login"}
	accessValues := CommandLineOptionsSetter(accessOptions, "Access")
	username = strings.Split(accessValues, accessStringConcatentaionValue)[0]
	password = strings.Split(accessValues, accessStringConcatentaionValue)[1]
	clear()
	println()

	playingOptions := []string{"1 player", "2 players"}
	playersValues := CommandLineOptionsSetter(playingOptions, "Players")
	clear()
	println()

	levelOptions := []string{"Easy", "Medium", "Hard"}
	level := CommandLineOptionsSetter(levelOptions, "Level")
	clear()
	println()

	text := readFile(level, "Dataset.txt") // text to have the test on
	if playersValues == "2Players" {
		port := ""
		fmt.Print("Enter Port: ")
		fmt.Scanln(&port)

		sender := ""
		fmt.Print("Sender? y/n :  ")
		fmt.Scanln(&sender)
		if sender == "y" {
			SendDataToServer(text, port)
		} else if sender == "n" {
			startTCPServer(port) // TODO: The server part of the program
		} else {
			fmt.Print("Enter only 'y' or 'n'")
		}
		clear()
		println()
	}
	fmt.Print(text)
	fmt.Println("\033[0;5H")

	start := time.Now()

	input := ""
	if scanner.Scan() {
		input = scanner.Text()
	}

	elapsed := time.Since(start).Minutes()
	clear()
	MetricsCalculation(username, password, input, text, averageWordLength, float32(elapsed))
}
