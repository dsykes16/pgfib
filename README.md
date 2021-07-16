PgFib
=====

Fibonacci algorithm implemented server-side on Postgres with a REST API server powered by Go

Performance
-----------
```
• [MEASUREMENT]
Fibonacci Tests
  Fib(500) performance with cold cache 

  Ran 5 samples:
  runtime:
    Fastest Time: 0.032s
    Slowest Time: 0.059s
    Average Time: 0.040s ± 0.010s
------------------------------
• [MEASUREMENT]
Fibonacci Tests
  Fib(500) performance with warm cache 

    Ran 5 samples:
    runtime:
      Fastest Time: 0.001s
      Slowest Time: 0.001s
      Average Time: 0.001s ± 0.000s
------------------------------
```
