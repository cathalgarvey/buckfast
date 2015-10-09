package main

import (
  "io/ioutil"
  "log"
  "github.com/alecthomas/kingpin"
  "github.com/cathalgarvey/buckfast/spritzlib"
)

var (
  infile = kingpin.Arg("infile", "file to read in.").Required().String()
  wpm = kingpin.Flag("wpm", "Approximate words per minute").Default("400").Short('w').Int()
  pivotColour = kingpin.Flag("pivot-colour", "Preferred pivot colour, one of red, green, blue, yellow, cyan, magenta, white, black.").
    Default("green").
    Short('p').
    Enum("red", "green", "blue", "cyan", "magenta", "yellow", "black", "white")
    //String()
  plainColour = kingpin.Flag("plain-colour", "Preferred non-pivot text colour, one of red, green, blue, yellow, cyan, magenta, white, black.").
    Default("white").
    Short('P').
    Enum("red", "green", "blue", "cyan", "magenta", "yellow", "black", "white")
//    String()
  bgColour = kingpin.Flag("background-colour", "Text background colour, one of red, green, blue, yellow, cyan, magenta, white, black.").
    Default("").
    Short('b').
    Enum("red", "green", "blue", "cyan", "magenta", "yellow", "black", "white")
//    String()
  boldText = kingpin.Flag("bold", "Whether to print bold.").Bool()
)

// Return single-letter codes required by terminal library.
func colourNameToCode(name string) string {
  switch name {
  case "red":     return "r"
  case "green":   return "g"
  case "blue":    return "b"
  case "cyan":    return "c"
  case "magenta": return "m"
  case "yellow":  return "y"
  case "black":   return "k"
  case "white":   return "w"
  default:        return ""  // Inputs are already enum'd so shouldn't happen.
  }
}

func main() {
  kingpin.Parse()
  log.Println("Opening infile: "+*infile)
  log.Println("Bold: ", *boldText)
  content, err := ioutil.ReadFile(*infile)
  if err != nil {
    log.Fatal(err.Error())
  }
  err = spritzlib.RenderSpritzed(string(content), *wpm, *boldText, colourNameToCode(*pivotColour), colourNameToCode(*plainColour), colourNameToCode(*bgColour))
  if err != nil {
    log.Fatal(err.Error())
  }
}
