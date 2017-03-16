package main

import "fmt"

func lyric() string {
  s := "smdlfkmlsdfk msldfkùqmsdfk"
  return s
}

func verse(n int) string {
  s := "%d bottles of beer on the wall, %d bottles of beer.\n" +
       "Take one down and pass it around, %d bottles of beer on the wall.\n"
  return fmt.Sprintf(s, n, n, n - 1)
}

func main() {
  fmt.Println(verse(4))
}
