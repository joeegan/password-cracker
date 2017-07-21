package main

import (
  "fmt"
  "io/ioutil"
  "flag"
  "strings"
  "sort"
  "time"
  "log"
)

func check(e error) {
  if e != nil {
      panic(e)
  }
}

func getTopPasswords() ([]string) {
  dat, err := ioutil.ReadFile("./top-passwords.txt")
  check(err)
  passwords := strings.Split(string(dat), "\n")
  // Remove the last blank line
  return passwords[:len(passwords)-1]
}

func main() {
  target := flag.String("password", "", "Supply a password to calculate time to crack")
  flag.Parse()
  start := time.Now()
  topPasswords := getTopPasswords()
  sort.Strings(topPasswords)

  i := sort.SearchStrings(topPasswords, *target)
  if i < len(topPasswords) && topPasswords[i] == *target {
      elapsed := time.Since(start)
      log.Printf("Time taken to crack %s", elapsed)
      fmt.Printf("Found %s in a sorted top passwords list at index %d\n", topPasswords[i], i)
      return;
  }
  elapsed := time.Since(start)
  log.Printf("Couldn't crack %s", elapsed)

}
