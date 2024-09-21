//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Count struct {
	count int
	sync.Mutex
}

func getWords(line string) []string {
	words := strings.Split(line, " ")
	return words
}

func countLetters(word string) int {
	letters := 0
	for _, letter := range word {
		if unicode.IsLetter(letter) {
			letters++
		}
	}
	return letters
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup
	letterCounter := Count{}
	for {
		if scanner.Scan() {
			line := scanner.Text()
			words := getWords(line)
			for _, word := range words {
				wordCopy := word
				wg.Add(1)
				go func(word string) {
					letterCounter.Lock()
					defer letterCounter.Unlock()
					defer wg.Done()
					letterCounter.count += countLetters(word)
				}(wordCopy)
			}
		} else {
			break
		}
	}
	wg.Wait()
	letterCounter.Lock()
	defer letterCounter.Unlock()
	println("Total letters counted:", letterCounter.count)

}
