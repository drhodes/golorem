include $(GOROOT)/src/Make.inc

TARG=lorem
GOFILES=\
	lorem.go\
	wordlist.go\

include $(GOROOT)/src/Make.pkg

format:
	gofmt -s -spaces=true -tabindent=false -tabwidth=4 -w lorem.go
	gofmt -s -spaces=true -tabindent=false -tabwidth=4 -w wordlist.go

