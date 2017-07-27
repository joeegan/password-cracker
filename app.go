package main

import (
  "fmt"
  "io/ioutil"
  "flag"
  "strings"
  "sort"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func getList(filename string) ([]string) {
  data, err := ioutil.ReadFile("./" + filename + ".txt")
  check(err)
  words := strings.Split(string(data), "\n")
  // Remove the last blank line
  trimmed := words[:len(words)-1]
  sort.Strings(trimmed)
  return trimmed
}

type queryInfo struct {
  i int
  listName string
}

func findInList(channel chan queryInfo, target string, list []string, listName string) () {
  i := sort.SearchStrings(list, target)
  if i < len(list) && list[i] == target {
    channel <- queryInfo{ i: i, listName: listName }
    close(channel)
  }
  return;
}

func main() {
  target := flag.String("password", "", "Supply a password to calculate time to crack")
  flag.Parse()
  
  passwords := getList("passwords")
  firstnames := getList("firstnames")
  surnames := getList("surnames")
  channel := make(chan queryInfo)

  go findInList(channel, *target, passwords, "passwords")
  go findInList(channel, *target, firstnames, "firstnames")
  go findInList(channel, *target, surnames, "surnames")

  msg1 := <-channel
  fmt.Println("Found password", *target, "at index", msg1.i, "in list", msg1.listName)

}
