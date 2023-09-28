.PHONY: all c go run

POINTLIBPATH="../clib"

all: c go run

c:
	gcc -c -Wall -Werror -fpic -o ${POINTLIBPATH}/point.o ${POINTLIBPATH}/point.c
	gcc -shared -o ${POINTLIBPATH}/libpoint.so ${POINTLIBPATH}/point.o

go:
	go build -o app *.go

run:
	./app