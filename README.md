# 🟩 Wordle Solver en Go

Un solveur **Wordle** écrit en **Go**, capable de suggérer des mots et de réduire
progressivement les candidats en fonction des retours (motif `G/Y/B`).

- `G` = Green (bonne lettre, bonne position)
- `Y` = Yellow (bonne lettre, mauvaise position)
- `B` = Black (lettre absente)

---

## 🚀 Lancer le programme

### 1. Avec `go run`
à la racine du projet :
```bash
go run ./cmd/solver
```

Options disponibles :
- `-first` : affiche seulement la **première suggestion** et quitte.
- `-dict <path>` : utilise un dictionnaire personnalisé (par défaut `assets/words.txt`).

### 2. Avec L'éxécutable:

```bash
./wordle-solver
```

Avec options :
```bash
./wordle-solver -first
./wordle-solver -dict ./assets/words.txt
```

---

## 📝 Exemple d’utilisation (mode interactif)

```text
Wordle solver — motif: G=vert Y=jaune B=gris (ex: GYBBY). Tape 'exit' pour quitter.
Suggestion: slate
Motif ? BBYBG
→ 245 candidats restants
Suggestion: round
Motif ? GBBYB
→ 14 candidats restants
Suggestion: coral
Motif ? GGGGG
Solution probable: coral
```
