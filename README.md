# golang-levenshtein
==========
Provide 2 main functions:
- Get 2 strings edit distance: LevenshteinDistance
- Get 2 strings edit operation list: LevenshteinEditOps

## Installation

```bash
$ go get github.com/david601508/golang-levenshtein
```

### Quick Start

Diffs are configured with Unified (or ContextDiff) structures, and can
be output to an io.Writer or returned as a string.

```Go
  ops := LevenshteinEditOps("Come on", "Came on")
  editCount := LevenshteinDistance("Come on", "Came on")
```


