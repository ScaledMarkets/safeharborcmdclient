PRODUCTNAME=Safe Harbor Command Line Client
ORG=Scaled Markets
PACKAGENAME=safeharborcmdclient
EXECNAME=safeharbor
CUCUMBER_CLASSPATH:=~/Library/Cucumber/*.jar

CUCUMBER_CLASSPATH=~/Library/Cucumber/cucumber-core-1.1.8.jar
CUCUMBER_CLASSPATH=$CUCUMBER_CLASSPATH:~/Library/Cucumber/cucumber-java-1.1.8.jar
CUCUMBER_CLASSPATH=$CUCUMBER_CLASSPATH:~/Library/Cucumber/cucumber-jvm-deps-1.0.3.jar
CUCUMBER_CLASSPATH=$CUCUMBER_CLASSPATH:~/Library/Cucumber/gherkin-2.12.2.jar
CUCUMBER_CLASSPATH=$CUCUMBER_CLASSPATH:~/Library/Cucumber/gherkin-jvm-deps-1.0.3.jar

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

.PHONY: all compile test clean info
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

test:
	java -cp $CUCUMBER_CLASSPATH cucumber.api.cli.Main \
		--plugin pretty --glue $(test_dir)/groovy $(test_dir)/resources \
		--tags @....

clean:
	rm -rf $(build_dir)/*

info:
	@echo "Makefile for $(PRODUCTNAME)."
