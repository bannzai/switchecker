PROJECT=switchecker

.PHONY: install
install: 
	go install github.com/bannzai/$(PROJECT)

.PHONY: test
test: install
	./scripts/test/run.sh
	go test ./

.PHONY: dry-run
dry-run: install
	$(PROJECT)

