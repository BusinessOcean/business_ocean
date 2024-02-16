run:
	@go run main.go

generate:
	cd .proto;
	buf generate proto;




.PHONY run
