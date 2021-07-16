-- TABLE: fibonacci: Memoization table for Fibonacci numbers
--   Columns:
--     ind - index of number in fibonacci sequence
--     result - term of fibonacci sequence at given index
CREATE TABLE IF NOT EXISTS fibonacci (
    ind integer PRIMARY KEY,
    result decimal(1000,0) NOT NULL
);


-- FUNC: get_fib(n integer) (res bigint)               
--   Parameters:
--     n - term of fibonacci sequence to return
--   Returns:
--     res - nth number in fibonacci seqence
CREATE OR REPLACE FUNCTION get_fib(n integer) RETURNS decimal(1000,0) AS $$
DECLARE
    res decimal(1000,0);
BEGIN
    IF n < 2 THEN
        RETURN n;
    END IF;

    SELECT INTO res result
    FROM fibonacci
    WHERE ind = n;

    IF res IS NULL THEN
        res := get_fib(n-1) + get_fib(n-2);
        INSERT INTO fibonacci (ind, result)
        VALUES (n, res);
    END IF;
    RETURN res;
END;
$$ LANGUAGE plpgsql;


-- FUNC: get_cache_size(end integer) (size integer)
--   Parameters:
--     range_end(integer) - End of range for measuring cache size
--   Returns:
--     size(integer) - Size of cache given range (non-inclusive) from 0 to `end`
CREATE OR REPLACE FUNCTION get_cache_size()
RETURNS integer as $$
DECLARE
    size integer;
BEGIN
    SELECT INTO size COUNT(*)
    FROM fibonacci;
    RETURN size;
END;
$$ LANGUAGE plpgsql;


-- FUNC: get_cache_size_lt(end integer) (size integer)
--   Parameters:
--     range_end(integer) - End of range for measuring cache size
--   Returns:
--     size(integer) - Size of cache given range (non-inclusive) from 0 to `end`
CREATE OR REPLACE FUNCTION get_cache_size_lt(range_end integer)
RETURNS integer as $$
DECLARE
    size integer;
BEGIN
    SELECT INTO size COUNT(*)
    FROM fibonacci
    WHERE ind < range_end;
    RETURN size;
END;
$$ LANGUAGE plpgsql;
