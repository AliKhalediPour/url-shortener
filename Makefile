BINARY_NAME=url-shortener

build:
	mkdir -p out/bin
	GO111MODULE=on go build -o out/bin/$(BINARY_NAME) ./cmd

run:
	export $(cat .env | xargs) && REDIS_HOST=localhost POSTGRES_HOST=localhost go run cmd/main.go

db-run:
	docker-compose -f docker-compose.yaml up -d

docker-run:
	docker build -t url -f Dockerfile . && docker run --env-file=.env --network=url-network -it --name url --rm -p 8000:8000 url
run:
