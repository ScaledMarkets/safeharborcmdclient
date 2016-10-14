PRODUCTNAME=Safe Harbor Command Line Client
ORG=Scaled Markets
PACKAGENAME=safeharborcmdclient
EXECNAME=safeharbor

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
UTILITIESDIR:=$(realpath $(CURDIR)/../Utilities)

src_dir = $(CURDIR)/src
build_dir = $(CURDIR)/bin
test_dir = $(CURDIR)/test

.PHONY: all compile test_prep test_compile test clean info
.DEFAULT: all

all: compile

$(build_dir):
	mkdir $(build_dir)

$(build_dir)/$(EXECNAME): $(build_dir) $(src_dir)/$(PACKAGENAME)/*.go

# 'make compile' builds the executable, which is placed in <build_dir>.
compile: $(build_dir) $(src_dir)/$(PACKAGENAME)/*.go
	@echo "UTILITIESDIR=$(UTILITIESDIR)"
	@GOPATH=$(CURDIR):$(UTILITIESDIR) go install $(PACKAGENAME) -o $(EXECNAME)

$(pkg_dir)/$(CPU_ARCH)/$(PACKAGENAME)/*.a : compile

$(build_dir)/$(PACKAGENAME): compile

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
