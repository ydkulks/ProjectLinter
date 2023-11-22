## Running .sh file
## sh <this_file>.sh

##Testing
# go run ./cmd/main.go

## Production
# go build ./cmd/main.go
## or (for windows)
# GOOS=windows GOARCH=amd64 go build ./cmd/main.go

## Man page
## Note: get the projectlinter.1 file from the latest release
sudo cp -v projectlinter.1 /usr/local/man/man1
sudo mandb
