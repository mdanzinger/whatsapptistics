package main

import (
	"fmt"
	//"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

func main() {
	line := "[02:23:03 - 02/02/1997] NameOfApp - Unique log message, could be anyting....."

	re := regexp.MustCompile(`(\d{2}:\d{2}:\d{2}).*(\d{2}\/\d{2}\/\d{4})`)
	start := time.Now()
	m := re.FindStringSubmatch(line)

	fmt.Println(m[2])

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func getDateTime(line string) string {
	date := line[strings.Index(line, "[")+1 : strings.Index(line, "]")]
	return date
}

func getDate(dateTime string) string {
	date := strings.Split(dateTime, " - ")
	return date[1]
}

func getTime(dateTime string) string {
	date := strings.Split(dateTime, " - ")
	return date[0]
}
