package fibonacci

import (
	"database/sql"
	"os"
	"path"

	_ "github.com/lib/pq"
)

const (
	sqlInitPath = "../sql/fibonacci.sql"
)

type Fibonacci interface {
	GetFib(n int) (result string, err error)
	GetCacheSize() (size int, err error)
	GetCacheSizeLT(end int) (size int, err error)
	ClearCache() error
}

type PgFib struct {
	db *sql.DB
}

func New(db *sql.DB) (fib Fibonacci, err error) {
	err = db.Ping()
	if err != nil {
		return
	}

	err = initDb(db)
	if err != nil {
		return
	}

	fib = &PgFib{db: db}
	return
}

func (f *PgFib) GetFib(n int) (result string, err error) {
	row := f.db.QueryRow(`SELECT get_fib($1);`, n)
	result = ""
	err = row.Scan(&result)
	return
}

func (f *PgFib) GetCacheSize() (size int, err error) {
	row := f.db.QueryRow(`SELECT get_cache_size();`)
	size = 0
	err = row.Scan(&size)
	return
}

// Returns the number of cache entries in the exclusive range from 0 to `end`
func (f *PgFib) GetCacheSizeLT(end int) (size int, err error) {
	row := f.db.QueryRow(`SELECT get_cache_size_lt($1);`, end)
	size = 0
	err = row.Scan(&size)
	return
}

func (f *PgFib) ClearCache() error {
	_, err := f.db.Exec(`TRUNCATE TABLE fibonacci;`)
	return err
}

func initDb(db *sql.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	s, err := os.ReadFile(path.Join(cwd, sqlInitPath))
	if err != nil {
		return err
	}
	_, err = db.Exec(string(s))
	if err != nil {
		return err
	}
	return nil
}
