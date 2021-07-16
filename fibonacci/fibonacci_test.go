package fibonacci_test

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dsykes16/pgfib/fibonacci"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

var _ = Describe("Fibonacci Tests", func() {
	var (
		pool  *dockertest.Pool
		pgcon *dockertest.Resource
		db    *sql.DB
		f     fibonacci.Fibonacci
	)

	DescribeTable("calculates fibonacci numbers correctly",
		func(index int, expected string) {
			res, err := f.GetFib(index)
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(expected))
		},
		Entry("Fib(0)", 0, "0"),
		Entry("Fib(1)", 1, "1"),
		Entry("Fib(2)", 2, "1"),
		Entry("Fib(3)", 3, "2"),
		Entry("Fib(4)", 4, "3"),
		Entry("Fib(5)", 5, "5"),
		Entry("Fib(6)", 6, "8"),
		Entry("Fib(7)", 7, "13"),
		Entry("Fib(8)", 8, "21"),
		Entry("Fib(9)", 9, "34"),
		Entry("Fib(10)", 10, "55"),
	)

	It("returns total cache size", func() {
		res, err := f.GetCacheSize()
		Expect(err).NotTo(HaveOccurred())
		Expect(res).To(Equal(0))

		f.GetFib(5)
		res, err = f.GetCacheSize()
		Expect(err).NotTo(HaveOccurred())
		Expect(res).To(Equal(4))
	})

	It("returns cache size given a range", func() {
		res, err := f.GetCacheSize()
		Expect(err).NotTo(HaveOccurred())
		Expect(res).To(Equal(0))

		f.GetFib(5)
		res, err = f.GetCacheSizeLT(3)
		Expect(err).NotTo(HaveOccurred())
		Expect(res).To(Equal(1))
	})

	JustAfterEach(func() {
		err := f.ClearCache()
		Expect(err).NotTo(HaveOccurred())
	})

	JustBeforeEach(func() {
		var err error
		f, err = fibonacci.New(db)
		Expect(err).NotTo(HaveOccurred())
	})

	BeforeSuite(func() {
		pool = newDockerPool()
		pgcon, db = startPostgres(pool)
	})

	AfterSuite(func() {
		if err := pool.Purge(pgcon); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	})

})

func newDockerPool() *dockertest.Pool {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker daemon: %s", err)
	}
	return pool
}

func startPostgres(pool *dockertest.Pool) (pgcontainer *dockertest.Resource, db *sql.DB) {
	var err error

	pgcontainer, err = pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: "postgres",
			Tag:        "13.3",
			Env: []string{
				"POSTGRES_USER=postgres",
				"POSTGRES_PASSWORD=testpass",
				"POSTGRES_DB=pgfib_test",
			},
		},
		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		log.Fatalf("could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		db, err = sql.Open(
			"postgres",
			fmt.Sprintf(
				"postgres://postgres:testpass@localhost:%s/pgfib_test?sslmode=disable",
				pgcontainer.GetPort("5432/tcp"),
			),
		)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to postgres: %s", err)
	}
	return
}
