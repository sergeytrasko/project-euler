# Project Euler Problem 279 - Triangles with integral sides and an integral angle

[https://projecteuler.net/problem=279](https://projecteuler.net/problem=279)

## Solution idea

For a triangle knowing two sides `a` and `b` and angle `alpha` between them it's possible to calculate third side `c` using the following equation (Law of Cosines): `c^2 = a^2 + b^2 - 2*a*b*cos(alpha)` (note if `alpha` is 90 degrees then `cos(alpha)` is 0 and it gives us Pythagorean theorem).
We know that `a^2`, `b^2` and `c^2` are integers. That means that `2*a*b*cos(alpha)` is also an integer. As `2*a*b` is integer then `cos(alpha)` should be at least a rational number.

According to Niven's theorem, the only rational values for `sin(alpha)` and `cos(alpha)` are: `0`, `1/2`, `-1/2`, `1` and `-1`. Applying it to triangle angle we can conclude that `cos(alpha)` can be only `1/2` (60 degrees), `0` (90 degrees) and `-1/2` (120 degrees).

Let's look at these 3 cases:
- `alpha = 90 degrees` and `cos(alpha) = 0` - this gives us right triange and we need to calculate number of Pythagorean tripples with a perimeter below `10^8`
- `alpha = 60 degrees` and `cos(alpha) = 1/2` - this converts to `c^2 = a^2 + b^2 - ab` that we need to solve over integers
- `alpha = 120 degrees` and `cos(alpha) = -1/2` - this converts to `c^2 = a^2 + b^2 + ab` that we need to solve over integers

Actually, integer triangles with 60 (this one is called Einstein's tripple) and 120 degrees have a generator formula similar to Pythagorean tripples.
We just need to carefully implement it.

## Alternative solutions

Naive implementation works well for very small numbers of perimeter sum. But it can be handy to check your main algorithm on smaller cases.

## External links

https://en.wikipedia.org/wiki/Niven%27s_theorem
https://en.wikipedia.org/wiki/Pythagorean_triple
https://en.wikipedia.org/wiki/Law_of_cosines
https://en.wikipedia.org/wiki/Eisenstein_triple
https://en.wikipedia.org/wiki/Integer_triangle#Integer_triangles_with_one_angle_with_a_given_rational_cosine