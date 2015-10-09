package spritzlib

import (
  "time"
  "fmt"
  "strconv"
)

func clearConsole() {
  // There is certainly a way to do this that doesn't suck like this.
  fmt.Print("\r                                                  \r")
}

func pivotPrint(wd string, bold bool, pivot_colour, plain_colour, background_colour string) {
    piv, _ := Pivot(wd)
    printColours(piv, bold, pivot_colour, plain_colour, background_colour)
}

// Split and spritz text, then render at intended WPM using
func RenderSpritzed(input string, wpm int, bold bool, pivot_colour, plain_colour, background_colour string) (int, int, int, error) {
    mpw := 1/wpm
    spritzed_paragraphs, err := SpritzifyText(input)
    if err != nil {
      return 0,0,0,err
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
    began := time.Now()
    total_wds := 0
    for _, para := range spritzed_paragraphs {
      for _, wd := range para {
        total_wds++
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
    took := time.Since(began)
    took_seconds := int(time.Since(began).Seconds())
    final_wpm := int(float64(total_wds)/took.Minutes())
    return took_seconds, total_wds, final_wpm, nil
}
