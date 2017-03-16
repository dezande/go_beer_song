package main

import "testing"

func TestVerse4(t *testing.T) {
  s := verse(4)
  s_t := "4 bottles of beer on the wall, 4 bottles of beer.\n" +
  "Take one down and pass it around, 3 bottles of beer on the wall.\n"
  if s != s_t {
    t.Error("Wrong song\n -", s_t, "\n\n+", s)
  }
}

func TestVerse99(t *testing.T) {
  s := verse(99)
  s_t := "99 bottles of beer on the wall, 99 bottles of beer.\n" +
      "Take one down and pass it around, 98 bottles of beer on the wall.\n"
  if s != s_t {
    t.Error("Wrong song\n -", s_t, "\n\n+", s)
  }
}
