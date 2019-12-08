gobuildalien:
	@go build -v -o ./build/alien ./cmd/alien/*.go

gorunalien:
	make gobuildalien
	@./build/alien