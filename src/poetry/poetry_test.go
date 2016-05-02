package poetry

import (
  "testing"
)

func TestNumStanzas(t *testing.T) {
    p := Poem{{"And from my pillow, looking forth by light",
    "Of moon or favouring stars, I could behold", "The antechapel where the statue stood", "Of Newton with his prism and silent face,", "The marble index of a mind for ever", "Voyaging through strange seas of Thought, alone."}}

    if p.NumStanzas() != 1 {
      t.Fatalf("Unexpected stanza count %d", p.NumStanzas())
    }

    emptyPoem := Poem{}

    if emptyPoem.NumStanzas() != 0 {
      t.Fatalf("Empty poem is not empty %d", p.NumStanzas())

    }
}

func TestNumLines(t *testing.T) {
    p := Poem{{"And from my pillow, looking forth by light",
    "Of moon or favouring stars, I could behold", "The antechapel where the statue stood", "Of Newton with his prism and silent face,", "The marble index of a mind for ever", "Voyaging through strange seas of Thought, alone."}}

    if p.NumLines() != 6 {
      t.Fatalf("Unexpected stanza count %d", p.NumStanzas())
    }

    emptyPoem := Poem{}

    if emptyPoem.NumStanzas() != 0 {
      t.Fatalf("Empty poem is not empty %d", p.NumStanzas())

    }
}

func TestStats(t *testing.T) {
    emptyPoem := Poem{}
    v, c := emptyPoem.Stats()

    if v != 0 || c != 0 {
      t.Fatalf("Bad number of vowels or consonants")
    }

    p := Poem{{"Hello"}}
    v, c = p.Stats()

    if v != 2 || c != 3 {
      t.Fatalf("Bad number of vowels or consonants")
    }

    p = Poem{{"Hello, world!"}}
    v, c = p.Stats()

    if v != 3 || c != 7 {
      t.Fatalf("Bad number of vowels or consonants v: %d c: %d", v,c)
    }

}

