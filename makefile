#!/usr/bin/env bash
###############################################################################
##
##       filename: makefile
##    description: make
##        created: 2023/09/16
##         author: ticktechman
##
###############################################################################
.PHONY: all build test install doc

build:
	go build -o build/srt2fcpxml cmd/main.go

clean:
	rm -Rf build

install: clean build
	sudo cp ./build/srt2fcpxml /usr/local/bin/
	cp -a ./res/* "${HOME}/Movies/Motion Templates.localized/Titles.localized/"

###############################################################################
