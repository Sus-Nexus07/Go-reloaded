# go-reloaded

A command-line text completion, editing, and auto-correction tool written in Go. It reads an input file, applies a series of smart transformations, and writes the cleaned result to an output file.

**Author:** Richard Obodo — https://acad.learn2earn.ng/git/robodo

---

## Table of Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Transformation Rules](#transformation-rules)
6. [Examples](#examples)
7. [Project Structure](#project-structure)
8. [How It Works](#how-it-works)
9. [Packages Used](#packages-used)

---

## Overview

`go-reloaded` is a smart text processor. You give it a file containing special tags and formatting problems, and it outputs a corrected, clean version. It handles number base conversions, word casing, article corrections (`a` → `an`), punctuation spacing, and quote formatting — all in one pass.

---

## Features

| Feature | Description |
|---|---|
| Hex conversion | Converts hexadecimal numbers to decimal |
| Binary conversion | Converts binary numbers to decimal |
| Uppercase | Converts word(s) to UPPERCASE |
| Lowercase | Converts word(s) to lowercase |
| Capitalize | Capitalizes the first letter of word(s) |
| Article fix | Changes `a` to `an` before vowels and `h` |
| Punctuation fix | Fixes spacing around `. , ! ? : ;` |
| Quote fix | Removes extra spaces inside single-quote pairs |

---

## Installation

**Prerequisites:** Go 1.18 or higher.

```bash
# Clone the repository
git clone https://acad.learn2earn.ng/git/robodo/go-reloaded

# Navigate into the project
cd go-reloaded

# Verify it builds
go build .
```

---

## Usage

```bash
go run . <input_file> <output_file>
```

| Argument | Description |
|---|---|
| `<input_file>` | Path to the text file you want to process |
| `<output_file>` | Path where the corrected text will be saved (created if it doesn't exist) |

**Example:**

```bash
go run . sample.txt result.txt
```

---

## Transformation Rules

### Number Conversions

`(hex)` — Replaces the word immediately before it with its decimal equivalent, treating the word as hexadecimal.

```
"1E (hex) files were added"  →  "30 files were added"
```

`(bin)` — Replaces the word immediately before it with its decimal equivalent, treating the word as binary.

```
"It has been 10 (bin) years"  →  "It has been 2 years"
```

---

### Word Casing

`(up)` — Converts the word immediately before the tag to UPPERCASE.

```
"Ready, set, go (up) !"  →  "Ready, set, GO!"
```

`(low)` — Converts the word immediately before the tag to lowercase.

```
"I should stop SHOUTING (low)"  →  "I should stop shouting"
```

`(cap)` — Capitalizes the word immediately before the tag.

```
"Welcome to the Brooklyn bridge (cap)"  →  "Welcome to the Brooklyn Bridge"
```

**With a number** — Adding `, <number>` to any casing tag applies the transformation to that many preceding words.

```
"This is so exciting (up, 2)"             →  "This is SO EXCITING"
"it was the age of foolishness (cap, 6)"  →  "It Was The Age Of Foolishness"
"IT WAS THE (low, 3) winter of despair"   →  "it was the winter of despair"
```

---

### Article Correction

Every standalone `a` or `A` is changed to `an` or `An` when the next word begins with a vowel (`a e i o u`) or the letter `h`.

```
"There it was. A amazing rock!"                         →  "There it was. An amazing rock!"
"There is no greater agony than bearing a untold story" →  "...bearing an untold story"
"It was a hour ago"                                     →  "It was an hour ago"
```

---

### Punctuation Spacing

The punctuation marks `. , ! ? : ;` must always sit directly after the previous word (no space before), with exactly one space after them.

```
"I was sitting over there ,and then BAMM !!"  →  "I was sitting over there, and then BAMM!!"
```

Groups of punctuation like `...` or `!?` are kept together as one unit.

```
"I was thinking ... You were right"                       →  "I was thinking... You were right"
"Punctuation tests are ... kinda boring ,what do you think ?"  →  "Punctuation tests are... kinda boring, what do you think?"
```

---

### Quote Formatting

Single quotes `'` always come in pairs. Any spaces between the quote mark and the word(s) inside are removed.

```
"I am exactly how they describe me: ' awesome '"
→  "I am exactly how they describe me: 'awesome'"

"As Elton John said: ' I am the most well-known homosexual in the world '"
→  "As Elton John said: 'I am the most well-known homosexual in the world'"
```

---

## Examples

### Full sample run

**Input (`sample.txt`):**

```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```

**Command:**

```bash
go run . sample.txt result.txt
```

**Output (`result.txt`):**

```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

### Number conversion

**Input:**
```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```

**Output:**
```
Simply add 66 and 2 and you will see the result is 68.
```

---

## Project Structure

```
go-reloaded/
├── main.go               # Entry point — argument validation and file I/O
├── go.mod                # Go module definition
├── README.md             # This file
└── processor/
    ├── pipeline.go       # Process() entry point and transformation order
    ├── hexbin.go         # (hex) and (bin) conversion
    ├── case_modifiers.go # (up), (low), (cap) transformations
    ├── articles.go       # a → an correction
    ├── punctuations.go   # Punctuation and quote spacing
    └── processor_test.go # Unit tests for all transformations
```

---

## How It Works

Transformations are applied in a fixed order. Each step completes before the next begins, so no stage interferes with another's patterns.

```
Read file
    │
    ▼
1. HandleHexBin       →  Convert (hex) and (bin) markers to decimal
    │
    ▼
2. HandleCaseModifiers →  Apply (up), (low), (cap) tags
    │
    ▼
3. FixPunctuation     →  Fix spacing around punctuation and quotes
    │
    ▼
4. FixArticles        →  Fix "a" → "an" before vowels and h
    │
    ▼
Write file
```

---

## Packages Used

Only Go standard library packages — no external dependencies.

| Package | Purpose |
|---|---|
| `os` | Reading and writing files, accessing command-line arguments |
| `fmt` | Printing error messages to the terminal |
| `regexp` | Finding and replacing text patterns (hex tags, casing tags, punctuation, etc.) |
| `strconv` | Converting between strings and numbers (`ParseInt`, `Atoi`) |
| `strings` | String utilities: `ToUpper`, `ToLower`, `Fields`, `TrimSpace`, `Join`, etc. |

---

## License

Built as part of the go-reloaded curriculum project at [Learn2Earn](https://learn2earn.ng).