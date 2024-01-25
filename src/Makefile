RMF=rm -f
BIN=gostatus
URL="http://127.0.0.1:6060/pkg/jepasq/gostatus/"

.PHONY: clean clean-binary

default: usage

usage:
	@echo "Please call 'make' with build, doc or check action"

build:
	go build

check:
	go test

# Thanks to https://askubuntu.com/a/1134795
open-doc:
	@echo "Opening $(URL)... (may need a refresh)"
	@xdg-open $(URL) || @sensible-browser $(URL) || @x-www-browser $(URL) || @gnome-open $(URL)
	@echo "Ctrl+C to exit"

doc: open-doc
	@godoc -http=:6060 > /dev/null

clean-binary:
	$(RMF) $(BIN)

clean: clean-binary
	$(RMF) *~
	$(RMF) */*~
