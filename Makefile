PROJECT=switchecker

.PHONY: install
install: 
	go install github.com/bannzai/$(PROJECT)

.PHONY: test
test:
	go test ./

.PHONY: dry-run
dry-run: install
	$(PROJECT)

