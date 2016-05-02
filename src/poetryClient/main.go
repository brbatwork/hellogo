package main

import (
    "fmt"
    "poetry"
)

func main() {
  p := poetry.Poem{{"And from my pillow, looking forth by light",
    "Of moon or favouring stars, I could behold", "The antechapel where the statue stood", "Of Newton with his prism and silent face,", "The marble index of a mind for ever", "Voyaging through strange seas of Thought, alone."}}
  v ,c, punc := p.Stats()
  fmt.Printf("Vowels: %d, Consonants: %d Punc: %d\n", v, c, punc)
  fmt.Printf("Stanzas: %d, Lines: %d\n", p.NumStanzas(), p.NumLines())
  fmt.Printf("%s\n",p)

}
