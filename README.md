# Go-reloaded

Outil en Go de complétion, d'édition et d'autocorrection de texte.

## Utilisation

```
go run . sample.txt result.txt
```

Le programme lit le fichier d'entrée (`sample.txt`), applique les transformations,
puis écrit le résultat dans le fichier de sortie (`result.txt`).

## Transformations

- `(hex)` : convertit le mot précédent (hexadécimal) en décimal
- `(bin)` : convertit le mot précédent (binaire) en décimal
- `(up)`, `(low)`, `(cap)` : majuscules, minuscules, initiale en majuscule
- `(up, 2)`, `(low, 3)`, `(cap, 2)` : même chose sur les N mots précédents
- Ponctuation `. , ! ? : ;` : collée au mot précédent, suivie d'un espace
- Groupes `...` et `!?` : conservés collés au mot précédent
- Apostrophes `' ... '` : collées au texte qu'elles entourent
- `a` devient `an` devant une voyelle ou un `h`

## Exemples

| Entrée | Sortie |
|---|---|
| `1E (hex) files were added` | `30 files were added` |
| `It has been 10 (bin) years` | `It has been 2 years` |
| `Ready, set, go (up) !` | `Ready, set, GO!` |
| `Welcome to the Brooklyn bridge (cap)` | `Welcome to the Brooklyn Bridge` |
| `This is so exciting (up, 2)` | `This is SO EXCITING` |
| `I was thinking ... You were right` | `I was thinking... You were right` |
| `I am exactly how they describe me: ' awesome '` | `I am exactly how they describe me: 'awesome'` |
| `There it was. A amazing rock!` | `There it was. An amazing rock!` |

## Tests

```
go test ./...
```

## Structure du projet

```
go-reload/
├── go.mod         module Go
├── main.go        code du programme
├── main_test.go   tests unitaires
├── sample.txt     exemple d'entrée
└── README.md      ce fichier
```

## Limites connues

La ponctuation est traitée par règle simple : un nombre décimal comme `3.14`
serait transformé en `3. 14`. Ce cas n'est pas couvert par le sujet.
