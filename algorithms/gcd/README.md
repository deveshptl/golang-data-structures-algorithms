#### GCD - Greatest Common Divisor

Euclid of Alexandria outlined an algorithm for solving this problem in one of the volumes of his Elements most famous for its systematic exposition of geometry. In modern terms, Euclid’s algorithm is based on applying repeatedly the equality

<code style="text-align: center; display: block;">gcd(m, n) = gcd(n, m mod n)</code>

For e.g:

<code style="text-align: center; display: block;">gcd(60, 24) = gcd(24, 12) = gcd(12, 0) = 12</code>
