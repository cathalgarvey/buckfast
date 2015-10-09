# Buckfast - Spritz on Go
A quick and dirty spritz-style speed-reading tool for colourised terminals. This might even work for Windows in the CMD prompt, if you're a masochist?

By Cathal Garvey, Copyright October 2015, released GNU AGPL, inclusive of the included (terrible) `spritzlib` back-end.

## Usage
```
# Print bold text with yellow pivot
buckfast ReadMe.md --bold --pivot-colour=yellow
# Print white text, white pivot, magenta background
buckfast ReadMe.md --pivot-colour=white --background-colour=magenta
# Red pivot, black text, white background, bold text, 800 words per minute.
buckfast --pivot-colour=red --background-colour=white --plain-colour=black --bold --wpm 800 ReadMe.md
# The above, but with shortcodes
buckfast -p red -P black -b white -w 800 --bold ReadMe.md
```

Suggestion: Find your favourite reading style, then make an alias of it in your `.bashrc` file.

## Thanks to
* Inspiration from [glance.wtf](http://glance.wtf).
* Flag parsing from [kingpin](https://github.com/alecthomas/kingpin).
* Concept from [Spritz](http://www.spritzinc.com/the-science/).
