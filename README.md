## Generate FIRST/FOLLOW/PREDICT Set from BNF.

### Feature
* [x] FirstSet generate. Output pretty.
* [x] FollowSet generate. Output pretty.
* [x] LL(1) Predicate Parsing Table. Output pretty.

### Usage
#### You can use demo input `cd cmd && go run main.go`
<details>
  <summary>Output to the terminal default</summary>

   ```
   FirstSet:
FIRST(D) = {g ε f}
FIRST(E) = {g ε}
FIRST(F) = {f ε}
FIRST(S) = {a}
FIRST(B) = {c}
FIRST(C) = {b ε}

FollowSet:
FOLLOW(F) = {h}
FOLLOW(S) = {$}
FOLLOW(B) = {g h f}
FOLLOW(C) = {g h f}
FOLLOW(D) = {h}
FOLLOW(E) = {f h}

PredictTable:
+---+----------------+------------+------------+------------+------------+------------+---+
| # | a              | b          | c          | f          | g          | h          | $ |
+---+----------------+------------+------------+------------+------------+------------+---+
| S | S -> {a B D h} |            |            |            |            |            |   |
+---+----------------+------------+------------+------------+------------+------------+---+
| B |                |            | B -> {c C} |            |            |            |   |
+---+----------------+------------+------------+------------+------------+------------+---+
| C |                | C -> {b C} |            | C -> {ε}   | C -> {ε}   | C -> {ε}   |   |
+---+----------------+------------+------------+------------+------------+------------+---+
| D |                |            |            | D -> {E F} | D -> {E F} | D -> {E F} |   |
+---+----------------+------------+------------+------------+------------+------------+---+
| E |                |            |            | E -> {ε}   | E -> {g}   | E -> {ε}   |   |
+---+----------------+------------+------------+------------+------------+------------+---+
| F |                |            |            | F -> {f}   |            | F -> {ε}   |   |
+---+----------------+------------+------------+------------+------------+------------+---+
   ```
</details>

#### Or use your own bnf grammar. `cd cmd && go run main.go -grammar your_own_grammar_file`

### Note
1. use `ε` indicate `EPSILON` (unicode is `'\u03B5'`)
2. use `$` indicate `input right end marker`.
3. use UpperCase letter indicate `Nonterminal`
4. use lowerCase letter indicate `Terminal`
5. `BNF` format with `|` support alternate.
6. use `->` distinguish `LHS` and `RHS`

More demo see the [`cmd/demo.bnf`](cmd/demo.bnf), or [`testdata`](testdata/testdata2.json)

### Ref
* [FirstFollow](https://www.cs.uaf.edu/~cs331/notes/FirstFollow.pdf)
* [Compiler Construction](https://learning.oreilly.com/library/view/compiler-construction/9789332524590/)
* [Parsing Topics](http://www.mollypages.org/page/grammar/index.mp#generaldeter) **Note**: some example is confused(some letter is not printed in the web page. so maybe confused when reading it.)
* [Online calculate FIRST/FOLLOW/PREDICT](http://hackingoff.com/compilers/predict-first-follow-set).
