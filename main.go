package main

import (
  "io/ioutil"
  "log"
  "strings"
  "github.com/alecthomas/kingpin"
  "github.com/cathalgarvey/buckfast/spritzlib"
  "github.com/cathalgarvey/buckfast/scrapedia"
)

var (
  infile = kingpin.Arg("infile", "File to read. If prefixed with 'wikipedia:', fetches the corresponding wikipedia page instead (needs to be exact target page title)").Required().String()
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
  var (
    bcontent []byte
    content string
    err error
  )
  kingpin.Parse()
  if strings.HasPrefix(*infile, "wikipedia:") {
    log.Println("Fetching from wikipedia:", strings.TrimPrefix(*infile, "wikiedia:"))
    content, err = scrapedia.GetMainText("https://en.wikipedia.org/wiki/"+strings.TrimPrefix(*infile, "wikipedia:"))
  } else {
    // Default: assume plaintext file and load contents.
    log.Println("Opening plaintext file: "+*infile)
    bcontent, err = ioutil.ReadFile(*infile)
    content = string(bcontent)
  }
  if err != nil {
    log.Fatal(err.Error())
  }
  log.Println("Beginning at", *wpm, "words per minute..")
  took_seconds, total_wds, final_wpm, err := spritzlib.RenderSpritzed(string(content), *wpm, *boldText, colourNameToCode(*pivotColour), colourNameToCode(*plainColour), colourNameToCode(*bgColour))
  if err != nil {
    log.Fatal(err.Error())
  }
  log.Println("Finished reading", total_wds, "words in", took_seconds, "seconds")
  log.Println("Final wpm:", final_wpm)
}
