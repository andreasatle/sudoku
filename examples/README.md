Example Sudokus
===============

A library of test examples for sudokus.

The program reads in a parameter-file on the format:

```
9 3 3
123456789

8-7-9-6-2
95286--4-
3-6-2-598
781934256
264-----9
5396-2--4
6-----421
12--4698-
4-821--65
```

```
6 3 2
123456

--36--
-2---4
5---6-
-3---5
3---1-
--14--
```

```
9 3 3
ABCDEFGHI

E---H--DI
---E---C-
-FGC----A
AE-------
---B-H---
-------AH
G----DAE-
-C---B---
DI--E---C
```

The first line contains the board size, and number of row-boxes and column-boxes. The second line contains all the values. There has to be exactly board size different values. Last is the initial board. The dash-character means no value, and can't be part of the values.