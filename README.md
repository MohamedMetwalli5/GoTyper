![Language](https://img.shields.io/badge/language-Go%20-cyan.svg)
![DBMS](https://img.shields.io/badge/DBMS-PostgreSQL%20-blue.svg)
![License](https://img.shields.io/badge/license-Apache_2.0-orange.svg)

				    ▄████  ▒█████  ▄▄▄█████▓▓██   ██▓ ██▓███  ▓█████  ██▀███  
				   ██▒ ▀█▒▒██▒  ██▒▓  ██▒ ▓▒ ▒██  ██▒▓██░  ██▒▓█   ▀ ▓██ ▒ ██▒
				  ▒██░▄▄▄░▒██░  ██▒▒ ▓██░ ▒░  ▒██ ██░▓██░ ██▓▒▒███   ▓██ ░▄█ ▒
				  ░▓█  ██▓▒██   ██░░ ▓██▓ ░   ░ ▐██▓░▒██▄█▓▒ ▒▒▓█  ▄ ▒██▀▀█▄  
				  ░▒▓███▀▒░ ████▓▒░  ▒██▒ ░   ░ ██▒▓░▒██▒ ░  ░░▒████▒░██▓ ▒██▒
				   ░▒   ▒ ░ ▒░▒░▒░   ▒ ░░      ██▒▒▒ ▒▓▒░ ░  ░░░ ▒░ ░░ ▒▓ ░▒▓░
				    ░   ░   ░ ▒ ▒░     ░     ▓██ ░▒░ ░▒ ░      ░ ░  ░  ░▒ ░ ▒░
				  ░ ░   ░ ░ ░ ░ ▒    ░       ▒ ▒ ░░  ░░          ░     ░░   ░ 
					░     ░ ░            ░ ░                 ░  ░   ░     
							     ░ ░                              
----------------------------------------------------------------------------------------------------------------------------------------------------
## Run
	go run game.go database_operations.go sender.go receiver.go

## About
This is a terminal-based typing game that can be played solo or with a friend on the same local network. It utilizes socket programming for multiplayer functionality.

## Metrics Calculations
### WPM
* WPM (Words Per Minute): It is used to measure how many words a typist can accurately type within a minute. The calculation is based on the total number of characters typed divided by 5 (since an average word is considered to be 5 characters long like in the dataset i used) and then divided by the time taken in minutes.

* Formula: `WPM = (Total characters typed / 5) / Time taken in minutes`

### ACC
* ACC (Accuracy): It is used to measure the percentage of correctly typed characters. The calculation is based on the total number of correctly typed characters divided by the total number of characters typed, and then multiplied by 100.
* Formula: `Accuracy = (Correctly typed characters / Total characters typed) x 100`

### RAW
* RAW (Raw Speed): It is used to measure the total number of characters typed, regardless of accuracy, within a specific time frame. The calculation is based on the total number of characters typed divided by the time taken.
* Formula: `RAW = Total characters typed / Time taken in seconds`

## Words Dataset
It contains 1000 of the most commonly used English words, as well as numbers and special characters.

## ERD of Users Database
![image](https://user-images.githubusercontent.com/58489322/231894616-d911fce6-75fe-44b0-be02-406d8fc4b82e.png)

## The Logo
![GoTyper](https://user-images.githubusercontent.com/58489322/229482931-17debad2-c555-4202-988b-ee6f0d423507.png)
