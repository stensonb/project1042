BINS=project1042

$(BINS): project1042.go
	GOOS=linux GOARCH=arm GOARM=6 go build .

clean:
	rm -rf $(BINS)

all: clean $(BINS)
