SRC = $(wildcard */main.go)
TARGETS = $(patsubst %.go,%,$(SRC))

all: $(TARGETS)

$(TARGETS): %: %.go
	go build -o $@ $^
clean:
	rm -f $(TARGETS)
