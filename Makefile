all: cover

test:
	./.github/test.sh

cover: test
	go tool cover -html=coverage.txt
