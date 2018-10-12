package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	line := "[2017-11-21, 8:40:55 PM] Martin: If i can find a ride from one of my friends ill go"

	dateTime := getDateTime(line)
	thedate := getDate(dateTime)
	thetime := getTime(dateTime)
	thename := getName(line)

	fmt.Println(thedate)
	fmt.Println(thetime)
	fmt.Println(thename)

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

func getName(line string) string {
	name := line[strings.Index(line, "]")+2 : strings.Index(line, ": ")]
	return name
}
