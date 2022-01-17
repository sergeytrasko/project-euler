# Project Euler Problem 479 - Roots on the Rise

[https://projecteuler.net/problem=479](https://projecteuler.net/problem=479)

## Solution idea

Simplify original equation and convert `1/x = (k/x)^2*(k+x^2)-kx` to `kx^3-k^2x^2+x-k^3=0`.

Now consider that `a`, `b` and `c` are roots of this expression.

According to Vieta's formula we get:
- `a+b+c = k`
- `ab+bc+ca = 1/k`
- `abc = k^2`

Notice that:
```
1 = k*(1/k) = (a+b+c)*(ab+bc+ca) = 
  a^2b + abc + a^2c + 
  ab^2 + b^2c + abc +
  abc + bc^2 + c^2a = 
```

Then expand `(a+b)*(b+c)*(c+a)`:

```
(a+b)*(b+c)*(c+a) = 
  (ab+ac+b^2+bc)*(c+a) = 
  (abc + ac^2 + b^2c + bc^2) + (a^2b + a^2c + b^2a + abc)
```

But notice that it is equal (see calculation above) to `(a+b)*(b+c)*(c+a) = (a+b+c)*(ab+bc+ca) - abc = 1 - k^2`.

This gives us: `(a+b)^p*(b+c)^p*(c+a)^p = [(a+b)*(b+c)*(c+a)]^p = (1-k^2)^p`

Sum of `s(n) = sum{p = 1..n} (1-k^2)^p` can be calulated a sum of geomeric series: `s(n) = sum{p = 1..n} (1-k^2)^p = (1 - (1-k^2)^(n+1))/k^2`

And now it comes to how to calculate this value under `mod 1'000'000'007`. Here modular multiplicative inverse operation is helpful:

```
s(n) = [
    (1 - (1-k^2)^(n+1)) mod m)
    * (k^2)^-1 mod m
  ] mod m
```

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Vieta%27s_formulas

https://en.wikipedia.org/wiki/Geometric_series

https://en.wikipedia.org/wiki/Modular_multiplicative_inverse