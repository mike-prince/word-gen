package main

import (
  "fmt"
  "math/rand"
  "time"
  "os"
  "bufio"
  "sort"
)


// Read words from file

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}


// Get random word from given slice

func getRandomWord(words []string) string {
  rand.Seed(time.Now().UTC().UnixNano())
  return words[rand.Intn(len(words))]
}


// Sort array by length

type ByLen []string

func (a ByLen) Len() int {
   return len(a)
}

func (a ByLen) Less(i, j int) bool {
   return len(a[i]) < len(a[j])
}

func (a ByLen) Swap(i, j int) {
   a[i], a[j] = a[j], a[i]
}


// Main function

func main() {

  words, _ := readLines("strings.txt")

  var result []string

  for i := 0; i < 100000; i++ {

    wordOne := getRandomWord(words)
    wordTwo := getRandomWord(words)

    if wordOne == wordTwo {
      wordTwo = getRandomWord(words)
    }

    var newWord = wordOne + wordTwo

    if len(newWord) <= 10 {
      result = append(result, newWord)
    }
  }

  sort.Sort(ByLen(result))

  for i := 0; i < len(result); i++ {
    fmt.Printf("%s\n", result[i])
  }
}
