# Project Euler Solutions

![My profile](https://projecteuler.net/profile/sergeytrasko.png)

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
- [#23 - Non-abundant sums](problems_1_50/problem_23) - to add description
- [#24 - Lexicographic permutations](problems_1_50/problem_24)
- [#25 - 1000-digit Fibonacci number](problems_1_50/problem_25)
- [#26 - Reciprocal cycles](problems_1_50/problem_26)
- [#27 - Quadratic primes](problems_1_50/problem_27)
- [#28 - Number spiral diagonals](problems_1_50/problem_28)
- [#29 - Distinct powers](problems_1_50/problem_29)
- [#30 - Digit fifth powers](problems_1_50/problem_30)
- [#31 - Coin sums](problems_1_50/problem_31)

More to come!!!
- [# - ](problems_1_50/problem_)
