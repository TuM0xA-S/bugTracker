.PHONY: build debug lint tests run stop linters clean

build:
	go build -ldflags "-s -w"

debug:
	go build .

lint: linters
	go vet ./...
	revive
	errcheck

tests: run
	docker exec -t bugtracker_bug-tracker_1 go test ./...

run: stop clean
	docker-compose up -d

stop:
	docker-compose down

linters: 
	go get github.com/kisielk/errcheck
	go get github.com/mgechev/revive

clean:
	docker rmi bugtracker_bug-tracker
