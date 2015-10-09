package spritzlib

import (
  "time"
  "fmt"
  "strconv"
)

func clearConsole() {
  // There is certainly a way to do this that doesn't suck like this.
  fmt.Print("\r                              \r")
}

func pivotPrint(wd string, bold bool, pivot_colour, plain_colour, background_colour string) {
    piv, _ := Pivot(wd)
    printColours(piv, bold, pivot_colour, plain_colour, background_colour)
}

// Split and spritz text, then render at intended WPM using
func RenderSpritzed(input string, wpm int, bold bool, pivot_colour, plain_colour, background_colour string) error {
    mpw := 1/wpm
    spritzed_paragraphs, err := SpritzifyText(input)
    if err != nil {
      return err
    }
    for i := 3; i>0 ; i-- {
      clearConsole()
      pivotPrint("Ready.."+strconv.Itoa(i), bold, pivot_colour, plain_colour, background_colour)
      <- time.After(time.Second)
    }
    clearConsole()
    pivotPrint("Go!", bold, pivot_colour, plain_colour, background_colour)
    <- time.After(time.Second)
    clearConsole()
    spinup := false
    for _, para := range spritzed_paragraphs {
      for _, wd := range para {
        clearConsole()
        if spinup {
          // If last word had a positive delay, increment this word's delay
          // to help eye catch up.
          wd.DelayScore = wd.DelayScore + 1
          spinup = false
        } else if wd.DelayScore > 0 {
          spinup = true
        }
        printColours(&wd, bold, pivot_colour, plain_colour, background_colour)
        <- wd.wpmDelay(wpm)
      }
      clearConsole()
      fmt.Print("---")
      <- time.After(time.Minute * time.Duration(mpw))
      clearConsole()
    }
    clearConsole()
    fmt.Println("Done!")
    return nil
}
