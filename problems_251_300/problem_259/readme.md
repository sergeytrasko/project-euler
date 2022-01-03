# Project Euler Problem 259 - Reachable Numbers

[https://projecteuler.net/problem=259](https://projecteuler.net/problem=259)

## Solution idea

Problem is somehow similar to [problem 93 - Arithmetic expressions](../problems_51_100/problem93) and I borrowed some code from it.

The idea is that we are given a string (or array) of digits from 1 to 9 and we try to generate all possible arithmetic expressions.

For instance, the sample from them problem can be represented with the following arithmetic expression tree.
```
                   (*)
                  /   \
                 /     \
                /       \
               /         \
              /           \
             /             \
           (*)              \
          /   \              \
         /     \              \
        /       \              \
       /         \              \
      /           \              \
     /             \              \
   (/)             (-)            (-)
  /   \           /   \          /   \
 /    ( )       (*)    \       ( )    \
/    /   \     /   \    \     /   \    \
1    2    3    4    5    6    7    8    9
```

This can be done using 'divide and conquer' approach.

Imagine we have array `digits`, two markers `left` and `right` and a function `calc(digits, left, right)` that returns all possible outcomes using `digits` between `left` and `right`

Then for each index `i` between `left` and `right` we can recursively calculate two sub-results:
- `calc(digits, left, i)`
- `calc(digits, i+1, right)`

Basically we are splitting the problem into 2 smaller problems.

Once we get results of both sub-problems we can combine the results and check all possible combinations (add, substract, multiply and divide).

Few implementation details:
- for performance optimizations we are dealing only with unique results (e.g. it is possible to get same sub-result in multiple ways)
- to avoid decimal overflow and precision issues we can emulate fractions using a `struct` and reduce it when necessary

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Divide-and-conquer_algorithm