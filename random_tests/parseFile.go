package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	timecount := make(map[string]int)
	file, err := os.Open("_chat.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(getDate(getDateTime(scanner.Text())))
		line := getDate(getDateTime(scanner.Text()))
		if _, ok := timecount[line]; ok {
			timecount[line] += 1
		} else {
			timecount[line] = 1
		}
		//if (getDate(getDateTime(scanner.Text()))) != "" {
		//}
	}

	fmt.Println(timecount["8:40:55 PM"])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func getDateTime(line string) string {
	date := line[strings.Index(line, "[")+1 : strings.Index(line, "]")]
	return date
}

func getDate(dateTime string) string {
	date := strings.Split(dateTime, ", ")
	return date[1]
}

func getTime(dateTime string) string {
	date := strings.Split(dateTime, ", ")
	return date[0]
}
