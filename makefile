PRODUCTNAME=Safe Harbor Command Line Client
ORG=Scaled Markets
PACKAGENAME=safeharborcmdclient
EXECNAME=$(PACKAGENAME)

.DELETE_ON_ERROR:
.ONESHELL:
.SUFFIXES:
.DEFAULT_GOAL: all

SHELL = /bin/sh

CURDIR:=$(shell pwd)
UTILITIESDIR:=$(realpath $(CURDIR)/../Utilities)

.PHONY: all compile clean info
.DEFAULT: all

src_dir = $(CURDIR)/src

build_dir = $(CURDIR)/bin

all: compile

$(build_dir):
	mkdir $(build_dir)

$(build_dir)/$(EXECNAME): $(build_dir) $(src_dir)

# 'make compile' builds the executable, which is placed in <build_dir>.
compile: $(build_dir)/$(PACKAGENAME)

$(build_dir)/$(PACKAGENAME): $(build_dir)
	echo "UTILITIESDIR=$(UTILITIESDIR)"
	@GOPATH=$(CURDIR):$(UTILITIESDIR) go install $(PACKAGENAME)
