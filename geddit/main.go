package main

import (
  "fmt"
  "os"
  "log"
  "github.com/jboursiquot/reddit"
)

func main(){

  if len(os.Args[1:]) == 0 {
    log.Fatal("Subreddit is required")
  }

  items, err := reddit.Get(os.Args[1])

  if err != nil {
    log.Fatal(err)
  }

  for _, item := range items {
    fmt.Println(item)
    fmt.Println()
  }
}

