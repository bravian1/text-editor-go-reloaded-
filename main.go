package main
import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)
func main() {
	// command line arguments.
	args := os.Args
	// Open the input file specified in the command line arguments.
	inputFile, err := os.Open(args[1])
	if err != nil {
		// log error and exit.
		log.Fatal(err)
	}
	// Ensure the file is closed when the function exits.
	defer inputFile.Close()
	// Create a new Scanner to read from the input file.
	scanner := bufio.NewScanner(inputFile)
	// Create the output file specified in the command line arguments.
	content, err := os.Create(args[2])
	if err != nil {
		log.Fatal(err)
	}
	// Ensure the file is closed when the function exits.
	defer content.Close()
	// Read the input file line by line.
	for scanner.Scan() {
		// Get the current line as a string.
		line := scanner.Text()
		// Split the line into words.
		words := strings.Fields(line)
		// Iterate over each word in the line.
		for i := 0; i < len(words); i++ {
			if words[i] == "(hex)" && i > 0 {
				hexnum := words[i-1]
				decnum, err := strconv.ParseInt(hexnum, 16, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(decnum, 10)
				}
				words = append(words[:i], words[i+1:]...)
			} else if words[i] == "(bin)" && i > 0 {
				binnum := words[i-1]
				decnum, err := strconv.ParseInt(binnum, 2, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(decnum, 10)
				}
				words = append(words[:i], words[i+1:]...)
			} else if words[i] == "(up)" {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
			} else if words[i] == "(low)" {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
			} else if words[i] == "(cap)" && i > 0 {
				words[i-1] = strings.ToUpper(string(words[i-1][0])) + words[i-1][1:]
				words = append(words[:i], words[i+1:]...)
			} else if words[i] == "(up," && i > 0 {
				num := extractNum(words[i+1])
				for j := 1; j <= num; j++ {
					a := (i - j)
					if a >= 0 && a < len(words) {
						words[a] = strings.ToUpper(words[a])
					}
				}
				words = append(words[:i], words[i+2:]...)
			} else if words[i] == "(low," && i > 0 {
				num := extractNum(words[i+1])
				for j := 1; j <= num; j++ {
					a := (i - j)
					if a >= 0 && a < len(words) {
						words[a] = strings.ToLower(words[a])
					}
				}
				words = append(words[:i], words[i+2:]...)
			} else if words[i] == "(cap," && i < len(words)-1 {
				num := words[i+1]
				b := extractNum(num)
				for j := 1; j <= b; j++ {
					a := (i - j)
					if a >= 0 && a < len(words) {
						words[a] = strings.ToUpper(string(words[a][0])) + words[a][1:]
					}
				}
				words = append(words[:i], words[i+2:]...)
			} else if words[i] == "a" && Vowel(words[i+1][0]) && i < len(words) {
				words[i] = "an"
			} else if words[i] == "A" && Vowel(words[i+1][0]) && i < len(words) {
				words[i] = "An"
			}
		}
		// Adjust punctuations.
		words = Punctuation(words)
		// Join the words back into a single string with spaces.
		newline := strings.Join(words, " ")
		// Write the processed line to the output file.
		_, err := content.WriteString(newline + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	// Check for errors that may have occurred during scanning.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
// extractNum extracts all numeric characters from a string and converts them to an integer. eg (cap, 10) -> 10
func extractNum(s string) int {
	numstr := "" // Initialize an empty string to hold the numeric characters.
	for _, i := range s {
		// Check if the character is a digit
		if i >= '0' && i <= '9' {
			// If it's a digit, append it to the numstr string.
			numstr += string(i)
		}
	}
	// Convert the concatenated string of digits to an integer.
	result, _ := strconv.Atoi(numstr)
	return result
}
func Vowel(s byte) bool {
	// Convert the byte to a lowercase string for comparison.
	char := strings.ToLower(string(s))
	// Return true if 'char' is a vowel or 'h', false otherwise.
	return char == "a" || char == "e" || char == "i" || char == "o" || char == "u" || char == "h"
}
// Punctuation adjusts the position of punctuation in a slice of strings.
func Punctuation(s []string) []string {
	// Define a slice of punctuation marks.
	puncs := []string{",", ".", ":", ";", "!", "?"}
	// Move leading punctuation to the end of the previous word.
	for i, word := range s {
		for _, punc := range puncs {
			// If the word starts with a punctuation mark and does not end with it,
			// remove the punctuation from the start and add it to the end of the previous word.
			if strings.HasPrefix(word, punc) && !strings.HasSuffix(word, punc) {
				s[i] = word[1:]
				s[i-1] += punc
			}
		}
	}
	// Handle words that are just a punctuation mark at the end of the slice.
	for i, word := range s {
		for _, punc := range puncs {
			// If the word is at the end of the slice and is just a punctuation mark,
			// add it to the previous word and remove it from the slice.
			if strings.HasPrefix(word, punc) && strings.HasSuffix(word, punc) && i == len(s)-1 {
				s[i-1] += word
				s = s[:len(s)-1]
				break
			}
		}
	}
	// Handle words that are just a punctuation mark not at the end of the slice.
	for i, word := range s {
		for _, punc := range puncs {
			// If the word is not at the end and is just a punctuation mark,
			// add it to the previous word and remove it from the slice.
			if strings.HasPrefix(word, punc) && strings.HasSuffix(word, punc) && s[i] != s[len(s)-1] {
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	// Handle the first apostrophe in the slice.
	// Initialize a variable to check for the first apostrophe.
	set := 0
	for i, word := range s {
		// If the word is an apostrophe and it's the first one,
		// add it to the next word and remove it from the slice.
		if word == "'" {
			if set == 0 && i+1 < len(s) {
				set++
				s[i+1] = word + s[i+1]
				s = append(s[:i], s[i+1:]...)
			} else if set == 1 && i-1 >= 0 {
				s[i-1] = s[i-1] + word
				s = append(s[:i], s[i+1:]...)
				set = 0
			}
		}
		//  Handle any remaining apostrophes.
		// If the word is an apostrophe,
		// add it to the previous word and remove it from the slice.
	}
	// Return the modified slice.
	return s
}
