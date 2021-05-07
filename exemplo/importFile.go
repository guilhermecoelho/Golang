package exemplo

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Readfile() []string {
	file, err := os.Open("../files/dup.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	slice := []string{}
	includeLine := false
	for scanner.Scan() {
		if includeLine {
			slice = append(slice, scanner.Text())
		}
		if strings.Contains(scanner.Text(), "Data mov.") {
			includeLine = true
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return slice
}
