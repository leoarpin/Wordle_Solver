package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	dictPath := flag.String("dict", "", "path to a custom words.txt (optional)")
	firstOnly := flag.Bool("first", false, "print only the first suggested guess and exit")
	flag.Parse()

	// Load dictionary (either external file or embedded)
	words, err := LoadWords(*dictPath) // from finder.go
	if err != nil {
		log.Fatalf("dictionary load error: %v", err)
	}
	if len(words) == 0 {
		log.Fatal("dictionary is empty")
	}

	cands := make([]string, len(words))
	copy(cands, words)

	if *firstOnly {
		fmt.Println(BestGuess(cands)) // from stats.go
		return
	}

	fmt.Println("Wordle solver — motif: G=vert Y=jaune B=gris (ex: GYBBY). Tape 'exit' pour quitter.")
	for {
		guess := BestGuess(cands)
		if guess == "" {
			fmt.Println("Plus de candidats.")
			return
		}
		fmt.Printf("Suggestion: %s\nMotif ? ", guess)

		var pattern string
		_, err := fmt.Fscan(os.Stdin, &pattern)
		if err != nil {
			log.Fatalf("read: %v", err)
		}
		if pattern == "exit" {
			return
		}
		if !ValidatePattern(pattern) {
			fmt.Println("Motif invalide (5 caractères parmi G/Y/B).")
			continue
		}

		cands = FilterCandidates(cands, guess, pattern)
		switch len(cands) {
		case 0:
			fmt.Println("Aucun candidat restant (vérifie le motif/dictionnaire).")
			return
		case 1:
			fmt.Printf("Solution probable: %s\n", cands[0])
			return
		default:
			fmt.Printf("→ %d candidats restants\n", len(cands))
		}
	}
}
