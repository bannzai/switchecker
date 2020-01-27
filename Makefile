PROJECT=switchecker

.PHONY: install
install: 
	go install github.com/bannzai/$(PROJECT)


.PHONY: dry-run
dry-run: install
	$(PROJECT)

