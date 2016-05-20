package main_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "bytes"
  main "imdbSearcher"
)

func setup() {
  main.ActorNames = []string{}
}

func Test_AskForName(t *testing.T) {
  setup()
  a := assert.New(t)
  b := []byte("Lee Majors\n")
  r := bytes.NewBuffer(b)
  main.AskForName(r)
  a.Equal(len(main.ActorNames), 1)
  a.Equal(main.ActorNames[0], "Lee Majors")
}

func Test_AskForNames(t *testing.T) {
  setup()
  a := assert.New(t)
  b := []byte("Lee\nMajors\nn\n")
  r := bytes.NewBuffer(b)
  main.AskForNames(r)
  a.Equal(len(main.ActorNames), 2)
  a.Equal(main.ActorNames[0], "Lee")
  a.Equal(main.ActorNames[1], "Majors")
}

func Test_AskForNames_FourNames(t *testing.T) {
  setup()
  a := assert.New(t)
  b := []byte("Lee\nMajors\ny\nMary\ny\nSmith\nn\n")
  r := bytes.NewBuffer(b)
  main.AskForNames(r)
  a.Equal(len(main.ActorNames), 4)
  a.Equal(main.ActorNames[0], "Lee")
  a.Equal(main.ActorNames[1], "Majors")
  a.Equal(main.ActorNames[2], "Mary")
  a.Equal(main.ActorNames[3], "Smith")
}
