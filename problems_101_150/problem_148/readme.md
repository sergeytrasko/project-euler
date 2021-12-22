# Project Euler Problem 148 - Exploring Pascal's triangle

[https://projecteuler.net/problem=148](https://projecteuler.net/problem=148)

## Solution idea

First obvious thing to mention (and to avoid huge numbers) - we don't need to calculate precise numbers in Pascal's triangle - instead it is sufficient to work `mod 7` - this can be easily checked. Even more we are not really interested in results of `mod 7` operation - it's ok to have `true`/`false` type of response.

Next thing I want to do - explore the triangle and display first rows. Drawing first 100 rows or so shows us quite clear pattern:
- obviously first 7 rows (starting from 0) are not divisible by 7 and they form a triangle of size 7
- for the next rows up to `7*7 = 49` we see that each 7 rows we add one more triangle. In total it forms a bigger triangle of size 49
- if we draw till `7*7*7 = 243` we notice that it has similar pattern, but with triangles of size 49
- and so on...

Playing around with this pattern I tried to convert row number to base 7 and see how we can use it.
Imagine we have row 18 (in base 10) that is 24 (in base 7). What it shows us? it shows us that there will be `2+1=3` triangles of 7 and each of them will have `4+1=5` non-divisible by 7 members. So in total in row 18 we will have `3*5=15` items we are interested in.
You can check the same for other rows as well.

So the algorithm is very simple:
- iterate from `0` to `10^9` and convert each row number to base 7.
- for each number in base 7 take each digit `d` and compose product of `1+d`
- sum the results

## Alternative solutions

There is a recursive solution as well.
Recursion anchor - if row number `n` is below 7 - just return `n` (see results of our exploration).
For row `n` we can find biggest `k` such that `7^k < n`. 
Set `mul = n / 7^k + 1` - this will be the number of triangles in the row.
Then we need recursively find the answer for row `n mod 7^k` and multiply it by `mul`.

Basically it's quite similar to converting row number `n` to base 7

## External links
