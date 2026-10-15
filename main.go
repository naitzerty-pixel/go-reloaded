package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	cmdRe   = regexp.MustCompile(`\(\s*(hex|bin|up|low|cap)\s*(?:,\s*(\d+)\s*)?\)`)
	tokenRe = regexp.MustCompile(`^<<(\w+),(\d*)>>$`)
	spaceRe = regexp.MustCompile(`\s+([.,!?:;]+)`)
	afterRe = regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;])`)
)

func capitalize(w string) string {
	r := []rune(strings.ToLower(w))
	if len(r) == 0 {
		return w
	}
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func apply(words []string, cmd string, n int) {
	if cmd == "hex" || cmd == "bin" {
		if len(words) == 0 {
			return
		}
		base := 16
		if cmd == "bin" {
			base = 2
		}
		last := len(words) - 1
		if v, err := strconv.ParseInt(words[last], base, 64); err == nil {
			words[last] = strconv.FormatInt(v, 10)
		}
		return
	}
	for i := len(words) - 1; i >= 0 && n > 0; i, n = i-1, n-1 {
		switch cmd {
		case "up":
			words[i] = strings.ToUpper(words[i])
		case "low":
			words[i] = strings.ToLower(words[i])
		case "cap":
			words[i] = capitalize(words[i])
		}
	}
}

func applyCommands(line string) string {
	line = cmdRe.ReplaceAllString(line, " <<$1,$2>> ")
	var out []string
	for _, w := range strings.Fields(line) {
		m := tokenRe.FindStringSubmatch(w)
		if m == nil {
			out = append(out, w)
			continue
		}
		n := 1
		if m[2] != "" {
			n, _ = strconv.Atoi(m[2])
		}
		apply(out, m[1], n)
	}
	return strings.Join(out, " ")
}

func fixQuotes(line string) string {
	var b strings.Builder
	open, noSpace := false, false
	for _, w := range strings.Fields(line) {
		if w == "'" {
			if open {
				b.WriteString("'")
				open = false
				noSpace = false
			} else {
				if b.Len() > 0 {
					b.WriteString(" ")
				}
				b.WriteString("'")
				open = true
				noSpace = true
			}
			continue
		}
		if b.Len() > 0 && !noSpace {
			b.WriteString(" ")
		}
		b.WriteString(w)
		noSpace = false
	}
	return b.String()
}

func fixPunctuation(line string) string {
	line = spaceRe.ReplaceAllString(line, "$1")
	return afterRe.ReplaceAllString(line, "$1 $2")
}

func fixArticles(line string) string {
	words := strings.Fields(line)
	for i := 0; i < len(words)-1; i++ {
		if words[i] != "a" && words[i] != "A" {
			continue
		}
		first := unicode.ToLower([]rune(words[i+1])[0])
		if strings.ContainsRune("aeiouh", first) {
			if words[i] == "a" {
				words[i] = "an"
			} else {
				words[i] = "An"
			}
		}
	}
	return strings.Join(words, " ")
}

// Process transforme tout le texte, ligne par ligne.
func Process(text string) string {
	lines := strings.Split(text, "\n")
	for i, l := range lines {
		l = strings.TrimRight(l, "\r")
		l = applyCommands(l)
		l = fixQuotes(l)
		l = fixPunctuation(l)
		l = fixArticles(l)
		lines[i] = l
	}
	return strings.Join(lines, "\n")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . input.txt output.txt")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("erreur de lecture :", err)
		os.Exit(1)
	}
	if err := os.WriteFile(os.Args[2], []byte(Process(string(data))), 0o644); err != nil {
		fmt.Println("erreur d'écriture :", err)
		os.Exit(1)
	}
}
