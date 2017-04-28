MODULE := pam_custom
NEED_SYMLINK := $(shell if ! stat -q .go/src/pam-custom 2>&1 > /dev/null ; then echo "yes" ; fi)

module: test
	GOPATH=${PWD}/.go go build -buildmode=c-shared -o ${MODULE}.so

test: *.go .go/src
	GOPATH=${PWD}/.go go test -cover

.go/src:
	-mkdir -p ${PWD}/.go/src
ifeq ($(NEED_SYMLINK),yes)
	ln -s ${PWD} ${PWD}/.go/src/pam-custom
endif

clean:
	go clean
	-rm -f ${MODULE}.so ${MODULE}.h
	-rm -rf .go/

.PHONY: test module download_deps clean
