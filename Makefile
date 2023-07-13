default: usage

usage:
	echo "Please call with build, doc or check action"

build:
	go build

check:
	go test

doc:
	@echo "Please open http://127.0.0.1:6060/pkg/jepasq/gostatus/ ..."
	@echo "Ctrl+C to exit"
	@godoc -http=:6060 > /dev/null
