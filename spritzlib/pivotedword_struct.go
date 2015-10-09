package spritzlib

import (
  "errors"
  "strings"
  "time"
  "unicode"
  "github.com/wsxiaoys/terminal"
)

type PivotedWord struct{
  Fore, Mid, Aft string
  IndentOffset int
  DelayScore int
}


func Pivot(word string) (pivoted *PivotedWord, err error) {
    var bestLetter int
    switch len(word) {
      case 0: {
        err = errors.New("Given zero-length string..")
        return
      }
      case 1: {
        bestLetter = 1
      }
      case 2,3,4,5: {
        bestLetter = 2
      }
      case 6,7,8,9: {
        bestLetter = 3
      }
      case 10,11,12,13: {
        bestLetter = 4
      }
      default: {
        bestLetter = 5
      }
    }
    delay := delayPercent(word)
    offset := 6 - bestLetter
    return &PivotedWord{ Fore: word[:bestLetter-1],
                         Mid: word[bestLetter-1:bestLetter],
                         Aft: word[bestLetter:],
                         IndentOffset: offset,
                         DelayScore: delay}, nil
}

// Provide an integer score for delay-value of a word.
// A "word" gets a boost for:
// * Ending in punctuation marks.
// * Being longer than N letters
func delayPercent(word string) int {
    wordScore := 0
    rword := []rune(word)
    clearword := make([]rune, 0, len(rword))
    for _, r := range rword {
      if unicode.IsLetter(r) || unicode.IsNumber(r) {
        clearword = append(clearword, r)
      }
    }
    if unicode.IsPunct(rword[len(rword)-1]) {
      wordScore = wordScore + 2
    }
    if unicode.IsPunct([]rune(rword)[0]) {
      wordScore = wordScore + 2
    }
    if len(clearword) > 8 {
      wordScore = wordScore + 1
    }
    if len(clearword) > 12 {
      wordScore = wordScore + 1
    }
    return 100 + (10 * wordScore)
}

func printColours(wd *PivotedWord, embolden bool, pivot_colour, plain_colour, background string) {
  if embolden {
    pivot_colour = "!"+pivot_colour
    plain_colour = "!"+plain_colour
  }
  pivot_colour = pivot_colour + strings.ToUpper(background)
  plain_colour = plain_colour + strings.ToUpper(background)
  terminal.Stdout.ClearLine().
    Color(plain_colour).Print(strings.Repeat(" ", wd.IndentOffset)).
    Color(strings.ToUpper(background)).
    Color(plain_colour).Print(wd.Fore).
    Color(pivot_colour).Print(wd.Mid).
    Color(plain_colour).Print(wd.Aft).
    Color(plain_colour).Print(strings.Repeat(" ", 20-(wd.Length()%20))).
    Reset()
}

func (self *PivotedWord) wpmDelay(wpm int) <-chan time.Time {
  MPW := ((time.Minute/100) * time.Duration(self.DelayScore)) / time.Duration(wpm)
  return time.After(MPW)
}

func (self *PivotedWord) Length() int {
  return len(self.Fore) + len(self.Mid) + len(self.Aft)
}
