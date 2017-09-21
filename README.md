I made this to solve my apartment's monthly crossword contests.

This supports n x m crossword puzzles, but *NOT* ragged matrices.

```Usage: crossword.go <gridfile> <word bank file (optional)>```

Where gridfile is the crossword grid, formatted as a newline separated file, and word bank file are the words to search, formatted the same as gridfile.

Currently does not support upper/lowercase mixing.

Directions go from 1 -> 8 with North-East / Diagonal Up-Right assigned to 1, and wrapping clockwise until North / Up is at 8.
