CWD := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
srcfiles := $(wildcard *.md)
destfiles := $(patsubst %.md,%.pdf,$(srcfiles))

all: $(destfiles)

%.pdf: %.md
	@echo "$< --> $@"
	@docker run --rm -v "$(CWD):/data" -u "1000:1000" -ti pandoc/latex --standalone --pdf-engine=xelatex --output=$@ $<
	@echo "done"

clean:
ifeq ($(OS),Windows_NT)
	@del $(destfiles)
else
	@rm -f $(destfiles)
endif

.PHONY: all clean

