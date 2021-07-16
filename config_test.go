package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dsykes16/pgfib"
)

var _ = Describe("Server Initialization", func() {
	It("Loads Configuration", func() {
		expected := Config{
			DBHost:        "localhost",
			DBUser:        "postgres",
			DBPass:        "testpass",
			DBPort:        5432,
			DBName:        "pgfib",
			DBSSL:         false,
			ServerAddress: "0.0.0.0:5000",
			SQLInitPath:   "./sql/fibonacci.sql",
		}
		config, err := LoadConfig("./")
		Expect(err).NotTo(HaveOccurred())
		Expect(config).To(Equal(expected))
		Expect(config.ConnectionString()).To(Equal("postgres://postgres:testpass@localhost:5432/pgfib?sslmode=disable"))
	})
})
