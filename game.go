package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	host           = "localhost"
	port           = 5432
	user           = "postgres"
	admin_password = "1234"
	dbname         = "GoTyper-DB"
)

const averageWordLength = 5
const theSpaceDelimiter = " "

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

func levelSelector(levelChoice string) string {
	if levelChoice == "1" {
		return "Easy"
	} else if levelChoice == "2" {
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

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func updateDataBase(username string, password string, wpm string, acc string, raw string) {
	psqlconn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, admin_password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	updateStatement := `UPDATE users
						SET username = $1, password = $2, wpm = $3, acc = $4, raw = $5 
						WHERE username = $1;
						`
	_, e := db.Exec(updateStatement, username, password, wpm, acc, raw)
	CheckError(e)
}

func Login(username string, password string) bool {
	psqlconn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, admin_password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	checkStatement := `SELECT username, password, wpm, acc, raw
					   FROM users
					   where username = $1 and password = $2;
					  `
	row := db.QueryRow(checkStatement, username, password)
	switch err := row.Scan(&username, &password); err {
	case sql.ErrNoRows:
		fmt.Println("Wrong username or password. Try again!")
		return false
	case nil:
		return true
	default:
		fmt.Println("Welcome back,", username+"!")
		return true
	}
}

func SignUp(username string, password string) {
	psqlconn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, admin_password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	insertDynamicStmt := `insert into "users" ("username", "password", "wpm", "acc", "raw") values ($1, $2, $3, $4, $5)`
	_, e := db.Exec(insertDynamicStmt, username, password, 0, 0, 0)
	CheckError(e)
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

	fmt.Println("(1) SignUp\n(2) Login")
	passingOption := ""
	if scanner.Scan() {
		passingOption = scanner.Text()
		if passingOption == "1" {
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
		}
	}
	clear()
	println()

	fmt.Println("Choose a level : \n(1) Easy\n(2) Medium\n(3) Hard")
	level := ""
	if scanner.Scan() {
		level = scanner.Text()
		level = levelSelector(level)
	}
	clear()
	println()

	text := readFile(level, "Dataset.txt") // text to have the test on
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
