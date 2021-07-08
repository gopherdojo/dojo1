# Jpeg to png

Find jpeg files in your directory and convert to png files.

## Usage

### convert jpg to png
`$ go run *.go [your path]`

### test

* 1
  - `$ go test -coverprofile=profile github.com/aakira/jpg2png/file`  
  - `$ go test -coverprofile=profile github.com/aakira/jpg2png/image`

* 2 coverage
  - `$ go tool cover -html=profile`
