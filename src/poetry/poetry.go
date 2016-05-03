package poetry

import (
    "unicode"
    "fmt"
    "os"
    "strings"
    "bufio"
)
type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem() Poem {
  return Poem{}
}

func LoadPoem(location string) (Poem, error) {
  f ,err := os.Open(location)
  if err != nil {
    return nil, err
  }

  defer f.Close()

  p := Poem{}
  var s Stanza
  scan := bufio.NewScanner(f)

  for scan.Scan() {
    l := scan.Text()

    if l == "" {
      p = append(p, s)
      s = Stanza{}
      continue
    }

    s = append(s, Line(l))
  }

  p = append(p, s)

  if scan.Err() != nil {
    return nil, err
  }

  return p, nil
}

func (p Poem) NumStanzas() int {
  return len(p)
}

func (s Stanza) NumLines() int {
  return len(s)
}

func (p Poem) NumLines() (count int) {
  for _, s := range p {
    count += s.NumLines()
  }
  return
}

func (p Poem) NumWords() int {
  results := 0
  for _, s := range p {
    for _, l := range s {
      sl := string(l) // cast Line to a string excplicitly
      parts := strings.Split(sl, " ")
      results += len(parts)
    }
  }
  return results
}

func (p Poem) Stats() (numVowels, numConsonants int, numPuncs int) {
  for _, stan := range p {
    for _, line := range stan {
      for _, rune := range line {
        if unicode.IsPunct(rune) {
          numPuncs += 1
        } else {
          switch rune {
          case 'a', 'e', 'i', 'o', 'u':
            numVowels += 1
          default:
            numConsonants += 1
          }
        }
      }
    }
  }

  return
}

func (s Stanza) String() string {
    result := ""
    for _, l := range s {
      result += fmt.Sprintf("%s\n", l)
    }
    return result
}

func (p Poem) String() string {
  result := ""

  for _, s := range p {
    result +=  fmt.Sprintf("%s\n", s)
  }

  return result
}
