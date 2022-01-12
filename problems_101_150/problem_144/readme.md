# Project Euler Problem 144 - Investigating multiple reflections of a laser beam

[https://projecteuler.net/problem=144](https://projecteuler.net/problem=144)

## Solution idea

For this problem you need to know few trigonometric formulae.

Start with laser beam equation:
- it can be expressed as `y = k*x + b` where `k` is the slope of the beam - this is a tangent relative to `y = 0`
- assuming that beam goes though points `(x1, y1)` and `(x2, y2)` we can say that `k = (y2 - y1)/(x2 - x1)`

In the problem we are given that the slope of a tangent line to the ellipse at point `(x, y)` is `-4x/y` (tangent relative to `y = 0`)

We can do a few drawings to understand the following:
- consider laser beam tangent is `tanA = -(y2-y1)/(x2-x1)`
- consider that laser beam hits ellipse at point `(x2, y2)`
- tangent slope `tanB = -4*x2/y2`
- the beam will be reflected with angle `tan(A+B) = (tanA + tanB)/(1 - tanA*tanB)`

Next we need to understand what will be the equation of the reflecting laser beam. It is `y = tan(A+B)*x + (y2 - x2*tan(A+B))`.

We know that elipse and line cross at 2 points. So we need to solve the system of equations (that can be reduced to a quadratic equation in `x`):
- `4x^2 + y^2 = 100`
- `y = tan(A+B)*x + (y2 - x2*tan(A+B))` (`x2`, `y2` and `tan(A+B)` are calculated by this moment)

Solve for `x`. There should be 2 solutions - one is equal to `x2` and another one that will be next reflection point (this will be `x` cooridnate, `y` can be calculated using `y = tan(A+B)*x + (y2 - x2*tan(A+B))`).

Repeat this algorithm till the beam exits the ellipse

## Alternative solutions

## External links
