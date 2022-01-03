package main

import(
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

func main() {
	
	defer fmt.Println("Bye!")
	
	var fileName string
	fmt.Scan(&fileName)
	
	tabooWords := make(map[string]struct{})
	
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	scanner.Split(bufio.ScanWords)
	
	for scanner.Scan() {
        tabooWords[scanner.Text()] = struct{}{}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	
	var line string
	
	for line != "exit"{
		fmt.Scan(&line)
		if line == "exit"{
			break
		}
		words := strings.Split(line, " ")
		for i := range words {
			if _, ok := tabooWords[strings.ToLower(words[i])]; ok {
				words[i] = strings.Repeat("*", len(words[i]))
			}	
		}
		
		for i := range words {
			fmt.Print(words[i], "")
		}
	}
}
