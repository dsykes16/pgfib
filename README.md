PgFib
=====

Fibonacci algorithm implemented server-side on Postgres with a REST API server powered by Go

All Fibonacci calculations and memoization is done in pure PL/pgSQL to take full advantage of Postgres and its performance capabilities. The upper-limit is Fib(4786) which, even with a cold cache, is amazingly performant. With the benefit of a warm cache, Fib(500) completes recursively in <0.01s. See performance benchmarks in the **Performance** section below.

The API is quite crude at this point, but as only a single value is ever returned from a given endpoint, I didn't see the need to add in additional complexity. A programmers favorite acronym: YAGNI (You Aren't Gonna Need It). For a more elaborate API, look at https://github.com/dsykes16/gofib/ where I implemented a full gRPC API along with an abstracted memoization cache allowing for interchangable backends (e.g. local, postgres, mysql, redis, etc). Unfortunately the performance of that project was sub-par compared to the pure PL/pgSQL implementation in this project.

Usage
-----
The fastest way to get up and running is docker-compose. Just run `docker-compose up` and the API should be available on port `5000` (or whatever you change it to in the `docker-compose.yml` file).

Endpoints
---------
- GET /fib/`[n]`
  - Returns the N-th number in the Fibonacci sequence as plaintext number
- GET /size
  - Returns the current memoization cache size as a plaintext number
- GET /size/`[upper-bound]`
  - Returns the number of cached values in the exclusive range from 1 to `[upper-bound]`
- DELETE /cache
  - Wipes the memoization cache

Building
--------
Install Go 1.16+
Run:
```
make build
```

Optionally, to build a docker image, run:
```
make docker-build
```

Installing
----------
To install to common binary directory on \*nix systems:
NOTE: It is recommended to edit the values in the `app.env` file prior to installation. Otherwise changes can be made later by overriding settings via environment variables, editing the installed file in `/etc/pgfib/config/`, or editing the `app.env` file in the working directory and re-installing
```
make install
```

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
