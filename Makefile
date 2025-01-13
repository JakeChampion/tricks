SRCS=$(wildcard netlify/go-functions/*.go)

OBJS=$(SRCS:.go=)

all: $(OBJS)

%:
	@if [ "$(@F)" != "all" ]; then\
		mkdir -p netlify/functions;\
		go get ./netlify/go-functions/;\
		go build -o netlify/functions/$(@F) ./netlify/go-functions/$(@F).go;\
    fi