# Project Euler Solutions

![My profile](https://projecteuler.net/profile/sergeytrasko.png?)

https://projecteuler.net

## Why I solve Project Euler problems

First of all - it is fun. Some 20 years back I used to participate in programming competitions (i.e. solving algorithmic problems).
Since that I'm mostly working in enterprise environment and solving different types of problems. It's always good to shake your memory.

Secondly - it allows me to switch focus from my day-to-day work and refresh my former knowledge.

Thirdly - it forces me to discover new things for me, new algorithms, new approaches, etc.

And last, but not least - I do enjoy cracking these quite challenging problems.

## Things I learned solving Project Euler problems

- Programming language is not important, I chose golang but I could have achieved the same results with any other language.
- First 50 or so problems can be solved without knowing anything about algorithms - mostly brute-force solutions are viable. I think I wrote solutions for them during two or maybe three evenings.
- The bigger the problem number, the more complex they are. I.e. 5% problem from first hundred problems is much easier than 5% problem from 7-th hundred.
- Getting to top 1% (at the time of writing it is equivalent of solving 111 or more problems) is not that hard. It took me about 2 weeks to reach this milestone (mosly solving problems in the evenings)
- Useful tools to use:
  - https://www.alpertron.com.ar/QUAD.HTM - generic quadratic Diophantine equation solver - will be used in a number of problems
  - https://www.wolframalpha.com/ - Wolfram Alpha is a great tool to quickly investigate some functions and expressions, this might not give you solution but can help to verify your assumptions
- Nice techniques to master (they are used in many problems):
  - Dynamic programming - https://en.wikipedia.org/wiki/Dynamic_programming
  - Prime numbers - https://en.wikipedia.org/wiki/Prime_number
  - Euler's totient function - https://en.wikipedia.org/wiki/Euler%27s_totient_function
  - Fibonacci numbers - https://en.wikipedia.org/wiki/Fibonacci_number
  - Greatest common divisor - https://en.wikipedia.org/wiki/Greatest_common_divisor
  - Pythagorean Triple - https://en.wikipedia.org/wiki/Pythagorean_triple

## My strategy

- Don't look for other people solutions (as it turned out you can quite easily find the solution for almost all problems) - try to solve it myself first. It will be more fun than just copying the result from the others.
- Try to figure out algorithm complexity beforehand. Efficient algorithm for Project Euler problem should check at most 50'000'000 possibilities (to fit into recommended time limit of 1 minute). For example if I see that need to find sum of numbers below 10^7 - it gives a clue that every number should be checked in constant time (i.e. O(1)).
- One cannot know all algorithms - it's OK to google algorithms and ideas of solutions.
- Solution that runs over 10 seconds on a modern hardware is most likely sub-optimal. Project Euler states that the reasonable limit for the program to find a solution is one minute, but I think only a couple of my solutions run slower than 15 seconds. If I see that I go over this limit I spend some time figuring out what can be improved.
- After I solved the problem I try to google solutions of other people to verify my ideas and learn something from others.
- Project Euler opens you an internal discussion thread for a problem once you solve it - it's good to check other people ideas there.

## Solved problems

Here are my solved problems. For all of them source code is available plus some very brief explanation of my solution. 
Mostly breakdown does not explain each and every aspect of the solution - my goal is to provide you a hint how to address the problem.

- [#1 - Multiples of 3 or 5](problems_1_50/problem_1)
- [#2 - Even Fibonacci numbers](problems_1_50/problem_2)
- [#3 - Largest prime factor](problems_1_50/problem_3)
- [#4 - Largest palindrome product](problems_1_50/problem_4)
- [#5 - Smallest multiple](problems_1_50/problem_5)
- [#6 - Sum square difference](problems_1_50/problem_6)
- [#7 - 10001st prime](problems_1_50/problem_7)
- [#8 - Largest product in a series](problems_1_50/problem_8)
- [#9 - Special Pythagorean triplet](problems_1_50/problem_9)
- [#10 - Summation of primes](problems_1_50/problem_10)
- [#11 - Largest product in a grid](problems_1_50/problem_11)
- [#12 - Highly divisible triangular number](problems_1_50/problem_12)
- [#13 - Large sum](problems_1_50/problem_13)
- [#14 - Longest Collatz sequence](problems_1_50/problem_14)
- [#15 - Lattice paths](problems_1_50/problem_15)
- [#16 - Power digit sum](problems_1_50/problem_16)
- [#17 - Number letter counts](problems_1_50/problem_17)
- [#18 - Maximum path sum I](problems_1_50/problem_18)
- [#19 - Counting Sundays](problems_1_50/problem_19)
- [#20 - Factorial digit sum](problems_1_50/problem_20)
- [#21 - Amicable numbers](problems_1_50/problem_21)
- [#22 - Names scores](problems_1_50/problem_22)
- [#23 - Non-abundant sums](problems_1_50/problem_23) - no description
- [#24 - Lexicographic permutations](problems_1_50/problem_24)
- [#25 - 1000-digit Fibonacci number](problems_1_50/problem_25)
- [#26 - Reciprocal cycles](problems_1_50/problem_26)
- [#27 - Quadratic primes](problems_1_50/problem_27)
- [#28 - Number spiral diagonals](problems_1_50/problem_28)
- [#29 - Distinct powers](problems_1_50/problem_29)
- [#30 - Digit fifth powers](problems_1_50/problem_30)
- [#31 - Coin sums](problems_1_50/problem_31)
- [#32 - Pandigital products](problems_1_50/problem_32) - no description
- [#33 - Digit cancelling fractions](problems_1_50/problem_33) - no description
- [#34 - Digit factorials](problems_1_50/problem_34) - no description
- [#35 - Circular primes](problems_1_50/problem_35) - no description
- [#36 - Double-base palindromes](problems_1_50/problem_36) - no description
- [#37 - Truncatable primes](problems_1_50/problem_37) - no description
- [#38 - Pandigital multiples](problems_1_50/problem_38) - no description
- [#39 - Integer right triangles](problems_1_50/problem_39) - no description
- [#40 - Champernowne's constant](problems_1_50/problem_40) - no description
- [#41 - Pandigital prime](problems_1_50/problem_41) - no description
- [#42 - Coded triangle numbers](problems_1_50/problem_42) - no description
- [#43 - Sub-string divisibility](problems_1_50/problem_43) - no description
- [#44 - Pentagon numbers](problems_1_50/problem_44) - no description
- [#45 - Triangular, pentagonal, and hexagonal](problems_1_50/problem_45) - no description
- [#46 - Goldbach's other conjecture](problems_1_50/problem_46) - no description
- [#47 - Distinct primes factors](problems_1_50/problem_47) - no description
- [#48 - Self powers](problems_1_50/problem_48) - no description
- [#49 - Prime permutations](problems_1_50/problem_49) - no description
- [#50 - Consecutive prime sum](problems_1_50/problem_50) - no description
- [#51 - Prime digit replacements](problems_51_100/problem_51) - no description
- [#52 - Permuted multiples](problems_51_100/problem_52) - no description
- [#53 - Combinatoric selections](problems_51_100/problem_53) - no description
- [#54 - Poker hands](problems_51_100/problem_54) - no description
- [#55 - Lychrel numbers](problems_51_100/problem_55) - no description
- [#56 - Powerful digit sum](problems_51_100/problem_56) - no description
- [#57 - Square root convergents](problems_51_100/problem_57) - no description
- [#58 - Spiral primes](problems_51_100/problem_58) - no description
- [#59 - XOR decryption](problems_51_100/problem_59) - no description
- [#60 - Prime pair sets](problems_51_100/problem_60) - no description
- [#61 - Cyclical figurate numbers](problems_51_100/problem_61) - no description
- [#62 - Cubic permutations](problems_51_100/problem_62) - no description
- [#63 - Powerful digit counts](problems_51_100/problem_63) - no description
- [#64 - Odd period square roots](problems_51_100/problem_64) - no description
- [#65 - Convergents of e](problems_51_100/problem_65) - no description
- [#66 - Diophantine equation](problems_51_100/problem_66) - no description
- [#67 - Maximum path sum II](problems_51_100/problem_67) - no description
- [#68 - Magic 5-gon ring](problems_51_100/problem_68) - no description
- [#69 - Totient maximum](problems_51_100/problem_69) - no description
- [#70 - Totient permutation](problems_51_100/problem_70) - no description
- [#71 - Ordered fractions](problems_51_100/problem_71) - no description
- [#72 - Counting fractions](problems_51_100/problem_72) - no description
- [#73 - Counting fractions in a range](problems_51_100/problem_73) - no description
- [#74 - Digit factorial chains](problems_51_100/problem_74) - no description
- [#75 - Singular integer right triangles](problems_51_100/problem_75) - no description
- [#76 - Counting summations](problems_51_100/problem_76) - no description
- [#77 - Prime summations](problems_51_100/problem_77) - no description
- [#78 - Coin partitions](problems_51_100/problem_78)
- [#79 - Passcode derivation](problems_51_100/problem_79) - no description
- [#80 - Square root digital expansion](problems_51_100/problem_80) - no description
- [#81 - Path sum: two ways](problems_51_100/problem_81) - no description
- [#82 - Path sum: three ways](problems_51_100/problem_82) - no description
- [#83 - Path sum: four ways](problems_51_100/problem_83) - no description
- [#85 - Counting rectangles](problems_51_100/problem_85) - no description
- [#86 - Cuboid route](problems_51_100/problem_86) - no description
- [#87 - Prime power triples](problems_51_100/problem_87) - no description
- [#89 - Roman numerals](problems_51_100/problem_89) - no description
- [#90 - Cube digit pairs](problems_51_100/problem_90) - no description
- [#91 - Right triangles with integer coordinates](problems_51_100/problem_91) - no description
- [#92 - Square digit chains](problems_51_100/problem_92) - no description
- [#93 - Arithmetic expressions](problems_51_100/problem_93) - no description
- [#94 - Almost equilateral triangles](problems_51_100/problem_94) - no description
- [#95 - Amicable chains](problems_51_100/problem_95) - no description
- [#96 - Su Doku](problems_51_100/problem_96) - no description
- [#97 - Large non-Mersenne prime](problems_51_100/problem_97) - no description
- [#98 - Anagramic squares](problems_51_100/problem_98) - no description
- [#99 - Largest exponential](problems_51_100/problem_99) - no description
- [#100 - Arranged probability](problems_51_100/problem_100) - no description
- [#101 - Optimum polynomial](problems_101_150/problem_101) - no description
- [#102 - Triangle containment](problems_101_150/problem_102) - no description
- [#103 - Special subset sums: optimum](problems_101_150/problem_103) - no description
- [#104 - Pandigital Fibonacci ends](problems_101_150/problem_104) - no description
- [#105 - Special subset sums: testing](problems_101_150/problem_105) - no description
- [#107 - Minimal network](problems_101_150/problem_107) - no description
- [#108 - Diophantine reciprocals I](problems_101_150/problem_108) - no description
- [#109 - Darts](problems_101_150/problem_109) - no description
- [#110 - Diophantine reciprocals II](problems_101_150/problem_110)
- [#111 - Primes with runs](problems_101_150/problem_111)
- [#112 - Bouncy numbers](problems_101_150/problem_112) - no description
- [#113 - Non-bouncy numbers](problems_101_150/problem_113) - no description
- [#114 - Counting block combinations I](problems_101_150/problem_114) - no description
- [#116 - Red, green or blue tiles](problems_101_150/problem_116) - no description
- [#117 - Red, green, and blue tiles](problems_101_150/problem_117) - no description
- [#118 - Pandigital prime sets](problems_101_150/problem_118) - no description
- [#119 - Digit power sum](problems_101_150/problem_119) - no description
- [#120 - Square remainders](problems_101_150/problem_120) - no description
- [#121 - Disc game prize fund](problems_101_150/problem_121)
- [#122 - Efficient exponentiation](problems_101_150/problem_122) - no description
- [#123 - Prime square remainders](problems_101_150/problem_123) - no description
- [#124 - Ordered radicals](problems_101_150/problem_124) - no description
- [#125 - Palindromic sums](problems_101_150/problem_125) - no description
- [#127 - abc-hits](problems_101_150/problem_127)
- [#129 - Repunit divisibility](problems_101_150/problem_129)
- [#130 - Composites with prime repunit property](problems_101_150/problem_130)
- [#131 - Prime cube partnership](problems_101_150/problem_131)
- [#132 - Large repunit factors](problems_101_150/problem_132)
- [#133 - Repunit nonfactors](problems_101_150/problem_133)
- [#134 - Prime pair connection](problems_101_150/problem_134) - no description
- [#135 - Same differences](problems_101_150/problem_135) - no description
- [#136 - Singleton difference](problems_101_150/problem_136) - no description
- [#137 - Fibonacci golden nuggets](problems_101_150/problem_137)
- [#138 - Special isosceles triangles](problems_101_150/problem_138) - no description
- [#139 - Pythagorean tiles](problems_101_150/problem_139) - no description
- [#140 - Modified Fibonacci golden nuggets](problems_101_150/problem_140)
- [#142 - Perfect Square Collection](problems_101_150/problem_142) - no description
- [#145 - How many reversible numbers are there below one-billion?](problems_101_150/problem_145) - no description
- [#146 - Investigating a Prime Pattern](problems_101_150/problem_146)
- [#148 - Exploring Pascal's triangle](problems_101_150/problem_148)
- [#151 - Paper sheets of standard sizes: an expected-value problem](problems_151_200/problem_151) - no description
- [#162 - Hexadecimal numbers](problems_151_200/problem_162) - no description
- [#164 - Numbers for which no three consecutive digits have a sum greater than a given value](problems_151_200/problem_164) - no description
- [#166 - Criss Cross](problems_151_200/problem_166) - no description
- [#168 - Number Rotations](problems_151_200/problem_168)
- [#170 - Find the largest 0 to 9 pandigital that can be formed by concatenating products](problems_151_200/problem_170)
- [#172 - Investigating numbers with few repeated digits](problems_151_200/problem_172)
- [#173 - Using up to one million tiles how many different "hollow" square laminae can be formed?](problems_151_200/problem_173) - no description
- [#174 - Counting the number of "hollow" square laminae that can form one, two, three, ... distinct arrangements](problems_151_200/problem_174)
- [#178 - Step Numbers](problems_151_200/problem_178)
- [#179 - Consecutive positive divisors](problems_151_200/problem_179) - no description
- [#182 - RSA encryption](problems_151_200/problem_182)
- [#183 - Maximum product of parts](problems_151_200/problem_183) - no description
- [#187 - Semiprimes](problems_151_200/problem_187) - no description
- [#188 - The hyperexponentiation of a number](problems_151_200/problem_188)
- [#190 - Maximising a weighted product](problems_151_200/problem_190) - no description
- [#191 - Prize Strings](problems_151_200/problem_191) - no description
- [#195 - Inscribed circles of triangles with one angle of 60 degrees](problems_151_200/problem_195)
- [#197 - Investigating the behaviour of a recursively defined sequence](problems_151_200/problem_197) - no description
- [#203 - Squarefree Binomial Coefficients](problems_201_250/problem_203) - no description
- [#204 - Generalised Hamming Numbers](problems_201_250/problem_204) - no description
- [#205 - Dice Game](problems_201_250/problem_205) - no description
- [#206 - Concealed Square](problems_201_250/problem_206) - no description
- [#211 - Divisor Square Sum](problems_201_250/problem_211)
- [#214 - Totient Chains](problems_201_250/problem_214) - no description
- [#216 - Investigating the primality of numbers of the form 2n^2-1](problems_201_250/problem_216)
- [#225 - Tribonacci non-divisors](problems_201_250/problem_225)
- [#231 - The prime factorisation of binomial coefficients](problems_201_250/problem_231) - no description
- [#235 - An Arithmetic Geometric sequence](problems_201_250/problem_235) - no description
- [#243 - Resilience](problems_201_250/problem_243)
- [#265 - Binary Circles](problems_251_300/problem_265)
- [#279 - Triangles with integral sides and an integral angle](problems_251_300/problem_279)
- [#291 - Panaitopol Primes](problems_251_300/problem_291)
- [#303 - Multiples with small digits](problems_301_350/problem_303) - no description
- [#315 - Digital root clocks](problems_301_350/problem_315) - no description
- [#345 - Matrix Sum](problems_301_350/problem_345)
- [#346 - Strong Repunits](problems_301_350/problem_346)
- [#347 - Largest integer divisible by two primes](problems_301_350/problem_347) - no description
- [#348 - Sum of a square and a cube](problems_301_350/problem_348) - no description
- [#357 - Prime generating integers](problems_351_400/problem_357) - no description
- [#358 - Cyclic numbers](problems_351_400/problem_358)
- [#381 - (prime-k) factorial](problems_351_400/problem_381) - no description
- [#387 - Harshad Numbers](problems_351_400/problem_387)
- [#491 - Double pandigital number divisible by 11](problems_451_500/problem_491) - no description
- [#493 - Under The Rainbow](problems_451_500/problem_493)
- [#500 - Problem 500!!](problems_451_500/problem_500)
- [#504 - Square on the Inside](problems_501_550/problem_504) - no description
- [#518 - Prime triples and geometric sequences](problems_501_550/problem_518) - no description
- [#587 - Concave triangle](problems_551_600/problem_587) - no description
- [#650 - Divisors of Binomial Product](problems_601_650/problem_650) - no description
- [#686 - Powers of Two](problems_651_700/problem_686) - no description
- [#719 - Number Splitting](problems_701_750/problem_719) - no description
- [#751 - Concatenation Coincidence](problems_751_800/problem_751) - no description
- [#757 - Stealthy Numbers](problems_751_800/problem_757)
