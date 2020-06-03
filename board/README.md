Package board
=============

The board keeps track of the current state of the Sudoku puzzle.

It has a variety of functions and methods that initialize and solve the sudoku puzzle.

Naked Singles
-------------
If a number shows up as the only candidate in a cell, you may set that cell value to be that number.
[Example](http://www.sudokusnake.com/nakedsingles.php)

Hidden Singles
--------------
If a number only shows up as a candidate in one cell of a row, column or box, you may set that cell value to be that number.
[Example](http://www.sudokusnake.com/hiddensingles.php)

Naked Pairs
-----------
If there are exactly two candidates in two cells in the same block (row, column or box) and the candidates are identical, then these candidates can be removed from the other cells in the block.
[Example](https://www.learn-sudoku.com/naked-pairs.html)

Pointing Pairs
--------------
If a candidate is present in only two cells of a box, then it must be the solution for one of these two cells. If these two cells belong to the same row or column, then this candidate can not be the solution in any other cell of the same row or column, respectively.
[Example](http://www.taupierbw.be/SudokuCoach/SC_PointingPair.shtml)

Claiming Pairs
--------------
Our implementation of pointing pairs cover this case. The difference between the two strategies is that we switch the box and row (or column) for pointing pair.

X-Wings
-------
If you can find a row that contains the same pencil mark in exactly two cells, as well as another parallel row that mirrors it – containing the same pencil mark in only the same two spots, then you can use this information to eliminate similar pencil marks in the columns passing through those spots. The rows and columns can be switched.
[Example](https://www.learn-sudoku.com/x-wing.html)

Hidden Pairs
------------
If there are atleast two candidates in two cells in the same block (row, column or box) and the candidates are identical, then the remaining candidates in those cells can be removed.
[Example](https://www.learn-sudoku.com/hidden-pairs.html)

Naked Triplets
--------------
This is similar to Naked Pair with 3 cells.

Not implemented

Naked Quads
-----------
This is similar to Naked Pair with 4 cells.

Not implemented

Coloring
--------
A chain of conjugate pairs exists when multiple conjugate pairs are linked together. For this technique, conjugate pairs, all on the same value, are linked together and the cells in the tree are colored with two alternating colors. This means cells from a single conjugate pair are never colored the same.
[Example](https://sudoku.ironmonger.com/howto/simpleColoring/docs.tpl?setreferrertoself=true)

Not implemented

Brute Force (Contradiction)
-----------
If an assumed final value leads to a non-valid sudoku, then the assumed value can be removed from the list of candidates.

Not implemented
