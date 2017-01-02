SRC = $(wildcard */main.go)
TARGETS = $(patsubst %.go,%,$(SRC))

$(TARGETS): %: %.go
	go build -o $@ $^
clean:
	rm -f $(TARGETS)
