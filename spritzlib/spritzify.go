package spritzlib

import (
  "strings"
  "errors"
)

// Returns sequential paragraphs of PivotedWords from input text.
func SpritzifyText(input string) (spritzed_paragraphs [][]PivotedWord, err error) {
  input = prepForSplit(input)
  if len(input) == 0 {
    return nil, errors.New("Empty or whitespace-only string provided.")
  }
  paragraphs := strings.Split(input, "\n")
  spritzed_paragraphs = make([][]PivotedWord, 0, len(paragraphs))
  for _, para := range paragraphs {
    spr_para, err := spritzifyParagraph(para)
    if err != nil {
      return nil, err
    }
    spritzed_paragraphs = append(spritzed_paragraphs, spr_para)
  }
  return spritzed_paragraphs, nil
}

// Accept a contiguous length of words ONLY separated by spaces/punctuation and return
// a slice of pivoted words.
func spritzifyParagraph(paragraph string) ([]PivotedWord, error) {
    paragraph = strings.TrimSpace(paragraph)
    words := strings.Split(paragraph, " ")
    pivoted := make([]PivotedWord, 0, len(words))
    for _, w := range words {
      pw, err := Pivot(w)
      if err != nil {
        return nil, err
      }
      pivoted = append(pivoted, *pw)
    }
    return pivoted, nil
}

// Kill all weird space characters, convert tabs to spaces, then remove duplicate
// spaces. Remove "\r" entirely. Remove duplicate newlines.
// Return trimmed out outside space.
// Todo: Likely includes loads of allocations. Make more efficient using a byte-string
// based method? Do many subs at once using Regex?
func prepForSplit(input string) string {
  input = strings.TrimSpace(input)
  input = strings.Replace(input, "\r\n", "\n", -1)
  input = strings.Replace(input, "\n\r", "\n", -1)
  input = strings.Replace(input, "\r", "\n", -1)  // After removing couplets to avoid duplications.
  input = strings.Replace(input, "\v", "", -1)
  input = strings.Replace(input, "\f", "", -1)
  input = strings.Replace(input, "\t", " ", -1)
  input = strings.Replace(input, string(0xA0), " ", -1)
  input = strings.Replace(input, string(0x85), " ", -1)
  // Want to split up compound-words to compound- words, the trailing dash gets
  // kept/displayed later and scores additional pause points.
  input = strings.Replace(input, "-", "- ", -1)
  // Remove duplicate spaces.
  for {
    if !strings.Contains(input, "  ") {
      break
    }
    input = strings.Replace(input, "  ", " ", -1)
  }
  // Remove duplicate newlines.
  for {
    if !strings.Contains(input, "\n\n") {
      break
    }
    input = strings.Replace(input, "\n\n", "\n", -1)
  }
  return input
}
