default:
	just --list
build os arch:
	if [ "{{os}}" = "windows" ]; then \
		CGO_ENABLED=0 GOOS={{os}} GOARCH={{arch}} go build -o ./bin/mif-maker-{{os}}-{{arch}}.exe . ; \
	else \
		CGO_ENABLED=0 GOOS={{os}} GOARCH={{arch}} go build -o ./bin/mif-maker-{{os}}-{{arch}} . ;\
	fi

build-all:
	#!/usr/bin/env sh
	COMBINATIONS="darwin-amd64 darwin-arm64 linux-amd64 windows-amd64 windows-arm64"
	for comb in $COMBINATIONS; do
		GOOS="$(echo $comb | cut -d'-' -f1)"
		GOARCH="$(echo $comb | cut -d'-' -f2)"
		just build $GOOS $GOARCH
	done
