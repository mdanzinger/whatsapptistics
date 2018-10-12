package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

func main() {
	timecount := make(map[string]int)
	file, err := os.Open("_chat.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	re := regexp.MustCompile(`(\d{4}-\d{1,2}-\d{1,2}).*(\d{1,2}:\d{2}:\d{2} [AP]M)`)

	start := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := re.FindStringSubmatch(scanner.Text())
		line := m[2]
		if _, ok := timecount[line]; ok {
			timecount[line] += 1
		} else {
			timecount[line] = 1
		}
		//fmt.Println(getDate(getDateTime(scanner.Text())))
	}
	fmt.Println(timecount["8:40:55 PM"])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
