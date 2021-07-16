test:
	ginkgo -r

build:
	go mod download
	go build -o bin/pgfib
	chmod +x bin/pgfib

install: build
	sudo mkdir -p /etc/pgfib/sql
	sudo cp sql/fibonacci.sql /etc/pgfib/sql/
	sudo cp bin/pgfib /usr/bin/pgfib
	sudo chmod +x /usr/bin/pgfib
	sudo mkdir -p /etc/pgfib/config
	sudo cp app.env /etc/pgfib/config

docker-build: build
	docker build -t pgfib .
