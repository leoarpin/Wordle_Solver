package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// LoadWords lit un fichier texte (par défaut ./assets/words.txt)
// et renvoie une slice de mots de 5 lettres (en minuscules).
func LoadWords(path string) ([]string, error) {
	if path == "" {
		path = "assets/words.txt"
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	seen := make(map[string]struct{})
	out := make([]string, 0, 10000)

	sc := bufio.NewScanner(strings.NewReader(string(data)))
	for sc.Scan() {
		w := strings.TrimSpace(strings.ToLower(sc.Text()))
		if w == "" || strings.HasPrefix(w, "#") {
			continue
		}
		if utf8.RuneCountInString(w) != 5 {
			continue
		}
		if _, ok := seen[w]; !ok {
			seen[w] = struct{}{}
			out = append(out, w)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

// Score compare guess vs target et renvoie un motif Wordle (G=vert, Y=jaune, B=gris).
func Score(guess, target string) string {
	n := 5
	p := []byte("BBBBB")
	used := make([]bool, n)

	// Greens
	for i := 0; i < n; i++ {
		if guess[i] == target[i] {
			p[i] = 'G'
			used[i] = true
		}
	}
	// Yellows
	for i := 0; i < n; i++ {
		if p[i] == 'G' {
			continue
		}
		for j := 0; j < n; j++ {
			if !used[j] && guess[i] == target[j] && guess[j] != target[j] {
				p[i] = 'Y'
				used[j] = true
				break
			}
		}
	}
	return string(p)
}

// ValidatePattern vérifie que le motif est de 5 caractères parmi G/Y/B.
func ValidatePattern(p string) bool {
	if len(p) != 5 {
		return false
	}
	for i := 0; i < 5; i++ {
		switch p[i] {
		case 'G', 'Y', 'B':
		default:
			return false
		}
	}
	return true
}

var ErrInvalidPattern = errors.New("invalid pattern")

// FilterCandidates garde les candidats compatibles avec le motif donné.
func FilterCandidates(cands []string, guess, pattern string) []string {
	out := make([]string, 0, len(cands))
	for _, w := range cands {
		if Score(guess, w) == pattern {
			out = append(out, w)
		}
	}
	return out
}
