# Project Euler Problem 195 - Inscribed circles of triangles with one angle of 60 degrees

[https://projecteuler.net/problem=195](https://projecteuler.net/problem=195)

## Solution idea

Integer triangles with 60 degree angle are called Eisenstein triples and they can be generated similarly to Pythagorean triples.
Note that we should ignore equilateral triangles (i.e. ones where `a = b = c`).

Next important fact is that radius of inscribed circle in a triangle depends on the triangle dimensions: `Area = r * (a+b+c)/2`. In the same time `Area` can be calculated using Heron's formula: `Area = sqrt(s * (s-a) * (s-b) * (s-c))` where `s = (a+b+c)/2`.

So the algorithm will be:
- generate all Eisenstein triples up to certain bound (I chose it after several attempts). Note that `a`, `b` and `c` can have common divisor of either 1 or 3
- calculate area of the triange
- calculate radius of inscribed circle
- good news is that the triples are generated in ascending order based on radius not taking `gcd(a, b, c)` into account - this gives us nice loop break condition
- last step - detect how many times we can scale radius still to have integer triange

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Eisenstein_triple
https://en.wikipedia.org/wiki/Integer_triangle#Integer_triangles_with_one_angle_with_a_given_rational_cosine
https://en.wikipedia.org/wiki/Incircle_and_excircles_of_a_triangle#Relation_to_area_of_the_triangle