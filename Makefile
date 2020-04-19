BINS=project1042

$(BINS): project1042.go
	GOOS=freebsd GOARCH=arm GOARM=6 go build .

clean:
	rm -rf $(BINS)

all: clean $(BINS)
