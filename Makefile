export GOPATH
TARGET := $(CURDIR)/bin

all:
	@echo "start building..."
	go build -o "$(TARGET)/raspberry"
	@echo "finish"