package main

// BestGuess : heuristique simple par fréquence de lettres uniques.
func BestGuess(cands []string) string {
	if len(cands) == 0 {
		return ""
	}
	// fréquence des lettres (présence unique par mot)
	freq := map[byte]int{}
	for _, w := range cands {
		seen := [26]bool{}
		for i := 0; i < 5; i++ {
			ch := w[i]
			if ch < 'a' || ch > 'z' {
				continue
			}
			idx := ch - 'a'
			if !seen[idx] {
				freq[ch]++
				seen[idx] = true
			}
		}
	}
	best := cands[0]
	bestScore := -1
	for _, w := range cands {
		seen := [26]bool{}
		score := 0
		for i := 0; i < 5; i++ {
			ch := w[i]
			if ch < 'a' || ch > 'z' {
				continue
			}
			idx := ch - 'a'
			if !seen[idx] {
				score += freq[ch]
				seen[idx] = true
			}
		}
		if score > bestScore {
			best = w
			bestScore = score
		}
	}
	return best
}
