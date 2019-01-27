#################
SEARCH_GOFILES =  find -not -path '*/vendor/*' -type f -name "*.go"
##################

.PHONY: build run

build: server client

server: 
	docker build -t redis .
	docker run --rm -v $(CURDIR)/main:/go/src/redis/main redis go build -o $@ server.go

client: 
	docker run --rm -v $(CURDIR)/main:/go/src/redis/main redis go build -o $@ client.go

run:
	docker run --rm -d --name server -p 9090:9090 -v $(CURDIR)/main:/go/src/redis/main redis go run server.go
	docker run --rm -it --name client -v $(CURDIR)/main:/go/src/redis/main redis go run client.go -h 172.17.0.1

check:
	docker run -v $(CURDIR)/main:/go/src/redis/main redis sh -xc '\
                test -z "`$(SEARCH_GOFILES) -exec gofmt -s -l {} \;`" \
                && test -z "`$(SEARCH_GOFILES) -exec golint {} \;`"'

