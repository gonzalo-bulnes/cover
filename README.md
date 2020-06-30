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

#### Step 1

- [x] Define a notion of `distance.ForErrorCorrection`.
- [ ] Figure out if some property can be applied to all items (e.g. words) of a group (e.g. nearest neighbours of a given word).

### Usage

```sh
go run github.com/gonzalo-bulnes/cover
```
