








build:
	go build .

test: build
	./go-compose tests/fixtures/composer_1.json
