MAKEFLAGS = -s

ql.go: ql.y
	goyacc -o ql.go -v ql.output -p ql ql.y
	gofmt -w ql.go

clean:
	rm -f ql.output ql.go
