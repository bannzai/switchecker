PROJECT=switchecker

.PHONY: install
install: 
	go install github.com/bannzai/$(PROJECT)

.PHONY: test
test: install
	./scripts/test/run.sh
	go test ./

.PHONY: ci-test
ci-test: install
	export PATH="${GOPATH}/bin:${PATH}"
	echo $(go env GOPATH)
	echo $(go env GOBIN)
	echo ${GOPATH}
	ls /bin
	which $(PROJECT)
	./scripts/test/run.sh
	make test

.PHONY: dry-run
dry-run: install
	$(PROJECT)

