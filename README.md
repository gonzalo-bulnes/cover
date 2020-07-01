Cover
=====

An exploration of the cover tree structure applied to the composition of wordlists suitable to create memorable passphrases.

Spike: create a cover tree of words
-----------------------------------

### Goal

Understsand the most basic use of [go-covertree][gct] and [golang-levenshtein][gl].

  [gct]: https://github.com/mandykoh/go-covertree
  [gl]: https://github.com/texttheater/golang-levenshtein

### Usage

```sh
go run github.com/gonzalo-bulnes/cover
```

Spike: proof of concept (edit distance)
---------------------------------------

### Goal

Understand if the cover tree can be used to pre-select groups of words that have a given edit distance between them (all of them). That would could be used to allow **automated error correction** across a wordlist.

> [In the second short list, all] words are at least an edit distance of 3 apart. This means that future software could correct any single typo in the user's passphrase (and in many cases more than one typo). (source: [EFF's New Wordlists for Random Passphrases](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases))

#### Steps

- [x] Define a notion of `distance.ForErrorCorrection`.
- [ ] Figure out if some property can be applied to all items (e.g. words) of a group (e.g. nearest neighbours of a given word).

### Usage

```sh
PRINT=true go run github.com/gonzalo-bulnes/cover
```

### Scratchpad

Typical constraints (to consider, unless they're at the corpus level):

- no word is an exact prefix of any other
- edit distance >= 3 between all the words
- words are as short as reasonable
- words are as recognized / memorable as possible (corpus)

Random thoughts:

- the edit distance increases when a word is longer than the other
- ideally for a passphrase wordlist, shorter words tend to be better (ignoring the words meaning for now)
- it may be that targetting "at minimum 3 and as close as possible to 3" as the edit distance constraint could be interesting
- it could be interesting to be able to modify independently the influence of the word length and the similarity with other words of same length


1. insert all word from corpus into tree
2. find 7776 nearest from some word
3. find the 7776 nearest from each word in that set
4. if there are any less than 7776 nearest, exclude the word that is being used as reference
5. and start again!

- https://en.wikipedia.org/wiki/Soundex
- https://wordnet.princeton.edu/

Cover tree for checking candidates!

1. insert all word from corpus into tree
2. find 7776 nearest from some word
3. insert the first word into a new cover tree (output tree)
4. query that output tree using the first word in the list of neighbours, inseet it (it will fit)
5. query the output tree using the second neigho bour, check that it results in all two previous words, and insert. Otherwise reject.
6. Iterate with the third neighbour etc.
