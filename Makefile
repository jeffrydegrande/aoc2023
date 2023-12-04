day=day3

default: $(day)
.PHONY: default $(day)

build:
	go build

$(day): build
	./aoc2023 $(day)
