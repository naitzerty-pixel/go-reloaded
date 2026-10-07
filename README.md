# Go-reloaded

Outil en Go de complétion, d'édition et d'autocorrection de texte.

## Utilisation

    go run . sample.txt result.txt

## Transformations

- (hex), (bin) : conversion en décimal
- (up), (low), (cap), avec ou sans nombre : (up, 2)
- ponctuation : . , ! ? : ; ... !?
- apostrophes : ' ... '
- a -> an devant une voyelle ou un h

## Tests

    go test ./...
