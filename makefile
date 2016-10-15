PRODUCTNAME=Safe Harbor Command Line Client
ORG=Scaled Markets
PACKAGENAME=safeharborcmdclient
EXECNAME=safeharbor
CPU_ARCH:=$(shell uname -s | tr '[:upper:]' '[:lower:]')_amd64

HOME:=/Users/cliffordberg
CUCUMBER_CLASSPATH:=$(HOME)/Library/Cucumber/cucumber-core-1.1.8.jar
CUCUMBER_CLASSPATH:=$(CUCUMBER_CLASSPATH):$(HOME)/Library/Cucumber/cucumber-java-1.1.8.jar
CUCUMBER_CLASSPATH:=$(CUCUMBER_CLASSPATH):$(HOME)/Library/Cucumber/cucumber-jvm-deps-1.0.3.jar
CUCUMBER_CLASSPATH:=$(CUCUMBER_CLASSPATH):$(HOME)/Library/Cucumber/gherkin-2.12.2.jar
CUCUMBER_CLASSPATH:=$(CUCUMBER_CLASSPATH):$(HOME)/Library/Cucumber/gherkin-jvm-deps-1.0.3.jar

JSON:=$(HOME)/Library/json/*

SAFEHARBOR_HOST=54.71.85.235
SAFEHARBOR_PORT=

.DELETE_ON_ERROR:
.ONESHELL:
.SUFFIXES:
.DEFAULT_GOAL: all

SHELL = /bin/sh

CURDIR:=$(shell pwd)
UTILITIESDIR:=$(realpath $(CURDIR)/../utilities)

src_dir = $(CURDIR)/src
build_dir = $(CURDIR)/bin
test_dir = $(CURDIR)/test
pkg_dir = $(CURDIR)/pkg

.PHONY: all compile test_prep test_compile test clean info
.DEFAULT: all

all: compile

$(build_dir):
	mkdir $(build_dir)

# Main executable depends on source files.
$(build_dir)/$(EXECNAME): $(build_dir) $(src_dir)/$(PACKAGENAME)/*.go

# The compile target depends on the main executable.
# 'make compile' builds the executable, which is placed in <build_dir>.
compile: $(build_dir)/$(EXECNAME)
	GOPATH=$(CURDIR):$(UTILITIESDIR) go install $(PACKAGENAME)

cukever:
	java -cp $(CUCUMBER_CLASSPATH) cucumber.api.cli.Main --version

cukehelp:
	java -cp $(CUCUMBER_CLASSPATH) cucumber.api.cli.Main --help

test_check:
	java -cp $(CUCUMBER_CLASSPATH):$(JSON) cucumber.api.cli.Main $(test_dir)/features

test_compile:
	javac -cp $(CUCUMBER_CLASSPATH):$(JSON) $(test_dir)/steps/test/*.java

test: test_compile
	java -cp $(CUCUMBER_CLASSPATH):$(JSON):$(test_dir)/steps cucumber.api.cli.Main \
		--glue test $(test_dir)/features \
		--tags @done

clean:
	rm -rf $(build_dir)/*
	rm -f $(test_dir)/steps/test/*.class

info:
	@echo "Makefile for $(PRODUCTNAME)."
