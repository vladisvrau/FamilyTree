setup:
	go install github.com/rafaelsq/wtc@latest
	go install

clear-ports:
	@sh ./etc/killports.sh

dev:
	@go run github.com/rafaelsq/wtc